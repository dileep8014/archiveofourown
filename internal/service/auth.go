package service

import (
	"github.com/shyptr/archiveofourown/internal/model"
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/errwrap"
	"math/rand"
	"strconv"
)

const TEMP_USER = "未登录用户"

// Auth: 获取token
func (svc *Service) Auth(idStr string) (token string, err error) {
	defer errwrap.Wrap(&err, "service.Auth")

	var (
		id       int64
		username string
		root     bool
	)
	if idStr != "" {
		var i int
		i, err = strconv.Atoi(idStr)
		if err != nil {
			return
		}
		var user UserResponse
		err = svc.db.Model(&model.User{}).Select("id,username,root").Where(&model.User{ID: int64(i)}).First(&user).Error
		if err != nil {
			return token, err
		}
		id, username, root = user.ID, user.Username, user.Root
	} else {
		id = int64(rand.Intn(8))
		username = TEMP_USER
	}
	token, err = app.GenerateToken(id, username, root)
	return
}
