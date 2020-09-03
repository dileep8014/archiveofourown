package model

import "time"

// Category example
type Category struct {
	ID        int64
	Name      string
	WorkNum   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c Category) TableName() string {
	return "`category`"
}
