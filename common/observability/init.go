package observability

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

// InitTracer creates a new trace provider instance and registers it as global trace provider.
func InitTracer(ctx context.Context) error {
	// First, create and install OTLP trace exporter, and configure it to export.
	// In this example, we will export to a collector running on the host machine.
	// Exporter is where traces are sent, usually we use OTLP(OpenTelemetry Protocol) to send traces to collector.
	exporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithEndpoint("localhost:4317"),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		return err
	}

	tp := tracesdk.NewTracerProvider(
		// Batcher is where traces are buffered before being sent to exporter
		// which can be used to reduce the number of calls to exporter.
		tracesdk.WithBatcher(exporter),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			// Service name is used to identify the service that generates the trace.
			// Usually we use package name as the service name.
			// In this example, we use "trace-example" as the service name.
			semconv.ServiceNameKey.String("trace-example"),
		)),
	)

	// we need to set the tracer provider to the global provider,
	// so we can use otel.Trace directly in other packages,
	// which is more convenient.
	otel.SetTracerProvider(tp)

	return nil
}
