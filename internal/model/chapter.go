package model

import (
	"time"
)

// Chapter example
type Chapter struct {
	ID           int64
	WorkId       int64
	Status       int8
	Title        string
	Content      string
	Seq          int64
	Version      int64
	SubsectionId int64
	Lock         bool
	CreatedBy    string
	UpdatedBy    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (c Chapter) TableName() string {
	return "chapter"
}
