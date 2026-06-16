package reporter

import (
	"fmt"
	"os"
)

// HandleError is the centralized error-reporting function.
// It reports the error with sufficient context to the active backend (e.g., Sentry) or stderr.
func HandleError(err error, contextMsg string) {
	if err == nil {
		return
	}

	// Since we are writing CLI applications without Sentry right now, output to stderr
	// This centralizes error handling, while keeping the output readable for the CLI
	fmt.Fprintf(os.Stderr, "Error: %s: %v\n", contextMsg, err)
}
