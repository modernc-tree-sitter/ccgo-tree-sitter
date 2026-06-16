package reporter

import (
	"fmt"
	"log/slog"
	"os"
)

// HandleError is the centralized error-reporting function.
// It reports the error with sufficient context to the active backend (e.g., Sentry) or stderr.
func HandleError(err error, contextMsg string) {
	if err == nil {
		return
	}

	// Log with sufficient context
	slog.Error("error occurred", "context", contextMsg, "error", err)

	// Since we are writing CLI applications, output to stderr as a fallback
	// This mirrors the prior scattered error handling behavior while centralizing it.
	fmt.Fprintf(os.Stderr, "Error: %s: %v\n", contextMsg, err)
}
