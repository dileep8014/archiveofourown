package model

import (
	"time"
)

// Topic
type Topic struct {
	ID            int64
	Title         string
	CategoryId    int64
	UserId        int64
	WorkNum       int64
	SubscribeNums int64
	Original      string
	OriginalUrl   string
	Description   string
	CreatedBy     string
	UpdatedBy     string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (t Topic) TableName() string {
	return "topic"
}
