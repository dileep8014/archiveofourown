package middleware

import (
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/errcode"
	"strconv"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		token, _ = c.Cookie("me")
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
				idStr, _ := hex.DecodeString(claims.ID)
				id, _ := strconv.ParseInt(string(idStr), 10, 64)
				username, _ := hex.DecodeString(claims.Username)
				c.Set("me.id", id)
				c.Set("me.name", string(username))
			}
		}
		if ecode != errcode.Success {
			res := app.NewResponse(c)
			res.ToErrorResponse(ecode)
			c.Abort()
			return
		}
		c.Next()
	}
}