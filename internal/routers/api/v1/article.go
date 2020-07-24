package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/shyptr/archiveofourown/pkg/app"
	"github.com/shyptr/archiveofourown/pkg/errcode"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

func (this Article) Router(api *gin.RouterGroup) {
	api.POST("/articles", this.Create)
	api.DELETE("/articles/:id", this.Delete)
	api.PUT("/articles/:id", this.Update)
	api.GET("/articles/:id", this.Get)
	api.GET("/articles", this.List)
}

// @Summary 获取文章信息
// @Tags 文章
// @Accept  json
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [get]
func (this Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}

// @Summary 获取文章列表
// @Tags 文章
// @Accept  json
// @Produce json
// @Param query query string true "查询内容" maxlength(100)
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [get]
func (this Article) List(c *gin.Context) {}

// @Summary 创建文章
// @Tags 文章
// @Accept  json
// @Produce json
// @Param title body string true "标题" maxlength(100)
// @Param subTitle body string true "简述" maxlength(100)
// @Param markID body int false "书签ID"
// @Param language body int true "语言" Enums(1,2) default(1)
// @Param chapterNums body int false "计划章节数"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [post]
func (this Article) Create(c *gin.Context) {}

// @Summary 修改文章信息
// @Tags 文章
// @Accept  json
// @Produce json
// @Param id path int true "文章ID"
// @Param title body string true "标题" maxlength(100)
// @Param subTitle body string true "简述" maxlength(100)
// @Param markID body int false "书签ID"
// @Param language body int true "语言" Enums(1,2) default(1)
// @Param chapterNums body int false "计划章节数"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [put]
func (this Article) Update(c *gin.Context) {}

// @Summary 删除文章
// @Tags 文章
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [delete]
func (this Article) Delete(c *gin.Context) {}
