package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/errcode"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		token = c.GetHeader("x-auth-token")
		if token == "" {
			ecode = errcode.InValidParams
		} else {
			claims, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeOut
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			} else {
				c.Set("me.id", claims.ID)
				c.Set("me.name", claims.Username)
				c.Set("me.root", claims.Root)
			}
		}
		if ecode != errcode.Success {
			res := app.NewResponse(c)
			res.ToErrorResponse(ecode)
			c.Abort()
			return
		}
		c.Next()
		c.Header("x-auth-token", token)
	}
}
