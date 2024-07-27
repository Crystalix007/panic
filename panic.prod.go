//go:build !dev && !GITHUB_COM_CRYSTALIX007_PANIC_ENABLED

package panic

// Panic will panic if in development mode.
func Panic(v interface{}) {
	// Do nothing in production mode.
}
