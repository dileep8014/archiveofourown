package v1

import "github.com/gin-gonic/gin"

type Subsection struct {
}

func NewSubsection() Subsection {
	return Subsection{}
}

func (s Subsection) Router(api gin.IRouter) {
	api.GET("/subsection/:workID", s.List)
	api.POST("/subsection", s.Create)
	api.PUT("/subsection/:id", s.Update)
	api.DELETE("/subsection/:id", s.Delete)
}

// @Summary 分卷列表
// @Tags 分卷
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param workID path int true "作品ID"
// @Success 200 {array} service.SubsectionResponse "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/subsection/{workID} [get]
func (s Subsection) List(ctx *gin.Context) {

}

// @Summary 分卷列表
// @Tags 分卷
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param param body service.SubsectionRequest true "请求参数"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/subsection [post]
func (s Subsection) Create(ctx *gin.Context) {

}

// @Summary 分卷列表
// @Tags 分卷
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id path int true "分卷ID"
// @Param param body service.SubsectionRequest true "请求参数"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/subsection/{id} [put]
func (s Subsection) Update(ctx *gin.Context) {

}

// @Summary 分卷列表
// @Tags 分卷
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id path int true "分卷ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/subsection/{id} [delete]
func (s Subsection) Delete(ctx *gin.Context) {

}
