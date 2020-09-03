package model

import (
	"time"
)

type User struct {
	ID        int64
	Username  string
	Email     string
	Password  string
	Root      bool
	CreatedBy string
	UpdatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u User) TableName() string {
	return "user"
}

type UserEx struct {
	UserId    int64
	WorksNums int64
	WorkDay   int64
	Words     int64
	FansNums  int64
	CreatedBy string
	UpdatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ex UserEx) TableName() string {
	return "user_ex"
}
