package model

import (
	"github.com/shyptr/archiveofourown/pkg/errwrap"
	"github.com/shyptr/sqlex"
)

// Category example
type Category struct {
	ID   int64  `json:"id" example:"1" format:"int64"`
	Name string `json:"name" example:"category name"`
}

func (c Category) Create(tx sqlex.BaseRunner) (category Category, err error) {
	defer errwrap.Wrap(&err, "model.category.Create")

	result, err := sqlex.Insert("category").Columns("name").Values(c.Name).RunWith(tx).Exec()
	if err != nil {
		return
	}
	c.ID, _ = result.LastInsertId()
	category = c
	return
}

func (c Category) List(tx sqlex.BaseRunner) (categories []Category, err error) {
	defer errwrap.Wrap(&err, "model.category.List")

	rows, err := sqlex.Select("id,name").From("category").RunWith(tx).Query()
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var category Category
		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return
		}
		categories = append(categories, category)
	}
	return
}

func (c Category) Update(tx sqlex.BaseRunner) (category Category, err error) {
	defer errwrap.Wrap(&err, "model.category.update")

	_, err = sqlex.Update("category").Set("name", c.Name).RunWith(tx).Exec()
	if err != nil {
		return
	}
	category = c
	return
}

func (c Category) Delete(tx sqlex.BaseRunner) (err error) {
	defer errwrap.Wrap(&err, "model.category.delete")

	_, err = sqlex.Delete("category").Where("id=?", c.ID).RunWith(tx).Exec()
	return
}
