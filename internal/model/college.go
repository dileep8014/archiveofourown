package model

import (
	"time"
)

// College example
type College struct {
	ID        int64
	UserId    int64
	Title     string
	Introduce string
	WorksNums int64
	CreatedBy string
	UpdatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c College) TableName() string {
	return "college"
}

// College work relation example
type CollegeWork struct {
	ID        int64
	CollegeId int64
	WorkId    int64
	CreatedBy string
	UpdatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c CollegeWork) TableName() string {
	return "college_work"
}
