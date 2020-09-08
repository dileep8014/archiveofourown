package service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/shyptr/archiveofourown/internal/model"
	"github.com/shyptr/archiveofourown/internal/mq"
	"github.com/shyptr/archiveofourown/pkg/errcode"
	"github.com/shyptr/archiveofourown/pkg/errwrap"
	"time"
)

// Chapter create request example
type ChapterNewRequest struct {
	WorkID       int64  `json:"workId" binding:"required"`
	Title        string `json:"title" binding:"required,maxlength=20"`
	Content      string `json:"content" binding:"required"`
	SubsectionID int64  `json:"subsectionId" binding:"required"`
}

// Chapter save request example
type ChapterSaveRequest struct {
	Title   string `json:"title" binding:"required,maxlength=20"`
	Content string `json:"content" binding:"required"`
}

// Chapter update path request example
type ChapterPathRequest struct {
	ID           int64 `json:"id" binding:"required"`
	SubsectionID int64 `json:"subsectionID" binding:"required"`
}

// Chapter content response example
type ChapterContentResponse struct {
	Title   string `json:"title" example:"第一章 起始"`
	Content string `json:"content" example:"正文内容"`
	Lock    bool   `json:"-"`
}

// Chapter history response example
type ChapterHistoryResponse struct {
	ID        int64     `json:"id" example:"1"`
	Title     string    `json:"title" example:"第一章 起始"`
	Content   string    `json:"content" example:"正文内容"`
	Version   int64     `json:"version" example:"1"`
	UpdatedAt time.Time `json:"updatedAt" example:"2020-1-1 00:00:00"`
}

var ChapterLockError = errors.New("章节已被管理员锁定，无法获取")

// GetChapter: 获取指定章节内容
func (svc Service) GetChapter(id int64) (res ChapterContentResponse, err error) {
	defer errwrap.Wrap(&err, "service.GetChapter")

	result := svc.db.Model(&model.Chapter{ID: id}).First(&res)
	err = CheckError(result, Select_OP)
	if err != nil {
		return
	}
	if res.Lock {
		return ChapterContentResponse{}, ChapterLockError
	}
	return
}

// GetHistoryChapter: 获取指定章节历史发布版本
func (svc Service) GetHistoryChapter(id int64) (list []ChapterHistoryResponse, err error) {
	defer errwrap.Wrap(&err, "service.GetHistoryChapter")

	chapter := model.Chapter{ID: id}
	err = svc.db.First(&chapter).Error
	if err != nil {
		return
	}
	result := svc.db.Where("work_id=? AND seq=? AND version<?", chapter.WorkId, chapter.Seq, chapter.Version).Find(&list)
	err = CheckError(result, Select_OP)
	return
}

// NewChapter: 新建章节
func (svc Service) NewChapter(req ChapterNewRequest) (err error) {
	defer errwrap.Wrap(&err, "service.NewChapter")

	// 验证资源所有者
	work := model.Work{ID: req.WorkID}
	err = svc.db.Select("user_id").First(&work).Error
	if err != nil {
		return
	}
	if work.UserId != svc.ctx.GetInt64("me.id") {
		return errcode.ErrorPermission
	}
	result := svc.db.Create(&model.Chapter{
		WorkId:       req.WorkID,
		Title:        req.Title,
		Content:      req.Content,
		SubsectionId: req.SubsectionID,
	})
	err = CheckError(result, Insert_OP)
	return
}

// SaveChapter: 保存章节内容
func (svc Service) SaveChapter(id int64, req ChapterSaveRequest) (err error) {
	defer errwrap.Wrap(&err, "service.NewChapter")

	err = svc.db.Transaction(func(tx *gorm.DB) error {
		chapter := model.Chapter{ID: id}
		// 获取章节信息
		err := tx.First(&chapter).Error
		if err != nil {
			return err
		}
		// 若章节为草稿章节,直接更新
		if chapter.Status == 0 {
			result := tx.Model(&chapter).Updates(&req)
			return CheckError(result, Update_OP)
		}
		if chapter.Status == 1 {
			//添加新的发布版本
			newChapter := model.Chapter{
				WorkId:       chapter.WorkId,
				Status:       1,
				Title:        req.Title,
				Content:      req.Content,
				Seq:          chapter.Seq,
				Version:      chapter.Version + 1,
				SubsectionId: chapter.SubsectionId,
			}
			result := tx.Create(&newChapter)
			err := CheckError(result, Insert_OP)
			if err != nil {
				return err
			}
			// 发送消息队列
			go mq.CalendarProvider{}.Send(mq.CalendarMessage{
				ChapterID: newChapter.ID,
				UserID:    svc.ctx.Value("me.id").(int64),
				Date:      time.Now(),
			}, time.NewTimer(20*time.Second))
		}
		return nil
	})
	return
}

// PublishChapter: 发布章节
func (svc Service) PublishChapter(id int64) (err error) {
	defer errwrap.Wrap(&err, "service.PublishChapter")

	err = svc.db.Transaction(func(tx *gorm.DB) error {
		chapter := model.Chapter{ID: id}
		// 获取章节信息
		err := tx.First(&chapter).Error
		if err != nil {
			return err
		}
		// 若当前章节已发布，直接返回
		if chapter.Status == 1 {
			return nil
		}
		// 若为草稿章节,查询当前最新发布章节序号，修改发布状态和发布序号
		if chapter.Status == 0 {
			err := tx.Model(&chapter).Where("work_id=?", chapter.WorkId).Select("max(seq)").First(&chapter.Seq).Error
			if err != nil {
				return err
			}
			chapter.Status = 1
			result := tx.Model(&chapter).Updates(chapter)
			err = CheckError(result, Update_OP)
			if err != nil {
				return err
			}
			// 发送消息队列
			go mq.CalendarProvider{}.Send(mq.CalendarMessage{
				ChapterID: chapter.ID,
				UserID:    svc.ctx.Value("me.id").(int64),
				Date:      time.Now(),
			}, time.NewTimer(20*time.Second))
		}
		return nil
	})
	return
}

// RecycleChapter: 回收章节
func (svc Service) RecycleChapter(id int64) (err error) {
	defer errwrap.Wrap(&err, "service.RecycleChapter")

	result := svc.db.Model(&model.Chapter{}).Update("status", 0).Where("id=? AND status=?", id, 2)
	err = CheckError(result, Update_OP)
	return
}

// UpdateChapterSubsection: 修改章节所在分卷
func (svc Service) UpdateChapterSubsection(req ChapterPathRequest) (err error) {
	defer errwrap.Wrap(&err, "service.UpdateChapterSubsection")

	result := svc.db.Model(&model.Chapter{ID: req.ID}).Update("subsection_id", req.SubsectionID)
	err = CheckError(result, Update_OP)
	return
}

// DeleteChapter: 删除章节
func (svc Service) DeleteChapter(id int64) (err error) {
	defer errwrap.Wrap(&err, "service.DeleteChapter")

	err = svc.db.Transaction(func(tx *gorm.DB) error {
		chapter := model.Chapter{ID: id}
		// 获取章节信息
		err := tx.First(&chapter).Error
		if err != nil {
			return err
		}
		// 若当前章节为发布章节，修改为回收状态
		if chapter.Status == 1 {
			result := tx.Model(&chapter).Update("status", 2)
			return CheckError(result, Update_OP)
		}
		// 直接删除
		result := tx.Delete(&chapter)
		return CheckError(result, Delete_OP)
	})
	return
}

// LockChapter: 锁住章节
func (svc Service) LockChapter(id int64) (err error) {
	defer errwrap.Wrap(&err, "service.LockChapter")

	result := svc.db.Model(&model.Chapter{ID: id}).Update("lock", true)
	err = CheckError(result, Update_OP)
	return
}

// UnLockChapter: 解锁章节
func (svc Service) UnLockChapter(id int64) (err error) {
	defer errwrap.Wrap(&err, "service.UnLockChapter")

	result := svc.db.Model(&model.Chapter{ID: id}).Update("lock", true)
	err = CheckError(result, Update_OP)
	return
}
