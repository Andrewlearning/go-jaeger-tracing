package main

import (
	"log"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"

	"go-jaeger-tracing/util"
)

func main() {
	util.InitLogger()
	defer util.DisposeLogger()

	cfg := config.Configuration{
		ServiceName: "service-a",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		log.Fatalf("Failed to initialize Jaeger tracer: %v", err)
	}
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)

	// start a new span in the head of the flow
	span := tracer.StartSpan("send-request")
	defer span.Finish()
	spanCtx := span.Context()
	traceID := spanCtx.(jaeger.SpanContext).TraceID()
	parentID := spanCtx.(jaeger.SpanContext).ParentID()
	spanID := spanCtx.(jaeger.SpanContext).SpanID()
	log.Printf("[Service A] traceID: %v, parentID: %v, spanID:%v\n", traceID, parentID, spanID)

	util.SendHttp("POST", "http://localhost:8080/process-request-b", span, tracer)
}
