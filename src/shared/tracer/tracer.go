package tracer

import (
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

// Init returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func Init(service string) (opentracing.Tracer, io.Closer) {
	// Add Opentracing instrumentation
	defcfg := config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	cfg, err := defcfg.FromEnv()
	if err != nil {
		panic("Could not parse Jaeger env vars: " + err.Error())
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		panic("Could not initialize jaeger tracer: " + err.Error())
	}

	return tracer, closer
}
