package model

import (
	"time"
)

// Calendar example
type Calendar struct {
	ID        int64
	UserId    int64
	Year      int64
	Month     int64
	Day       int64
	Words     int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c Calendar) TableName() string {
	return "calender"
}
