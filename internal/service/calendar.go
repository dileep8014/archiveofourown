package service

import (
	"github.com/shyptr/archiveofourown/internal/model"
	"github.com/shyptr/archiveofourown/pkg/errwrap"
)

// calendar request example
type CalendarRequest struct {
	Year  int64 `json:"year" binding:"required"`
	Month int64 `json:"month" binding:"required,min=1,max=12"`
}

// calendar response example
type CalendarResponse struct {
	Day   int64 `json:"day" example:"1"`
	Words int64 `json:"words" example:"4000"`
}

// 获取指定用户的创作日历
func (svc Service) ListCalendar(userID int64, req CalendarRequest) (list []CalendarResponse, err error) {
	defer errwrap.Wrap(&err, "service.ListCalendar")

	result := svc.db.Where(&model.Calendar{UserId: userID, Year: req.Year, Month: req.Month}).Find(&list)
	err = CheckError(result, Select_OP)

	return
}
