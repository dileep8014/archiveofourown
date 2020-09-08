package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/shyptr/archiveofourown/internal/service"
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/errcode"
)

type Chapter struct {
}

func NewChapter() Chapter {
	return Chapter{}
}

func (c Chapter) Router(api gin.IRouter) {
	api.POST("chapter/:id/lock", c.Lock)
	api.POST("chapter/:id/unlock", c.Lock)
}

// @Summary 锁住章节
// @Tags 章节
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id path int true "章节ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/chapter/{id}/lock [post]
func (c Chapter) Lock(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	id, err := app.ShouldParamConvertInt(ctx, "id")
	if err != nil {
		res.ToErrorResponse(errcode.InValidParams.WithDetails(err.Error()))
		return
	}

	svc := service.NewService(ctx)

	err = svc.LockChapter(int64(id))
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorLockChapter)
		return
	}
	res.ToResponse(gin.H{})
}

// @Summary 解锁章节
// @Tags 章节
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id path int true "章节ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/chapter/{id}/unlock [post]
func (c Chapter) UnLock(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	id, err := app.ShouldParamConvertInt(ctx, "id")
	if err != nil {
		res.ToErrorResponse(errcode.InValidParams.WithDetails(err.Error()))
		return
	}

	svc := service.NewService(ctx)

	err = svc.UnLockChapter(int64(id))
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorUnLockChapter)
		return
	}
	res.ToResponse(gin.H{})
}
