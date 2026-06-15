package reporter

import (
	"log/slog"
)

// HandleError provides a centralized way to report unexpected errors.
func HandleError(msg string, err error, attrs ...any) {
	if err != nil {
		args := append([]any{"error", err}, attrs...)
		slog.Error(msg, args...)
	} else {
		slog.Error(msg, attrs...)
	}
}
