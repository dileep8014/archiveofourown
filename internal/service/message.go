package service

import "time"

// Message response example
type MessageResponse struct {
	ID        int64        `json:"id"`
	Sender    UserResponse `json:"sender"`
	Content   string       `json:"content"`
	Readed    bool         `json:"readed"`
	CreatedAt time.Time    `json:"createdAt"`
}

// Message page response example
type MessagePageResponse struct {
	List  []MessageResponse `json:"list"`
	Total int               `json:"total"`
}

// Message request example
type MessageRequest struct {
	ReceiverID int64  `json:"receiverId" binding:"required"`
	Content    string `json:"content" binding:"required,maxlength=200"`
}
