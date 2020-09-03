package model

import (
	"time"
)

type WorkType int

const (
	Doujin   = iota // 同人
	Original        // 原创
)

type Work struct {
	ID         int64
	Type       WorkType
	Title      string
	Introduce  string
	UserId     int64
	CategoryId int64
	TopicId    int64
	Lock       bool
	CreatedBy  string
	UpdatedBy  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (w Work) TableName() string {
	return "work"
}

type WorkEx struct {
	WorkId         int64
	Words          int64
	ViewNums       int64
	TalkNums       int64
	CollegeNums    int64
	SubscribeNums  int64
	ChapterNums    int64
	SubsectionNums int64
	DraftNums      int64
	RecycleNums    int64
	CreatedBy      string
	UpdatedBy      string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (w WorkEx) TableName() string {
	return "work_ex"
}
