package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/shyptr/archiveofourown/internal/service"
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/errcode"
)

func Verify(call func(c *gin.Context) (bool, error)) gin.HandlerFunc {
	return func(c *gin.Context) {

		verify, err := call(c)
		res := app.NewResponse(c)
		if err != nil {
			res.ToErrorResponse(errcode.VerifyError.WithError(err))
			c.Abort()
			return
		}
		if !verify {
			res.ToErrorResponse(errcode.ErrorPermission)
			c.Abort()
			return
		}
		c.Next()
	}
}

func LoginNeed(c *gin.Context) (bool, error) {
	userName := c.GetString("me.name")
	if userName == service.TEMP_USER {
		return false, nil
	}
	return true, nil
}

func ResourceOwner(param string, belong func(param string, c *gin.Context) (int64, error)) func(c *gin.Context) (bool, error) {
	return func(c *gin.Context) (bool, error) {
		userId := c.GetInt64("me.id")
		userName := c.GetString("me.name")
		root := c.GetBool("me.root")
		if userName == service.TEMP_USER {
			return false, nil
		}
		if root {
			return true, nil
		}
		p, err := belong(param, c)
		if err != nil {
			return false, err
		}
		return userId == p, nil
	}
}
