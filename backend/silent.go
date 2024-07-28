//go:build GITHUB_COM_CRYSTALIX007_PANIC_SILENT

package backend

// Panic will panic if in development mode.
func Panic(v any) {
	// Do nothing in production mode.
}
