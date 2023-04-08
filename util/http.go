package util

import (
	"fmt"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
	"io/ioutil"
	"log"
	"net/http"
)

func SendHttp(serviceName, method, url string, span opentracing.Span, tracer opentracing.Tracer) {
	httpClient := &http.Client{
		Transport: &nethttp.Transport{},
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatalf("Failed to create HTTP request: %v", err)
	}

	if err := tracer.Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header)); err != nil {
		log.Printf("Failed to inject span context into HTTP headers: %v", err)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to send HTTP request: %v", err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	defer res.Body.Close()

	// Print out the result
	fmt.Printf("[%v] Response status: %s, res.Body:%v\n", serviceName, res.Status, string(body))
}
