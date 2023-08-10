package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"github.com/uber/jaeger-client-go/config"
)

func main() {
	viper.AutomaticEnv()
	cfg := config.Configuration{
		ServiceName: "test-service",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:          true,
			CollectorEndpoint: viper.GetString("JAEGER_ENDPOINT"),
		},
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = closer.Close()
	}()
	opentracing.SetGlobalTracer(tracer)

	num := 1
	for {
		spanName := fmt.Sprintf("test-operation-%d-%s", num, time.Now().String())
		span, _ := opentracing.StartSpanFromContext(context.Background(), spanName)
		time.Sleep(viper.GetDuration("INTERVAL"))
		span.Finish()
		log.Println(spanName, "finished")
		num++
	}
}
