package v1

import "github.com/gin-gonic/gin"

type Subscription struct {
}

func NewSubscription() Subscription {
	return Subscription{}
}

func (s Subscription) Router(api gin.IRouter) {
	api.POST("/subscription", s.Create)
	api.DELETE("/subscription", s.Delete)
}

// @Summary 添加订阅
// @Tags Subscription
// @Accept json
// @Produce json
// @Param param body service.SubscriptionRequest
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/subscription [post]
func (s Subscription) Create(ctx *gin.Context) {

}

// @Summary 添加订阅
// @Tags Subscription
// @Accept json
// @Produce json
// @Param param body service.SubscriptionRequest
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/subscription [delete]
func (s Subscription) Delete(ctx *gin.Context) {

}
