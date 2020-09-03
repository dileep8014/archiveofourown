package global

import (
	"github.com/opentracing/opentracing-go"
	"github.com/shyptr/archiveofourown/pkg/errwrap"
	"github.com/shyptr/archiveofourown/pkg/tracer"
)

var Tracer opentracing.Tracer

func SetupTracer() (err error) {
	defer errwrap.Wrap(&err, "init.setupTracer")

	jaegerTracer, _, err := tracer.NewJaegerTracer("pointer", "127.0.0.1:6831")
	if err != nil {
		return
	}
	Tracer = jaegerTracer
	return
}
