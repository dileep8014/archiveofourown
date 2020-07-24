package service

import (
	"github.com/shyptr/archiveofourown/internal/model"
	"github.com/shyptr/archiveofourown/pkg/errwrap"
)

func (svc *Service) GetUser(id int64) (user model.User, err error) {
	defer errwrap.Wrap(&err, "service.user.get")

	user.ID = id
	user, err = user.Get(svc.tx)
	return
}
