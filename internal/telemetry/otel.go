package telemetry

import (
	"os"
	"strings"
)

func otelEnabled() bool {
	for _, v := range os.Environ() {
		if strings.HasPrefix(v, "OTEL_") {
			return true
		}
	}
	return false
}
