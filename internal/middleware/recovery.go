package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/shyptr/archiveofourown/global"
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/email"
	"github.com/shyptr/archiveofourown/pkg/errcode"
	"time"
)

func Recovery() gin.HandlerFunc {
	defaultMailer := email.NewEmail(&email.SMPTInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})

	return func(c *gin.Context) {
		c.Set("mailer", defaultMailer)
		defer func() {
			logger := c.Value("logger").(zerolog.Logger)
			if r := recover(); r != nil {
				logger.Error().Interface("panic recover", r).Send()

				err := defaultMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间: %s", time.Now().Format(time.RFC3339)),
					fmt.Sprintf("错误信息: %v", r),
				)
				if err != nil {
					logger.Panic().AnErr("panic recover", err).Send()
				}
				app.NewResponse(c).ToErrorResponse(errcode.ServerError.WithError(fmt.Errorf("panic: %v", r)))
				c.Abort()
			}
		}()
		c.Next()
	}
}
