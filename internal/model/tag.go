package model

import (
	"time"
)

// Tag
type Tag struct {
	ID        int64
	Name      string
	CreatedBy string
	UpdatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t Tag) TableName() string {
	return "`tag`"
}

// WorkTag
type WorkTag struct {
	ID        int64
	WorkID    int64
	TagID     int64
	CreatedBy string
	UpdatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (w WorkTag) TableName() string {
	return "work_tag"
}
