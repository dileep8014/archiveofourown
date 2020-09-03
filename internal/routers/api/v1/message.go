package v1

import "github.com/gin-gonic/gin"

type Message struct {
}

func NewMessage() Message {
	return Message{}
}

func (m Message) Router(api gin.IRouter) {
	api.GET("/message/unread")
	api.GET("/message/all/:type", m.List)
	api.POST("/message", m.Send)
}

// @Summary 未读私信数
// @Tags Message
// @Accept json
// @Produce json
// @Success 200 {int} int "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/message/unread [get]
func (m Message) Unread(ctx *gin.Context) {

}

// @Summary 查看私信
// @Tags Message
// @Accept json
// @Produce json
// @Param type path int true "私信类型" Enums(0,1)
// @Param pageSize query int true "页大小"
// @Param page query int true "当前页"
// @Success 200 {object} service.MessagePageResponse "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/message/{type} [get]
func (m Message) List(ctx *gin.Context) {

}

// @Summary 发送私信
// @Tags Message
// @Accept json
// @Produce json
// @Param param body service.MessageRequest true "参数"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/message [post]
func (m Message) Send(ctx *gin.Context) {

}
