package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/errcode"
)

func Verify(calls ...func(c *gin.Context) (bool, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		res := app.NewResponse(c)
		for _, call := range calls {
			verify, err := call(c)
			if err != nil {
				res.CheckErrorAndResponse(err, errcode.UnauthorizedAuthNotExist)
				c.Abort()
				return
			}
			if !verify {
				res.ToErrorResponse(errcode.ErrorPermission)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

func LoginNeed(c *gin.Context) (bool, error) {
	_, ok := c.Get("me.name")
	if !ok {
		return false, errcode.UnauthorizedAuthNotExist
	}
	return true, nil
}
