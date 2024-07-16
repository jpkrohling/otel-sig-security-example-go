package telemetry

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func InitTelemetry() (func() error, error) {
	if otelEnabled() {
		exp, err := otlptracegrpc.New(context.Background())
		if err != nil {
			return nil, fmt.Errorf("failed to initialize traces exporter: %v", err)
		}
		tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exp))
		otel.SetTracerProvider(tp)
		return func() error {
			return tp.Shutdown(context.Background())
		}, nil
	}

	return func() error {
		return nil
	}, nil
}
