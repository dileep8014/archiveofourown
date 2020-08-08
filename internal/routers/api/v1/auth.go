package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/shyptr/archiveofourown/internal/service"
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/errcode"
)

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
// @Success 200 {object} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /auth [get]
func (a Auth) Get(c *gin.Context) {
	logger := c.Value("logger").(zerolog.Logger)
	res := app.NewResponse(c)
	token := c.GetHeader("x-auth-token")
	if token != "" {
		res.Header("x-auth-token", token)
		res.ToResponse(gin.H{})
		return
	}
	idStr := c.Param("id")
	svc := service.NewService(c)
	defer svc.Finish()
	token, err := svc.Auth(idStr)
	if err != nil {
		logger.Error().AnErr("auth.get", err).Send()
		res.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
	}
	res.Header("x-auth-token", token)
	res.ToResponse(gin.H{})
}
