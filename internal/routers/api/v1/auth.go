package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/shyptr/archiveofourown/internal/service"
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/errcode"
)

// Auth: 认证
type Auth struct {
}

func NewAuth() Auth {
	return Auth{}
}

// @Summary 获取token
// @Tags Auth
// @Accept json
// @Produce json
// @Param id path int false "用户ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /auth [get]
func (a Auth) Get(c *gin.Context) {
	res := app.NewResponse(c)
	token := c.GetHeader("x-auth-token")
	if token != "" {
		res.Header("x-auth-token", token)
		res.ToResponse(gin.H{})
		return
	}
	idStr := c.Param("id")
	svc := service.NewService(c)
	token, err := svc.Auth(idStr)
	if err != nil {
		res.ToErrorResponse(errcode.UnauthorizedTokenGenerate.WithError(err))
		return
	}
	res.Header("x-auth-token", token)
	res.ToResponse(gin.H{})
}
