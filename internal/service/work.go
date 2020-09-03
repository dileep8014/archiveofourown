package service

import (
	"github.com/shyptr/archiveofourown/internal/model"
	"time"
)

// Work response example
type WorkResponse struct {
	ID             int64             `json:"id"`
	Type           model.WorkType    `json:"type"`
	Title          string            `json:"title"`
	Introduce      string            `json:"introduce"`
	UserId         int64             `json:"-"`
	User           UserResponse      `json:"user"`
	CategoryId     int64             `json:"-"`
	Category       CategoryResponse  `json:"category"`
	TopicId        int64             `json:"-"`
	Topic          TopicNameResponse `json:"topic"`
	Tags           []TagResponse     `json:"tags"`
	Words          int64             `json:"words"`
	ViewNums       int64             `json:"viewNums"`
	TalkNums       int64             `json:"talkNums"`
	CollegeNums    int64             `json:"collegeNums"`
	SubscribeNums  int64             `json:"subscribeNums"`
	ChapterNums    int64             `json:"chapterNums"`
	SubsectionNums int64             `json:"subsectionNums"`
	DraftNums      int64             `json:"draftNums"`
	RecycleNums    int64             `json:"recycleNums"`
	CreatedAt      time.Time         `json:"createdAt"`
	UpdatedAt      time.Time         `json:"updatedAt"`
}
