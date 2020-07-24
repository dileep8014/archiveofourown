package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/shyptr/archiveofourown/internal/service"
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/errcode"
	"strconv"
)

type Category struct {
}

func NewCategory() Category {
	return Category{}
}

func (this Category) Router(api *gin.RouterGroup) {
	api.POST("/category", this.Create)
	api.GET("/category/:id", this.List)
	api.PUT("/category/:id", this.Update)
	api.DELETE("/category/:id", this.Delete)
}

// @Summary 创建分类
// @Tags 分类
// @Accept json
// @Produce json
// @Param category body service.CategoryRequest true "分类"
// @Success 200 {object} model.Category "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/category [post]
func (this Category) Create(c *gin.Context) {
	logger := c.Value("logger").(zerolog.Logger)
	res := app.NewResponse(c)
	req := service.CategoryRequest{}
	valid, errs := app.BindAndValid(c, &req)
	if valid {
		logger.Error().AnErr("app.BindAndValid", errs).Send()
		res.ToErrorResponse(errcode.InValidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.NewService(c)
	defer svc.Finish()
	category, err := svc.CreateCategory(req)
	if err != nil {
		logger.Error().Err(err).Send()
		res.ToErrorResponse(errcode.ErrorCreateCategoryFail)
		return
	}
	res.ToResponse(category)
}

// @Summary 修改分类
// @Tags 分类
// @Accept json
// @Produce json
// @Param id path int true "分类ID"
// @Param category body service.CategoryRequest true "分类"
// @Success 200 {object} model.Category "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/category/{id} [put]
func (this Category) Update(c *gin.Context) {
	logger := c.Value("logger").(zerolog.Logger)
	res := app.NewResponse(c)
	req := service.CategoryRequest{}
	valid, errs := app.BindAndValid(c, &req)
	if valid {
		logger.Error().AnErr("app.BindAndValid", errs).Send()
		res.ToErrorResponse(errcode.InValidParams.WithDetails(errs.Errors()...))
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error().AnErr("app.ID", err).Send()
		res.ToErrorResponse(errcode.IDParseError)
		return
	}
	svc := service.NewService(c)
	defer svc.Finish()
	category, err := svc.UpdateCategory(id, req)
	if err != nil {
		logger.Error().Err(err).Send()
		res.ToErrorResponse(errcode.ErrorUpdateCategoryFail)
		return
	}
	res.ToResponse(category)
}

// @Summary 获取所有分类
// @Tags 分类
// @Accept json
// @Produce json
// @Success 200 {array} model.Category "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/category [get]
func (this Category) List(c *gin.Context) {
	logger := c.Value("logger").(zerolog.Logger)
	res := app.NewResponse(c)
	svc := service.NewService(c)
	defer svc.Finish()
	categories, err := svc.ListCategories()
	if err != nil {
		logger.Error().Err(err).Send()
		res.ToErrorResponse(errcode.ErrorListCategoryFail)
		return
	}
	res.ToResponse(categories)
}

// @Summary 删除分类
// @Tags 分类
// @Accept json
// @Produce json
// @Param id path int true "分类ID"
// @Success 200 {object} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/category/{id} [delete]
func (this Category) Delete(c *gin.Context) {
	logger := c.Value("logger").(zerolog.Logger)
	res := app.NewResponse(c)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error().AnErr("app.ID", err).Send()
		res.ToErrorResponse(errcode.IDParseError)
		return
	}
	svc := service.NewService(c)
	defer svc.Finish()
	err = svc.DeleteCategory(id)
	if err != nil {
		logger.Error().Err(err).Send()
		res.ToErrorResponse(errcode.ErrorDeleteCategoryFail)
		return
	}
	res.ToResponse(gin.H{})
}
