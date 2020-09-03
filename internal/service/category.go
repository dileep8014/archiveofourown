package service

import (
	"github.com/shyptr/archiveofourown/internal/model"
	"github.com/shyptr/archiveofourown/pkg/errwrap"
)

// category request example
type CategoryRequest struct {
	Name string `json:"name" example:"category name" binding:"required,max=100"`
}

// Category response example
type CategoryResponse struct {
	ID      int64  `json:"id" example:"1"`
	Name    string `json:"name" example:"电影"`
	WorkNum int64  `json:"workNum" example:"1"`
}

// CreateCategory: 创建分类
func (svc *Service) CreateCategory(req CategoryRequest) (err error) {
	defer errwrap.Wrap(&err, "service.category.create")

	result := svc.db.Create(&model.Category{Name: req.Name})
	err = CheckError(result, Insert_OP)
	return
}

// UpdateCategory: 修改分类信息
func (svc *Service) UpdateCategory(id int64, req CategoryRequest) (err error) {
	defer errwrap.Wrap(&err, "service.category.update")

	result := svc.db.Model(&model.Category{}).Updates(model.Category{ID: id, Name: req.Name})
	err = CheckError(result, Update_OP)
	return
}

// ListCategories: 分类列表
func (svc *Service) ListCategories() (list []CategoryResponse, err error) {
	defer errwrap.Wrap(&err, "service.category.list")

	result := svc.db.Model(&model.Category{}).Find(&list)
	err = CheckError(result, Select_OP)
	return
}

// DeleteCategory: 删除分类
func (svc *Service) DeleteCategory(id int64) (err error) {
	defer errwrap.Wrap(&err, "service.category.delete")

	result := svc.db.Delete(&model.Category{ID: id})
	err = CheckError(result, Delete_OP)
	return
}
