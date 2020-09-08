package v1

import "github.com/gin-gonic/gin"

type Reply struct {
}

func NewReply() Reply {
	return Reply{}
}

func (r Reply) Router(api gin.IRouter) {
	api.GET("/reply/:commentID", r.List)
	api.POST("/reply", r.Create)
	api.DELETE("/reply/:id", r.Delete)
}

// @Summary 回复列表
// @Tags 回复
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param commentID path int true "评论ID"
// @Success 200 {array} service.ReplyResponse "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/reply/{commentID} [get]
func (r Reply) List(ctx *gin.Context) {

}

// @Summary 添加回复
// @Tags 回复
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param param body service.ReplyRequest true "请求参数"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/reply [post]
func (r Reply) Create(ctx *gin.Context) {

}

// @Summary 回复列表
// @Tags 回复
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id path int true "回复ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/reply/{id} [delete]
func (r Reply) Delete(ctx *gin.Context) {

}
