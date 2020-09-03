package v1

import "github.com/gin-gonic/gin"

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Router(api gin.IRouter) {
	api.GET("/tag/:name", t.List)
}

// @Summary 查询匹配标签
// @Tags Tag
// @Accept json
// @Produce json
// @Param name path string true "标签名"
// @Success 200 {array} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tag/{name} [get]
func (t Tag) List(ctx *gin.Context) {

}
