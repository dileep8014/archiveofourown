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
// @Tags 公告
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param param body service.NewsRequest true "请求参数"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/news [post]
func (n News) Create(ctx *gin.Context) {

}

// @Summary 修改公告
// @Tags 公告
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id path int true "公告ID"
// @Param param body service.NewsUpdateRequest true "请求参数"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/news/{id} [put]
func (n News) Update(ctx *gin.Context) {

}

// @Summary 删除公告
// @Tags 公告
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id path int true "公告ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/news/{id} [delete]
func (n News) Delete(ctx *gin.Context) {

}
