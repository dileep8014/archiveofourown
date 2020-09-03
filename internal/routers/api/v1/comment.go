package v1

import "github.com/gin-gonic/gin"

type Comment struct {
}

func NewComment() Comment {
	return Comment{}
}

func (c Comment) Router(api gin.IRouter) {
	api.POST("/comment", c.Create)
	api.GET("/comment/:objType/:objID", c.List)
	api.DELETE("/comment/:id", c.Delete)
}

// @Summary 添加评论
// @Tags Comment
// @Accept json
// @Produce json
// @Param param body service.CommentRequest
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/comment [post]
func (c Comment) Create(ctx *gin.Context) {

}

// @Summary 获取评论列表
// @Tags Comment
// @Accept json
// @Produce json
// @Param objType path int true "评论对象类型,0-作品,1-章节,2-选段,3-公告" Enums(0,1,2,3)
// @Param objID path int true "评论对象ID"
// @Param pageSize query int true "页大小"
// @Param page query int true "当前页"
// @Success 200 {object} service.CommentPageResponse "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/comment/{objType}/{objID} [get]
func (c Comment) List(ctx *gin.Context) {

}

// @Summary 删除评论
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path int true "评论ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/comment/{id} [delete]
func (c Comment) Delete(ctx *gin.Context) {

}
