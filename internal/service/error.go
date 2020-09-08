package service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/shyptr/archiveofourown/pkg/errcode"
)

var (
	InsertError = errors.New("插入失败，插入行数0条")
	UpdateError = errors.New("更新失败，更新行数0条")
	DeleteError = errors.New("删除失败，删除行数0条")
)

type GormOperator int

const (
	Insert_OP GormOperator = iota
	Update_OP
	Delete_OP
	Select_OP
)

var opErrs = []error{InsertError, UpdateError, DeleteError}

func CheckError(result *gorm.DB, op GormOperator) error {
	if result.Error == gorm.ErrRecordNotFound {
		return errcode.NotFound
	}
	if result.Error != nil {
		return result.Error
	}
	affected := result.RowsAffected
	if op != Select_OP && affected == 0 {
		return opErrs[op]
	}
	return nil
}
