package app

import (
	"github.com/gin-gonic/gin"
	"github.com/shyptr/archiveofourown/pkg/errcode"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

// Pager example
type Pager struct {
	Page      int `json:"page" example:"1"`
	PageSize  int `json:"pageSize" example:"10"`
	TotalRows int `json:"totalRows" example:"100"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	res := gin.H{"code": err.Code, "msg": err.Msg}
	details := err.Details
	if len(details) > 0 {
		res["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), res)
}
