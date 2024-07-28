package panic

import (
	"log/slog"

	"github.com/crystalix007/panic/backend"
)

// SetLogger sets the logger for the backend.
func SetLogger(l *slog.Logger) {
	backend.SetLogger(l)
}

// GetLogger returns the logger for the backend.
func GetLogger() *slog.Logger {
	return backend.GetLogger()
}
