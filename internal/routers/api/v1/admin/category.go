package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/shyptr/archiveofourown/internal/service"
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/errcode"
)

type Category struct {
}

func NewCategory() Category {
	return Category{}
}

func (c Category) Router(api gin.IRouter) {
	api.POST("/category", c.Create)
	api.PUT("/category/:id", c.Update)
	api.DELETE("/category/:id", c.Delete)
}

// @Summary 创建分类
// @Tags 分类
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param category body service.CategoryRequest true "分类"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/admin/category [post]
func (c Category) Create(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	req := service.CategoryRequest{}
	valid, errs := app.BindAndValid(ctx, &req, ctx.BindJSON)
	if !valid {
		res.ToErrorResponse(errcode.InValidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.NewService(ctx)
	err := svc.CreateCategory(req)
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorCreateCategoryFail)
		return
	}
	res.ToResponse(gin.H{})
}

// @Summary 修改指定分类信息
// @Tags 分类
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id path int true "分类ID"
// @Param category body service.CategoryRequest true "分类"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/admin/category/{id} [put]
func (c Category) Update(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	id, err := app.ShouldParamConvertInt(ctx, "id")
	if err != nil {
		res.ToErrorResponse(errcode.InValidParams.WithDetails(err.Error()))
		return
	}
	req := service.CategoryRequest{}
	valid, errs := app.BindAndValid(ctx, &req, ctx.BindJSON)
	if !valid {
		res.ToErrorResponse(errcode.InValidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.NewService(ctx)
	err = svc.UpdateCategory(int64(id), req)
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorUpdateCategoryFail)
		return
	}
	res.ToResponse(gin.H{})
}

// @Summary 删除分类
// @Tags 分类
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id path int true "分类ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/admin/category/{id} [delete]
func (c Category) Delete(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	id, err := app.ShouldParamConvertInt(ctx, "id")
	if err != nil {
		res.ToErrorResponse(errcode.InValidParams.WithDetails(err.Error()))
		return
	}
	svc := service.NewService(ctx)
	err = svc.DeleteCategory(int64(id))
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorDeleteCategoryFail)
		return
	}
	res.ToResponse(gin.H{})
}
