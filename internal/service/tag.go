package service

import "github.com/shyptr/archiveofourown/internal/model"

// Tag response
type TagResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func NewTagListResponse(tags []model.Tag) []TagResponse {
	list := make([]TagResponse, len(tags))
	for index, item := range tags {
		list[index] = TagResponse{
			ID:   item.ID,
			Name: item.Name,
		}
	}
	return list
}
