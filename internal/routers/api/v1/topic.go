package v1

import "github.com/gin-gonic/gin"

type Topic struct {
}

func NewTopic() Topic {
	return Topic{}
}

func (t Topic) Router(api gin.IRouter) {
	api.GET("/topic", t.List)
	api.POST("/topic", t.Create)
	api.PUT("/topic/:id", t.Update)
	api.DELETE("/topic/:id", t.Delete)
}

// @Summary 专题列表
// @Tags Topic
// @Accept json
// @Produce json
// @Param pageSize query int true "页大小"
// @Param page query int true "当前页"
// @Success 200 {object} service.TopicPageResponse "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/topic [get]
func (t Topic) List(ctx *gin.Context) {

}

// @Summary 创建专题
// @Tags Topic
// @Accept json
// @Produce json
// @Param param body service.TopicRequest
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/topic [post]
func (t Topic) Create(ctx *gin.Context) {

}

// @Summary 修改专题
// @Tags Topic
// @Accept json
// @Produce json
// @Param id path int true "专题ID"
// @Param param body service.TopicUpdateRequest
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/topic/{id} [put]
func (t Topic) Update(ctx *gin.Context) {

}

// @Summary 修改专题
// @Tags Topic
// @Accept json
// @Produce json
// @Param id path int true "专题ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/topic/{id} [delete]
func (t Topic) Delete(ctx *gin.Context) {

}
