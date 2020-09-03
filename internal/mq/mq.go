package mq

import (
	"fmt"
	"github.com/shyptr/archiveofourown/global"
	"github.com/shyptr/archiveofourown/pkg/email"
	"github.com/streadway/amqp"
	"sync"
	"time"
)

var (
	pool sync.Pool
	once sync.Once
)

func InitMQ() {
	once.Do(func() {
		defaultMailer := email.NewEmail(&email.SMPTInfo{
			Host:     global.EmailSetting.Host,
			Port:     global.EmailSetting.Port,
			IsSSL:    global.EmailSetting.IsSSL,
			UserName: global.EmailSetting.UserName,
			Password: global.EmailSetting.Password,
			From:     global.EmailSetting.From,
		})
		pool.New = func() interface{} {
			log := global.Get()
			defer global.Put(log)
			conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/", global.MQSetting.Username,
				global.MQSetting.Password, global.MQSetting.Host))
			if err != nil {
				log.Error().Caller().AnErr("dial rabbitmq", err).Send()
				defaultMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("连接MQ失败，发生时间: %s", time.Now().Format(time.RFC3339)),
					fmt.Sprintf("错误信息: %v", err),
				)
				return nil
			}
			return conn
		}
	})
}

func getConn() *amqp.Connection {
	for c := pool.Get(); c != nil; c = pool.Get() {
		conn := c.(*amqp.Connection)
		if conn.IsClosed() {
			continue
		}
		return conn
	}
	return nil
}

func putConn(conn *amqp.Connection) {
	if !conn.IsClosed() {
		pool.Put(conn)
	}
}
