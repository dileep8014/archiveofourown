package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/shyptr/archiveofourown/pkg/errcode"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) Header(key, value string) {
	r.Ctx.Header(key, value)
}

func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list":  list,
		"total": totalRows,
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	if err.Err != nil {
		r.Ctx.Error(err)

		logger := r.Ctx.Value("logger").(zerolog.Logger)
		logger.Error().Caller(1).Err(err.Err).Send()
	}

	res := gin.H{"code": err.Code, "msg": err.Msg}
	details := err.Details
	if len(details) > 0 {
		res["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), res)
}
