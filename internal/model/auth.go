package model

import (
	"github.com/shyptr/archiveofourown/pkg/errcode"
	"github.com/shyptr/archiveofourown/pkg/errwrap"
	"github.com/shyptr/sqlex"
)

// auth example
type Auth struct {
	AppKey    string `json:"appKey" example:"auth key"`
	AppSecret string `json:"appSecret" example:"auth secret"`
}

func (a Auth) Get(tx sqlex.BaseRunner) (auth Auth, err error) {
	defer errwrap.Wrap(&err, "model.auth.get")

	var count int
	err = sqlex.Select("count(*)").From("`auth`").
		Where("app_key=? AND app_secret=?", a.AppKey, a.AppSecret).
		RunWith(tx).QueryRow().Scan(&count)
	if err != nil {
		return
	}
	if count == 0 {
		err = errcode.NotExist
	}
	return
}
