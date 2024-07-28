//go:build !(GITHUB_COM_CRYSTALIX007_PANIC_LOG || prod)

package backend

import "log/slog"

// Fallback implementations of the options for the log-specific backend, when
// the log backend is not enabled.
//
// These implementations do nothing.

// SetLogger does nothing in the fallback implementation.
func SetLogger(l *slog.Logger) {
	// Do nothing.
}

// GetLogger does nothing in the fallback implementation.
func GetLogger() *slog.Logger {
	return nil
}
