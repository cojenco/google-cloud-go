package storage

import (
	"context"

	"cloud.google.com/go/storage/internal"
	sharedtrace "cloud.google.com/go/internal/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	otelcodes "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

const defaultTracerName = "cloud.google.com/go/storage"

func tracer() trace.Tracer {
	return otel.Tracer(defaultTracerName, trace.WithInstrumentationVersion(internal.Version))
}

// isOpenCensusTracingEnabled returns a boolean. If true, we use the shared internal/trace
// module for Opencensus support.
// To be removed upon OpenCensus planned sunset in Dec 2024.
func isOpenCensusTracingEnabled() bool {
	return sharedtrace.IsOpenCensusTracingEnabled()
}

// startSpan accepts SpanStartOption and is used to replace internal/trace/StartSpan.
func startSpan(ctx context.Context, name string, opts ...trace.SpanStartOption) (context.Context) {
	// We can append commonTraceOptions here or preprocess additional info...

	if isOpenCensusTracingEnabled() {
		// Use the shared internal/trace module for existing OpenCencus support.
		ctx = sharedtrace.StartSpan(ctx, name)
	} else {
		ctx, _ = tracer().Start(ctx, name, opts...)
	}
	return ctx
}

// endSpan is used to replace internal/trace/EndSpan.
func endSpan(ctx context.Context, err error) {
	if isOpenCensusTracingEnabled() {
		// Use the shared internal/trace module for existing OpenCencus support.
		sharedtrace.EndSpan(ctx, err)
	} else {
		span := trace.SpanFromContext(ctx)
		if err != nil {
			span.SetStatus(otelcodes.Error, err.Error())
			span.RecordError(err)
		}
		span.End()
	}
}

// getCommonTraceOptions() compiles the common span attributes we want to append.
// Similar helper functions can be created for such operations as reads and writes.
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