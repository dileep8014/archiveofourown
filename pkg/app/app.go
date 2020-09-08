package app

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
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

// for which has none data returns.
func (r *Response) ToSuccessResponse() {
	res := gin.H{"code": errcode.Success.Code, "msg": errcode.Success.Msg}
	r.Ctx.JSON(errcode.Success.StatusCode(), res)
}

func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, gin.H{"code": errcode.Success.Code, "data": data})
}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list":  list,
		"total": totalRows,
	})
}

// if err == errocde.*Error, then use err, otherwise use codeErr.WithErr(err)
func (r *Response) CheckErrorAndResponse(err error, codeErr errcode.Error) {
	if errcode.IsError(err) {
		r.ToErrorResponse(err.(errcode.Error))
		return
	}
	r.ToErrorResponse(codeErr.WithError(err), 1)
}

func (r *Response) ToErrorResponse(err errcode.Error, skip ...int) {
	span := r.Ctx.Value("span").(opentracing.Span)
	span.SetTag("gin.error", err.ErrorWithDetails())

	if err.Err != nil {
		var k = 1
		if len(skip) > 0 {
			k += skip[0]
		}
		span.SetTag("sys.error", err.Err)
		ext.Error.Set(span, true)

		logger := r.Ctx.Value("logger").(zerolog.Logger)
		logger.Error().Caller(k).Err(err.Err).Send()
	}

	res := gin.H{"code": err.Code, "msg": err.Msg}
	details := err.Details
	if len(details) > 0 {
		res["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), res)
}
