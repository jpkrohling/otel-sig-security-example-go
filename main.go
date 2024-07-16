package main

import (
	"context"
	"fmt"

	"github.com/jpkrohling/otel-sig-security-example-go/internal/telemetry"
	"go.opentelemetry.io/otel"
)

func main() {
	// added some telemetry just to have a few dependencies on the go.mod, and therefore, a go.sum file
	ctx := context.Background()
	close := telemetry.InitTelemetry()
	defer close()

	tr := otel.Tracer("github.com/jpkrohling/otel-sig-security-example-go")

	_, span := tr.Start(ctx, "main")
	defer span.End()

	fmt.Println("Hello, World!")
}
