//go:build dev || GITHUB_COM_CRYSTALIX007_PANIC_ENABLED

package panic

// Panic will panic with the given value if in development mode.
func Panic(v interface{}) {
	panic(v)
}
