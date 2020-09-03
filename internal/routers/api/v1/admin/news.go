package admin

import "github.com/gin-gonic/gin"

type News struct {
}

func NewNews() News {
	return News{}
}

func (n News) Router(api gin.IRouter) {
	api.POST("/news", n.Create)
	api.PUT("/news/:id", n.Update)
	api.DELETE("/news/:id", n.Delete)
}

// @Summary 添加公告
// @Tags News
// @Accept json
// @Produce json
// @Param param body service.NewsRequest
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/news [post]
func (n News) Create(ctx *gin.Context) {

}

// @Summary 修改公告
// @Tags News
// @Accept json
// @Produce json
// @Param id path int true "公告ID"
// @Param param body service.NewsUpdateRequest
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/news/{id} [put]
func (n News) Update(ctx *gin.Context) {

}

// @Summary 删除公告
// @Tags News
// @Accept json
// @Produce json
// @Param id path int true "公告ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/news/{id} [delete]
func (n News) Delete(ctx *gin.Context) {

}
