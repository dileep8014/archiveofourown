package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/shyptr/archiveofourown/internal/service"
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/errcode"
)

// College: 书单
type College struct {
}

func NewCollege() College {
	return College{}
}

func (c College) Router(api gin.IRouter) {
	// 获取书单内作品列表
	api.GET("/college/:id/works", c.Works)
	// 获取书单列表
	api.GET("/college", c.List)
	// 创建书单
	api.POST("/college", c.Create)
	// 添加作品
	api.POST("/college/:id/add", c.Add)
	// 修改书单信息
	api.PUT("/college/:id", c.Update)
	// 删除书单
	api.DELETE("/college/:id", c.Delete)
}

// @Summary 获取书单内作品列表
// @Tags College
// @Accept json
// @Produce json
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
		res.ToErrorResponse(errcode.ErrorCollegeWorks.WithError(err))
		return
	}
	res.ToResponse(works)
}

// @Summary 查询书单列表
// @Tags College
// @Accept json
// @Produce json
// @Param pageSize query int true "页大小"
// @Param page query int true "当前页"
// @Success 200 {object} service.CollegePageResponse "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/college [get]
func (c College) List(ctx *gin.Context) {

}

// @Summary 创建书单
// @Tags College
// @Accept json
// @Produce json
// @Param param body service.CollegeCreateRequest true "参数"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/college [post]
func (c College) Create(ctx *gin.Context) {

}

// @Summary 添加作品
// @Tags Chapter
// @Accept json
// @Produce json
// @Param id path int true "书单ID"
// @Param workID path int true "作品ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/college/add/{id} [post]
func (c College) Add(ctx *gin.Context) {

}

// @Summary 修改书单信息
// @Tags College
// @Accept json
// @Produce json
// @Param id path int true "书单ID"
// @Param param body service.CollegeUpdateRequest true "参数"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/college/{id} [put]
func (c College) Update(ctx *gin.Context) {

}

// @Summary 删除
// @Tags College
// @Accept json
// @Produce json
// @Param id path int true "书单ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/college/{id} [delete]
func (c College) Delete(ctx *gin.Context) {

}
