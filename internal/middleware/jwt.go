package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/shyptr/archiveofourown/internal/cache"
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/errcode"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		token = c.GetHeader("Authorization")
		if token != "" {
			// jwt 验证token
			claims, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeOut
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			} else {
				// redis 判断token是否存在
				// 存在则获取用户部分信息，设置到上下文中
				exist, err := cache.Token{V: token}.Exists()
				if err != nil {
					ecode = errcode.UnauthorizedTokenError
				} else if !exist {
					ecode = errcode.UnauthorizedAuthNotExist
				} else {
					c.Set("me.id", claims.ID)
					c.Set("me.name", claims.Username)
					c.Set("me.root", claims.Root)
				}
			}
		}
		if ecode.Code != errcode.Success.Code {
			res := app.NewResponse(c)
			res.ToErrorResponse(ecode)
			c.Abort()
			return
		}
		c.Next()
	}
}
