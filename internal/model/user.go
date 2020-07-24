package model

import (
	"github.com/shyptr/archiveofourown/pkg/errwrap"
	"github.com/shyptr/archiveofourown/pkg/runner"
	"github.com/shyptr/sqlex"
	"time"
)

// user example
type User struct {
	ID        int64     `json:"id" example:"用户ID" format:"int64"`
	Username  string    `json:"username" example:"user name"`
	Email     string    `json:"email" example:"xxx@qq.com" format:"email"`
	Password  string    `json:"-"`
	Root      bool      `json:"root" example:"false"`
	CreatedAt time.Time `json:"createdAt" example:"2000-01-01 00:00:00" format:"date"`
	UpdatedAt time.Time `json:"updatedAt" example:"2000-01-01 00:00:00" format:"date"`
}

func (u User) Get(tx *runner.Runner) (user User, err error) {
	defer errwrap.Wrap(&err, "model.user.get")

	queryRow := sqlex.Select("username,email,password,root,created_at,updated_at").From("`user`").
		Where("id=?", u.ID).RunWith(tx).QueryRow()
	err = queryRow.Scan(&user.Username, &user.Email, &user.Password, &user.Root, &user.CreatedAt, &user.UpdatedAt)
	user.ID = u.ID
	return
}
