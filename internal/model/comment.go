package model

import "time"

type Comments struct {
	ID        int64
	UserId    int64
	ObjType   int
	ObjId     int64
	Content   string
	CreatedBy string
	UpdatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c Comments) TableName() string {
	return "comments"
}
