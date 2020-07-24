package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/shyptr/archiveofourown/global"
	"github.com/uber/jaeger-client-go"
)

func Tracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx context.Context
		span := opentracing.SpanFromContext(c.Request.Context())
		if span != nil {
			span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(), global.Tracer,
				c.Request.URL.Path, opentracing.ChildOf(span.Context()))
		} else {
			span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(), global.Tracer,
				c.Request.URL.Path)
		}
		defer span.Finish()

		var traceID, spanID string
		var spanContext = span.Context()
		switch spanContext := spanContext.(type) {
		case jaeger.SpanContext:
			traceID = spanContext.TraceID().String()
			spanID = spanContext.SpanID().String()
		}
		c.Set("x-trace-id", traceID)
		c.Set("x-span-id", spanID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
