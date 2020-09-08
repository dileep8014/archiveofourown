package model

import (
	"time"
)

type User struct {
	ID        int64
	Username  string
	Email     string
	Password  string
	Avatar    string
	Gender    int
	Root      bool
	Introduce string
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

type UserSt struct {
	UserId            int64
	ShowEmail         bool
	DisableSearch     bool
	ShowAdult         bool
	HiddenGrade       bool
	HiddenTag         bool
	SubscriptionEmail bool
	TopicEmail        bool
	CommentEmail      bool
	SystemEmail       bool
	CreatedBy         string
	UpdatedBy         string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (u UserSt) TableName() string {
	return "user_st"
}
