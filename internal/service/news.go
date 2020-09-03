package service

import "time"

// News request
type NewsRequest struct {
	Title   string `json:"title" binding:"required,maxlength=20"`
	Content string `json:"content" binding:"required,maxlength=300"`
}

type NewsUpdateRequest struct {
	Title   string `json:"title" binding:"maxlength=20"`
	Content string `json:"content" binding:"maxlength=300"`
}

// News page response
type NewsPageResponse struct {
	List []struct {
		ID          int64     `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		CommentNums int       `json:"commentNums"`
		UpdatedAt   time.Time `json:"updatedAt"`
	} `json:"list"`
	Total int `json:"total"`
}

// News response
type NewsResponse struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	CommentNums int       `json:"commentNums"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
