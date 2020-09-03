package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog"
	"github.com/shyptr/archiveofourown/global"
	"github.com/uber/jaeger-client-go"
	"time"
)

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		span := c.Value("span").(opentracing.Span)
		spanContext := span.Context().(jaeger.SpanContext)
		log := global.Logger.With().Str("trace_id", spanContext.TraceID().String()).Logger().
			With().Str("span_id", spanContext.SpanID().String()).Logger()
		log = log.Hook(AccessLogHook(c, span))
		c.Set("logger", log)
		beginTime := time.Now()
		c.Next()
		log.Info().Int("status", c.Writer.Status()).TimeDiff("takeUp", time.Now(), beginTime).
			Str("ip", c.ClientIP()).Str("method", c.Request.Method).Str("path", c.Request.RequestURI).Send()
	}
}

func AccessLogHook(c *gin.Context, span opentracing.Span) zerolog.HookFunc {
	return func(e *zerolog.Event, level zerolog.Level, message string) {
		if level >= zerolog.ErrorLevel {
			span.SetTag("gin.method", c.Request.Method)
			span.SetTag("gin.error", c.Errors.String())
		}
	}
}
