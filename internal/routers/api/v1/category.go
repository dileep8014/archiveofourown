package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/shyptr/archiveofourown/internal/service"
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/errcode"
)

// Category: 作品分类
type Category struct {
}

func NewCategory() Category {
	return Category{}
}

func (c Category) Router(api *gin.RouterGroup) {
	api.GET("/category", c.List)
}

// @Summary 获取所有分类
// @Tags 分类
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {array} service.CategoryResponse "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/category [get]
func (c Category) List(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	svc := service.NewService(ctx)
	categories, err := svc.ListCategories()
	if err != nil {
		res.CheckErrorAndResponse(err, errcode.ErrorListCategoryFail)
		return
	}
	res.ToResponse(categories)
}
