package service

import "time"

// Reply request
type ReplyRequest struct {
	CommentID int64  `json:"commentId" binding:"required"`
	Content   string `json:"content" binding:"required,maxlength=100,minlength=10"`
}

// Reply response
type ReplyResponse struct {
	ID        int64        `json:"id"`
	User      UserResponse `json:"user"`
	Content   string       `json:"content"`
	CreatedAt time.Time    `json:"createdAt"`
}
