package v1

import "github.com/gin-gonic/gin"

type News struct {
}

func NewNews() News {
	return News{}
}

func (n News) Router(api gin.IRouter) {
	api.GET("/news", n.List)
	api.GET("/news/:id", n.Get)
}

// @Summary 公告列表
// @Tags News
// @Accept json
// @Produce json
// @Param pageSize query int true "页大小"
// @Param page query int true "当前页"
// @Success 200 {object} service.NewsPageResponse "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/news/list [get]
func (n News) List(ctx *gin.Context) {

}

// @Summary 公告内容
// @Tags News
// @Accept json
// @Produce json
// @Param id path int true "公告ID"
// @Success 200 {object} service.NewsResponse "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/news/{id} [get]
func (n News) Get(ctx *gin.Context) {

}
