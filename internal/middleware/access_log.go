package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/shyptr/archiveofourown/global"
	"github.com/uber/jaeger-client-go"
	"time"
)

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		span := c.Value("span").(opentracing.Span)
		span.SetTag("gin.method", c.Request.Method)
		span.SetTag("gin.url", c.Request.RequestURI)
		span.SetTag("gin.status", c.Writer.Status())

		spanContext := span.Context().(jaeger.SpanContext)
		logger := global.Logger.With().Str("trace_id", spanContext.TraceID().String()).Logger().
			With().Str("span_id", spanContext.SpanID().String()).Logger()
		c.Set("logger", logger)

		beginTime := time.Now()
		c.Next()
		logger.Info().Int("status", c.Writer.Status()).TimeDiff("takeUp", time.Now(), beginTime).
			Str("ip", c.ClientIP()).Str("method", c.Request.Method).Str("path", c.Request.RequestURI).Send()
	}
}
