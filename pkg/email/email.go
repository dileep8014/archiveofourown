package email

import (
	"crypto/tls"
	"github.com/shyptr/archiveofourown/pkg/errwrap"
	"gopkg.in/gomail.v2"
)

type Email struct {
	*SMPTInfo
}

type SMPTInfo struct {
	Host     string
	Port     int
	UserName string
	Password string
	From     string
	IsSSL    bool
}

func NewEmail(info *SMPTInfo) *Email {
	return &Email{SMPTInfo: info}
}

func (e *Email) SendMail(to []string, subject, body string) (err error) {
	defer errwrap.Add(&err, "email.send")

	m := gomail.NewMessage()
	m.SetHeader("From", e.From)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSL}
	err = dialer.DialAndSend()
	return
}
