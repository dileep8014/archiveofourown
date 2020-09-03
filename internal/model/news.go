package model

import "time"

// News
type News struct {
	ID           int64
	Title        string
	Description  string
	Content      string
	CommentsNums int
	CreatedBy    string
	UpdatedBy    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
