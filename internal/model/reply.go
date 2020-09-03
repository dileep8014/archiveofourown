package model

import "time"

// Reply
type Reply struct {
	ID        int64
	UserId    int64
	CommentId int64
	Content   string
	CreatedBy string
	UpdatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
}
