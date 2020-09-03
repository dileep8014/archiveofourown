package service

import "time"

// Comment create request example
type CommentRequest struct {
	ObjType int    `json:"objType" binding:"required,one of 0 1 2 3"`
	ObjID   int64  `json:"objId" binding:"required"`
	Content string `json:"content" binding:"required,minlength=10,maxlength=200"`
}

// Comment response example
type CommentResponse struct {
	ID        int64        `json:"id"`
	User      UserResponse `json:"user"`
	Content   string       `json:"content"`
	UpdatedAt time.Time    `json:"updatedAt"`
}

// Comment page response example
type CommentPageResponse struct {
	List  []CommentResponse `json:"list"`
	Total int64             `json:"total"`
}
