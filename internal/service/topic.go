package service

import (
	"github.com/shyptr/archiveofourown/internal/model"
	"time"
)

// Topic request
type TopicRequest struct {
	Title       string `json:"title" binding:"required,maxlength=20"`
	CategoryID  int64  `json:"categoryId" binding:"required"`
	Original    string `json:"original" binding:"required"`
	OriginalUrl string `json:"originalUrl" binding:"required"`
	Description string `json:"description" binding:"required,maxlength=300"`
}

// Topic update request
type TopicUpdateRequest struct {
	Title       string `json:"title" binding:"maxlength=20"`
	CategoryID  int64  `json:"categoryId" `
	Original    string `json:"original" `
	OriginalUrl string `json:"originalUrl" `
	Description string `json:"description" binding:"maxlength=300"`
}

// Topic name response
type TopicNameResponse struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

func NewTopicNameResponse(t model.Topic) TopicNameResponse {
	return TopicNameResponse{
		ID:    t.ID,
		Title: t.Title,
	}
}

// Topic response
type TopicResponse struct {
	ID            int64            `json:"id"`
	Title         string           `json:"title"`
	Category      CategoryResponse `json:"category"`
	WorkNum       int64            `json:"workNum"`
	SubscribeNums int64            `json:"subscribeNums"`
	Original      string           `json:"original"`
	OriginalUrl   string           `json:"originalUrl"`
	Description   string           `json:"description"`
	CreatedAt     time.Time        `json:"createdAt"`
	UpdatedAt     time.Time        `json:"updatedAt"`
}

func NewTopicResponse(t model.Topic) TopicResponse {
	return TopicResponse{
		ID:            t.ID,
		Title:         t.Title,
		Category:      CategoryResponse{},
		WorkNum:       t.WorkNum,
		SubscribeNums: t.SubscribeNums,
		Original:      t.Original,
		OriginalUrl:   t.OriginalUrl,
		Description:   t.Description,
		CreatedAt:     t.CreatedAt,
		UpdatedAt:     t.UpdatedAt,
	}
}

// Topic page response
type TopicPageResponse struct {
	List  []TopicResponse `json:"list"`
	Total int             `json:"total"`
}
