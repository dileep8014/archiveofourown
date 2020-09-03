package model

import "time"

// Message
type Message struct {
	ID         int64
	Type       int
	ReceiverId int64
	SenderId   int64
	Content    string
	Readed     bool
	CreatedBy  string
	UpdatedBy  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
