//go:build !dev && !GITHUB_COM_CRYSTALIX007_PANIC_ENABLED && !GITHUB_COM_CRYSTALIX007_PANIC_CUSTOM

package panic

// Panic will panic if in development mode.
func Panic(v any) {
	// Do nothing in production mode.
}
