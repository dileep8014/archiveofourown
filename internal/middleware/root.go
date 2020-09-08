package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/errcode"
)

func Root() gin.HandlerFunc {
	return func(c *gin.Context) {
		root, ok := c.Get("me.root")
		if !ok || !root.(bool) {
			res := app.NewResponse(c)
			res.ToErrorResponse(errcode.ErrorPermission.WithDetails("非管理员用户"))
			c.Abort()
			return
		}
		c.Next()
	}
}
