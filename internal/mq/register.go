package mq

import (
	"encoding/json"
	"fmt"
	"github.com/shyptr/archiveofourown/global"
	"github.com/shyptr/archiveofourown/pkg/email"
	"time"
)

var registerQueueName = "register_queue"

type RegisterMsg struct {
	Email string `json:"email"`
	Path  string `json:"path"`
}

type RegisterConsumer struct {
	quit <-chan struct{}
}

func (r RegisterConsumer) Start() {
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
		registerQueueName,
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

	defaultMailer := email.NewEmail(&email.SMPTInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})

	body := `<h1>欢迎注册Pointer同人小说网</h1>
<p>尊敬的用户您好：</p>
<br/>
<br/>
<br/>
<p>感谢您注册Pointer同人小说网，您的注册信息完善地址为%s,请您完成信息填写，填写完毕后系统将为您注册账户。</p>`

	for {
		select {
		case <-r.quit:
			return
		case d := <-delivery:
			if d.Body == nil {
				continue
			}
			global.Logger.Info().Str("Received a message", string(d.Body)).Send()
			// 解析消息
			var msg RegisterMsg
			err := json.Unmarshal(d.Body, &msg)
			if err != nil {
				global.Logger.Error().Caller().AnErr("unmarshal register queue msg", err).Send()
				_ = d.Ack(false)
				continue
			}
			// 发送邮件
			global.Logger.Info().Str("register mq", "sending mail").Send()
			err = defaultMailer.SendMail([]string{msg.Email}, "请完善注册信息【Pointer同人小说网】", fmt.Sprintf(body, msg.Path))
			if err != nil {
				global.Logger.Error().Caller().AnErr("register email send failed", err).Send()
				_ = d.Reject(true)
				continue
			}
			global.Logger.Info().Str("register mq", "send mail success").Send()
			_ = d.Ack(false)
		}
	}
}

type RegisterProvider struct {
}

func (r RegisterProvider) Send(msg RegisterMsg, timer *time.Timer) {
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
