package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/shyptr/archiveofourown/global"
	"github.com/shyptr/archiveofourown/internal/middleware"
	"github.com/shyptr/archiveofourown/internal/model"
	"github.com/shyptr/archiveofourown/internal/service"
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/errcode"
	"strconv"
)

// College: 书单
type College struct {
}

func NewCollege() College {
	return College{}
}

func (c College) Owner(ctx *gin.Context) (bool, error) {
	id, err := app.ShouldParamConvertInt(ctx, "id")
	if err != nil {
		return false, errcode.InValidParams.WithDetails(err.Error())
	}
	ctx.Set("id", int64(id))
	var userId int64
	err = global.Engine.Model(&model.College{}).Select("user_id").First(&userId).Error
	return userId == ctx.GetInt64("me.id"), err
}

func (c College) Router(api gin.IRouter) {
	owner := middleware.Verify(middleware.LoginNeed, c.Owner)
	// 获取书单内作品列表
	api.GET("/college/:id/works", c.Works)
	// 获取书单列表
	api.GET("/college", c.List)
	// 创建书单
	api.POST("/college", middleware.Verify(middleware.LoginNeed), c.Create)
	// 添加作品
	api.POST("/college/:id/add/:workId", owner, c.Add)
	// 修改书单信息
	api.PUT("/college/:id", owner, c.Update)
	// 删除书单
	api.DELETE("/college/:id", owner, c.Delete)
}

// @Summary 获取书单内作品列表
// @Tags 书单
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id path int true "书单ID"
// @Param pageSize query int false "页大小"
// @Param page query int false "当前页"
// @Success 200 {object} service.CollegeWorksPageResponse "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/college/{id}/works [get]
func (c College) Works(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	id, err := app.ShouldParamConvertInt(ctx, "id")
	if err != nil {
		res.ToErrorResponse(errcode.InValidParams.WithDetails(err.Error()))
		return
	}

	svc := service.NewService(ctx)

	works, err := svc.ListCollegeWorks(int64(id))
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorCollegeWorks)
		return
	}
	res.ToResponse(works)
}

// @Summary 查询书单列表
// @Tags 书单
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param pageSize query int false "页大小"
// @Param page query int false "当前页"
// @Param userId query int true "用户ID"
// @Success 200 {object} service.CollegePageResponse "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/college [get]
func (c College) List(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	query, ok := ctx.GetQuery("userId")
	if !ok {
		res.ToErrorResponse(errcode.InValidParams.WithDetails("用户ID为空"))
		return
	}
	userId, err := strconv.Atoi(query)
	if err != nil {
		res.ToErrorResponse(errcode.InValidParams.WithDetails("用户ID解析错误"))
		return
	}
	svc := service.NewService(ctx)
	list, err := svc.CollegeList(int64(userId))
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorCollegeList)
		return
	}
	res.ToResponse(list)
}

// @Summary 创建书单
// @Tags 书单
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param param body service.CollegeCreateRequest true "参数"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/college [post]
func (c College) Create(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	req := service.CollegeCreateRequest{}
	valid, errs := app.BindAndValid(ctx, &req, ctx.BindJSON)
	if !valid {
		res.ToErrorResponse(errcode.InValidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.NewService(ctx)
	err := svc.CreateCollege(req)
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorCreateCollege)
		return
	}
	res.ToResponse(gin.H{})
}

// @Summary 添加作品
// @Tags Chapter
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id path int true "书单ID"
// @Param workId path int true "作品ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/college/{id}/add/{workId} [post]
func (c College) Add(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	workId, err := app.ShouldParamConvertInt(ctx, "workId")
	if err != nil {
		res.ToErrorResponse(errcode.InValidParams.WithDetails(err.Error()))
		return
	}
	svc := service.NewService(ctx)
	err = svc.CollegeAddWork(ctx.GetInt64("id"), int64(workId))
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorCollegeAddWork)
		return
	}
	res.ToResponse(gin.H{})
}

// @Summary 修改书单信息
// @Tags 书单
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id path int true "书单ID"
// @Param param body service.CollegeUpdateRequest true "参数"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/college/{id} [put]
func (c College) Update(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	req := service.CollegeUpdateRequest{}
	valid, errs := app.BindAndValid(ctx, &req, ctx.BindJSON)
	if !valid {
		res.ToErrorResponse(errcode.InValidParams.WithDetails(errs.Errors()...))
	}
	svc := service.NewService(ctx)
	err := svc.UpdateCollege(ctx.GetInt64("id"), req)
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorCollegeUpdate)
		return
	}
	res.ToResponse(gin.H{})
}

// @Summary 删除
// @Tags 书单
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id path int true "书单ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/college/{id} [delete]
func (c College) Delete(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	svc := service.NewService(ctx)
	err := svc.DeleteCollege(ctx.GetInt64("id"))
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorCollegeDelete)
		return
	}
	res.ToResponse(gin.H{})
}
