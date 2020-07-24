package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/shyptr/archiveofourown/global"
	"github.com/shyptr/archiveofourown/internal/service"
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/errcode"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const TEMP_USER = "未登录用户"

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
// @Success 200 {object} service.AuthResponse "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /auth [get]
func (a Auth) Get(c *gin.Context) {
	logger := c.Value("logger").(zerolog.Logger)
	res := app.NewResponse(c)
	cookie, err := c.Cookie("me")
	if err != nil && err != http.ErrNoCookie {
		logger.Error().Caller().AnErr("auth.get", err).Send()
		res.ToErrorResponse(errcode.UnauthorizedTokenError)
		return
	}
	if cookie != "" {
		res.ToResponse(service.AuthResponse{Token: cookie})
		return
	}
	idStr := c.Param("id")
	var id int64
	var username string
	if idStr != "" {
		i, err := strconv.Atoi(idStr)
		if err != nil {
			logger.Error().Caller().AnErr("auth.get", err).Send()
			res.ToResponse(errcode.IDParseError)
			return
		}
		svc := service.NewService(c)
		defer svc.Finish()
		user, err := svc.GetUser(int64(i))
		if err != nil {
			logger.Error().Caller().AnErr("auth.get", err).Send()
			res.ToResponse(errcode.ErrorGetUserFail)
			return
		}
		id, username = user.ID, user.Username
	} else {
		id = int64(rand.Intn(8))
		username = TEMP_USER
	}
	token, err := app.GenerateToken(id, username)
	if err != nil {
		logger.Error().Caller().AnErr("auth.get", err).Send()
		res.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	c.SetCookie("me", token, global.JWTSetting.Expire*int(time.Second), "/", "localhost", false, false)
	res.ToResponse(service.AuthResponse{Token: token})
}
