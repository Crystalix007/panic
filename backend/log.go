//go:build GITHUB_COM_CRYSTALIX007_PANIC_LOG || (prod && !GITHUB_COM_CRYSTALIX007_PANIC_CUSTOM && !GITHUB_COM_CRYSTALIX007_PANIC_SILENT && !GITHUB_COM_CRYSTALIX007_PANIC_PANIC)

package backend

import "log/slog"

var logger *slog.Logger = slog.Default()

// Panic will log the given value using the configured logger.
func Panic(v any) {
	logger.Warn("caught panic", slog.Any("value", v))
}

// SetLogger sets the logger to use for logging panics.
func SetLogger(l *slog.Logger) {
	logger = l
}

// GetLogger returns the logger used for logging panics.
func GetLogger() *slog.Logger {
	return logger
}
