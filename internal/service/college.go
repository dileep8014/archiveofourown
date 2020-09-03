package service

import (
	"github.com/shyptr/archiveofourown/internal/model"
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/errwrap"
	"gorm.io/gorm"
)

// College create request example
type CollegeCreateRequest struct {
	Title     string `json:"title" binding:"required,maxlength=10"`
	Introduce string `json:"introduce" binding:"required,minlength=10,maxlength=200"`
}

// College update request example
type CollegeUpdateRequest struct {
	Title     string `json:"title" binding:"maxlength=10"`
	Introduce string `json:"introduce" binding:"minlength=10,maxlength=200"`
}

// College response example
type CollegeResponse struct {
	ID        int64
	Title     string
	Introduce string
	WorksNums int64
}

// College page response example
type CollegePageResponse struct {
	List  []CollegeResponse `json:"list"`
	Total int64             `json:"total"`
}

// College works page response example
type CollegeWorksPageResponse struct {
	List  []WorkResponse `json:"list"`
	Total int64          `json:"total"`
}

// ListCollegeWorks: 查询书单内作品列表
func (svc Service) ListCollegeWorks(id int64) (res CollegeWorksPageResponse, err error) {
	defer errwrap.Wrap(&err, "service.ListCollegeWorks")

	err = svc.db.Transaction(func(tx *gorm.DB) error {
		// 查询书单内作品列表
		size, offset := app.GetPage(svc.ctx)
		err := tx.Model(&model.CollegeWork{}).Joins("join work on college_work.work_id=work.id").
			Joins("left join work_ex on work_ex.word_id = work.id").
			Select("work.id, work.type, work.title, work.introduce, work.user_id, work.category_id,"+
				"work.topic_id, work.lock, work.created_at, work.updated_at, work_ex.words, work_ex.view_nums,"+
				"work_ex.talk_nums, work_ex.college_nums, work_ex.subscribe_nums, work_ex.chapter_nums,"+
				"work_ex.subsection_nums, work_ex.draft_nums, work_ex.draft_nums,work_ex.recycle_nums").
			Where("college_work.college_id=?", id).Count(&res.Total).Limit(size).Offset(offset).Find(&res.List).Error
		if err != nil {
			return err
		}
		for index, work := range res.List {
			// 用户信息
			err := tx.Model(&model.User{}).First(&res.List[index].User, &model.User{ID: work.UserId}).Error
			if err != nil {
				return err
			}
			// 分类信息
			err = tx.Model(&model.Category{}).First(&res.List[index].Category, &model.Category{ID: work.CategoryId}).Error
			if err != nil {
				return err
			}
			// 专题信息
			err = tx.Model(&model.Topic{}).First(&res.List[index].Topic, &model.Topic{ID: work.TopicId}).Error
			if err != nil {
				return err
			}
			// 标签列表
			err = tx.Model(&model.WorkTag{}).Joins("join tag on tag.id = work_tag.tag_id").
				Select("tag.id,tag.name").Where("work_tag.work_id=?", id).Find(&res.List[index].Tags).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	return
}
