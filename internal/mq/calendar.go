package mq

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/shyptr/archiveofourown/global"
	"github.com/shyptr/archiveofourown/internal/model"
	"time"
)

var calendarQueueName = "calendar_queue"

type CalendarMessage struct {
	ChapterID int64     `json:"chapterId"`
	UserID    int64     `json:"userId"`
	Date      time.Time `json:"date"`
}

type CalendarConsumer struct {
	quit <-chan struct{}
}

func (c CalendarConsumer) Start() {
	// 获取MQ连接
	conn := getConn()
	defer conn.Close()
	channel, err := conn.Channel()
	if err != nil {
		global.Logger.Fatal().Caller().AnErr("open channel", err).Send()
	}
	defer channel.Close()
	// 定义queue
	queue, err := channel.QueueDeclare(
		calendarQueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		global.Logger.Fatal().Caller().AnErr("queue declare", err).Send()
	}
	// 设置channel的Qos
	err = channel.Qos(1, 0, false)
	if err != nil {
		global.Logger.Fatal().Caller().AnErr("set qos", err).Send()
	}
	// 注册消费者
	delivery, err := channel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		global.Logger.Fatal().Caller().AnErr("consume register", err).Send()
	}
	// 消费信息，直到退出
	for {
		select {
		case <-c.quit:
			return
		case d := <-delivery:
			if d.Body == nil {
				continue
			}
			global.Logger.Info().Str("Received a message", string(d.Body)).Send()
			// 解析消息
			var msg CalendarMessage
			err := json.Unmarshal(d.Body, &msg)
			if err != nil {
				global.Logger.Error().Caller().AnErr("unmarshal calendar queue msg", err).Send()
				_ = d.Ack(false)
				continue
			}
			// 开启数据库事务
			err = global.Engine.Transaction(func(tx *gorm.DB) error {
				// 获取当前版本章节内容
				chapter := model.Chapter{ID: msg.ChapterID}
				err := tx.First(&chapter).Error
				if err != nil {
					return err
				}
				words := int64(len([]rune(chapter.Content)))
				// 检查此时是否已经有日历数据
				calendar := model.Calendar{
					UserId: msg.UserID,
					Year:   int64(msg.Date.Year()),
					Month:  int64(msg.Date.Month()),
					Day:    int64(msg.Date.Day()),
				}
				err = tx.Where(&calendar).First(&calendar).Error
				// 没有日历数据则直接插入
				if err == gorm.ErrRecordNotFound {
					calendar.Words = words
					return tx.Create(&calendar).Error
				}
				if err != nil {
					return err
				}
				// 若无前置版本
				if chapter.Version == 1 {
					return tx.Model(&calendar).Update("words", calendar.Words+words).Error
				}
				// 查找前置版本
				lastVersionChapter := model.Chapter{WorkId: chapter.WorkId, Seq: chapter.Seq, Version: chapter.Version - 1}
				err = tx.Where(&lastVersionChapter).First(&lastVersionChapter).Error
				if err != nil {
					return err
				}
				calendar.Words = calendar.Words + words - int64(len([]rune(lastVersionChapter.Content)))
				err = tx.Updates(&calendar).Error
				if err != nil {
					return err
				}
				return nil
			})
			if err != nil {
				global.Logger.Error().Caller().Err(err).Send()
				_ = d.Reject(true)
				continue
			}
			_ = d.Ack(false)
		}
	}
}

type CalendarProvider struct {
}

func (c CalendarProvider) Send(msg CalendarMessage, timer *time.Timer) {

	body, _ := json.Marshal(msg)
	for {
		select {
		case <-timer.C:
			return
		default:
			err := SendMessage(registerQueueName, body)
			if err == nil {
				return
			}
		}
	}
}
