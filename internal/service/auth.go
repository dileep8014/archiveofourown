package service

import (
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/errwrap"
	"math/rand"
	"strconv"
)

const TEMP_USER = "未登录用户"

func (svc *Service) Auth(idStr string) (token string, err error) {
	defer errwrap.Wrap(&err, "service.Auth")

	var id int64
	var username string
	if idStr != "" {
		i, err := strconv.Atoi(idStr)
		if err != nil {
			return
		}
		user, err := svc.GetUser(int64(i))
		if err != nil {
			return
		}
		id, username = user.ID, user.Username
	} else {
		id = int64(rand.Intn(8))
		username = TEMP_USER
	}
	token, err = app.GenerateToken(id, username)
	return
}
