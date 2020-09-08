package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/shyptr/archiveofourown/internal/service"
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/errcode"
)

// Calendar: 创作日历
type Calendar struct {
}

func NewCalendar() Calendar {
	return Calendar{}
}

func (c Calendar) Router(api gin.IRouter) {
	// 获取指定用户的创作日历
	api.GET("/calendar/:userID", c.List)
}

// @Summary 获取指定用户的创作日历
// @Tags 创作日历
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param userID path int true "用户ID"
// @Param year query int true "年份"
// @Param month query int true "月份" mininum(1) maxinum(12)
// @Success 200 {array} service.CalendarResponse "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/calendar [get]
func (c Calendar) List(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	// 解析参数
	userID, err := app.ShouldParamConvertInt(ctx, "userID")
	if err != nil {
		res.ToErrorResponse(errcode.InValidParams.WithDetails(err.Error()))
		return
	}
	req := service.CalendarRequest{}
	valid, errs := app.BindAndValid(ctx, &req, ctx.BindQuery)
	if !valid {
		res.ToErrorResponse(errcode.InValidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.NewService(ctx)

	list, err := svc.ListCalendar(int64(userID), req)
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorListCalendar)
		return
	}

	res.ToResponse(list)
}
