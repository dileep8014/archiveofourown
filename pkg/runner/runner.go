package runner

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog"
	"github.com/shyptr/archiveofourown/global"
	"github.com/shyptr/archiveofourown/pkg/logger"
	"github.com/shyptr/sqlex"
	"github.com/uber/jaeger-client-go"
	"strings"
)

type Runner struct {
	*sql.Tx
	Logger zerolog.Logger
	ctx    context.Context
	span   opentracing.Span
}

func NewRunner(c *gin.Context) *Runner {
	var ctx context.Context
	span := opentracing.SpanFromContext(c.Request.Context())
	if span != nil {
		span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(), global.Tracer,
			"sql", opentracing.ChildOf(span.Context()))
	} else {
		span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(), global.Tracer,
			"sql")
	}

	logger := logger.Get()

	var spanContext = span.Context()
	switch spanContext := spanContext.(type) {
	case jaeger.SpanContext:
		logger = logger.With().Str("trace_id", spanContext.TraceID().String()).Logger().
			With().Str("span_id", spanContext.SpanID().String()).Logger()
	}

	return &Runner{
		Tx:     c.Value("tx").(*sql.Tx),
		Logger: logger,
		span:   span,
		ctx:    ctx,
	}
}

func (r *Runner) Close() {
	r.span.Finish()
	logger.Put(r.Logger)
}

func (r *Runner) QueryRow(query string, args ...interface{}) sqlex.RowScanner {
	r.Logger.Info().Str("sql", fmt.Sprintf(strings.ReplaceAll(query, "?", "%v"), args...)).Send()
	return r.Tx.QueryRow(query, args...)
}

func (r *Runner) Exec(query string, args ...interface{}) (sql.Result, error) {
	r.Logger.Info().Str("sql", fmt.Sprintf(strings.ReplaceAll(query, "?", "%v"), args...)).Send()
	return r.Tx.Exec(query, args...)
}

func (r *Runner) Query(query string, args ...interface{}) (*sql.Rows, error) {
	r.Logger.Info().Str("sql", fmt.Sprintf(strings.ReplaceAll(query, "?", "%v"), args...)).Send()
	return r.Tx.Query(query, args...)
}
