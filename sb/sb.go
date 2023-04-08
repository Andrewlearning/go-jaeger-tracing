package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"

	"go-jaeger-tracing/util"
)

func main() {
	util.InitLogger()
	defer util.DisposeLogger()

	cfg := config.Configuration{
		ServiceName: "service-b",
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

	http.HandleFunc("/process-request-b", func(w http.ResponseWriter, r *http.Request) {
		// fetch spanCtx from request header
		spanCtx, _ := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))

		// start a new span base on previous spanCtx
		span := opentracing.StartSpan("process-request-b", opentracing.ChildOf(spanCtx))
		defer span.Finish()
		traceID := span.Context().(jaeger.SpanContext).TraceID()
		parentID := span.Context().(jaeger.SpanContext).ParentID()
		spanID := span.Context().(jaeger.SpanContext).SpanID()
		log.Printf("[Service B] traceID: %v, parentID: %v, spanID:%v\n", traceID, parentID, spanID)
		fmt.Fprint(w, "Request processed successfully")

		util.SendHttp(cfg.ServiceName, "POST", "http://localhost:8082/process-request-c", span, tracer)
	})

	log.Fatal(http.ListenAndServe(":8081", nethttp.Middleware(tracer, http.DefaultServeMux)))
}
