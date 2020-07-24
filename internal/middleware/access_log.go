package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/shyptr/archiveofourown/pkg/logger"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.Get().With().Str("trace_id", c.Value("x-trace-id").(string)).Logger().
			With().Str("span_id", c.Value("x-span-id").(string)).Logger()
		c.Set("logger", log)
		defer logger.Put(log)
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter
		beginTime := time.Now()
		c.Next()
		log.Info().Int("status", bodyWriter.Status()).TimeDiff("takeUp", time.Now(), beginTime).
			Str("ip", c.ClientIP()).Str("method", c.Request.Method).Str("path", c.Request.RequestURI).Send()
	}
}
