package storage

import (
	"context"
	// "fmt"
	// "log"
	// "sync"


	"cloud.google.com/go/storage/internal"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	otelcodes "go.opentelemetry.io/otel/codes"
	// "go.opentelemetry.io/otel/propagation"
	// semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

const defaultTracerName = "cloud.google.com/go/storage"

func tracer() trace.Tracer {
	return otel.Tracer(defaultTracerName, trace.WithInstrumentationVersion(internal.Version))
}

func startSpan(ctx context.Context, name string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	// Maybe do something 
	// return otel.GetTracerProvider().Tracer(defaultTracerName).Start(ctx, name, opts...)
	return tracer().Start(ctx, name, opts...)
}

func endSpan(ctx context.Context, err error) {
	span := trace.SpanFromContext(ctx)
	if err != nil {
		span.SetStatus(otelcodes.Error, err.Error())
		span.RecordError(err)
	}
	span.End()
}

func getCommonTraceOptions() []trace.SpanStartOption {
	opts := []trace.SpanStartOption{
		trace.WithAttributes(
			attribute.String("gcp.client.service", "google.storage.v2.Storage"),
			attribute.String("gcp.client.repo", "googleapis/google-cloud-go/storage"),
			attribute.String("gcp.client.artifact", "com.google.cloud.google-cloud-storage"),
			attribute.String("foo", "bar"),
			attribute.Int64("inject-some-id", 888),
		),
	}
	return opts
}
