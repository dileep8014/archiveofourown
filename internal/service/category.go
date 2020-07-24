package service

import (
	"github.com/shyptr/archiveofourown/internal/model"
	"github.com/shyptr/archiveofourown/pkg/errwrap"
)

// category request example
type CategoryRequest struct {
	Name string `json:"name" example:"category name" binding:"required,max=100"`
}

func (svc *Service) CreateCategory(req CategoryRequest) (category model.Category, err error) {
	defer errwrap.Wrap(&err, "service.category.create")

	category.Name = req.Name
	category, err = category.Create(svc.tx)
	return
}

func (svc *Service) UpdateCategory(id int64, req CategoryRequest) (category model.Category, err error) {
	defer errwrap.Wrap(&err, "service.category.update")

	category.ID, category.Name = id, req.Name
	category, err = category.Update(svc.tx)
	return
}

func (svc *Service) ListCategories() (categories []model.Category, err error) {
	defer errwrap.Wrap(&err, "service.category.list")

	category := model.Category{}
	categories, err = category.List(svc.tx)
	return
}

func (svc *Service) DeleteCategory(id int64) (err error) {
	defer errwrap.Wrap(&err, "service.category.delete")

	category := model.Category{ID: id}
	err = category.Delete(svc.tx)
	return
}
