package v1

import "github.com/gin-gonic/gin"

type Identify struct {
}

func NewIdentify() Identify {
	return Identify{}
}

func (ident Identify) Router(api gin.IRouter) {
	api.GET("/identify/:path", ident.Check)
}

// @Summary 检测认证信息
// @Tags Identify
// @Accept json
// @Produce json
// @Param path path int true "认证路径"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/identify/{path} [get]
func (ident Identify) Check(ctx *gin.Context) {

}
