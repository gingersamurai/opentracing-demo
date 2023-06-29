package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"log"
	"time"
)

func main() {
	cfg := config.Configuration{
		ServiceName: "test-service",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "jaeger:6831",
		},
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		panic(err)
	}
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)
	for i := 0; i < 100; i++ {
		spanName := fmt.Sprintf("test-operation-%d", i)
		log.Println(spanName, "started")
		span, _ := opentracing.StartSpanFromContext(context.Background(), spanName)
		time.Sleep(1 * time.Second)
		span.Finish()
		log.Println(spanName, "finished")
	}
	time.Sleep(50 * time.Second)
}
