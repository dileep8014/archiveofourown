package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/shyptr/archiveofourown/global"
	"github.com/shyptr/archiveofourown/internal/middleware"
	"github.com/shyptr/archiveofourown/internal/model"
	"github.com/shyptr/archiveofourown/internal/service"
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/errcode"
)

// Chapter: 章节
type Chapter struct {
}

func NewChapter() Chapter {
	return Chapter{}
}

func (c Chapter) ChapterOwner(ctx *gin.Context) (bool, error) {
	id, err := app.ShouldParamConvertInt(ctx, "id")
	if err != nil {
		return false, errcode.InValidParams.WithDetails(err.Error())
	}
	ctx.Set("id", int64(id))
	chapter := model.Chapter{ID: int64(id)}
	var res int64
	err = global.Engine.Model(&chapter).Joins("right join work on work.id=chapter.work_id").
		Select("work.user_id").Where("chapter.id=?", chapter.ID).First(&res).Error
	return res == ctx.GetInt64("me.id"), err
}

func (c Chapter) Router(api gin.IRouter) {
	owner := middleware.Verify(middleware.LoginNeed, c.ChapterOwner)
	// 获取具体章节内容
	api.GET("/chapter/:id", c.Get)
	// 章节历史发布版本
	api.GET("/chapter/:id/history", owner, c.History)
	// 新建章节
	api.POST("/chapter", middleware.Verify(middleware.LoginNeed), c.New)
	// 保存章节
	api.POST("/chapter/:id", owner, c.Save)
	// 发布章节
	api.POST("/chapter/:id/publish", owner, c.Publish)
	// 回收章节
	api.POST("/chapter/:id/recycle", owner, c.Recycle)
	// 修改章节分卷
	api.PATCH("chapter/:id/:subsectionID", owner, c.Update)
	// 删除章节
	api.DELETE("/chapter/:id", owner, c.Delete)
}

// @Summary 查询章节内容
// @Tags 章节
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id path int true "章节ID"
// @Success 200 {object} service.ChapterContentResponse "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/chapter/{id} [get]
func (c Chapter) Get(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	svc := service.NewService(ctx)

	chapter, err := svc.GetChapter(ctx.GetInt64("id"))
	if err != nil && err == service.ChapterLockError {
		res.ToErrorResponse(errcode.ErrorGetChapter.WithDetails(err.Error()))
		return
	}
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorGetChapter)
		return
	}
	res.ToResponse(chapter)
}

// @Summary 查询章节历史版本
// @Tags 章节
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id path int true "章节ID"
// @Success 200 {array} service.ChapterHistoryResponse "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/chapter/{id}/history [get]
func (c Chapter) History(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	svc := service.NewService(ctx)

	list, err := svc.GetHistoryChapter(ctx.GetInt64("id"))
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorGetHistoryChapter)
		return
	}
	res.ToResponse(list)
}

// @Summary 新建草稿章节
// @Tags 章节
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param param body service.ChapterNewRequest true "参数"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/chapter [post]
func (c Chapter) New(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	req := service.ChapterNewRequest{}
	valid, errs := app.BindAndValid(ctx, &req, ctx.BindJSON)
	if !valid {
		res.ToErrorResponse(errcode.InValidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.NewService(ctx)

	err := svc.NewChapter(req)
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorNewChapter)
		return
	}
	res.ToResponse(gin.H{})
}

// @Summary 保存章节内容
// @Tags 章节
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id path int true "章节ID"
// @Param param body service.ChapterSaveRequest true "参数"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/chapter/{id} [post]
func (c Chapter) Save(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	req := service.ChapterSaveRequest{}
	valid, errs := app.BindAndValid(ctx, &req, ctx.BindJSON)
	if !valid {
		res.ToErrorResponse(errcode.InValidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.NewService(ctx)

	err := svc.SaveChapter(ctx.GetInt64("id"), req)
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorSaveChapter)
		return
	}
	res.ToResponse(gin.H{})
}

// @Summary 发布章节
// @Tags 章节
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id path int true "章节ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/chapter/{id}/publish [post]
func (c Chapter) Publish(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	svc := service.NewService(ctx)

	err := svc.PublishChapter(ctx.GetInt64("id"))
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorPublishChapter)
		return
	}
	res.ToResponse(gin.H{})
}

// @Summary 回收章节
// @Tags 章节
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id path int true "章节ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/chapter/{id}/publish [post]
func (c Chapter) Recycle(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	svc := service.NewService(ctx)

	err := svc.RecycleChapter(ctx.GetInt64("id"))
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorPublishChapter)
		return
	}
	res.ToResponse(gin.H{})
}

// @Summary 修改章节分卷信息
// @Tags 章节
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id path int true "章节ID"
// @Param subsectionID path int true "分卷ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/chapter/{id}/{subsectionID} [patch]
func (c Chapter) Update(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	req := service.ChapterPathRequest{}
	valid, errs := app.BindAndValid(ctx, &req, ctx.BindUri)
	if !valid {
		res.ToErrorResponse(errcode.InValidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.NewService(ctx)

	err := svc.UpdateChapterSubsection(req)
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorUpdateChapter)
		return
	}
	res.ToResponse(gin.H{})
}

// @Summary 删除章节
// @Tags 章节
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id path int true "章节ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/chapter/{id} [delete]
func (c Chapter) Delete(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	svc := service.NewService(ctx)

	err := svc.DeleteChapter(ctx.GetInt64("id"))
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorDeleteChapter)
		return
	}
	res.ToResponse(gin.H{})
}
