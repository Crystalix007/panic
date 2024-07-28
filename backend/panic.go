//go:build !GITHUB_COM_CRYSTALIX007_PANIC_CUSTOM && !GITHUB_COM_CRYSTALIX007_PANIC_SILENT && !GITHUB_COM_CRYSTALIX007_PANIC_LOG && !prod

// The default backend, if no build tags are provided.

package backend

// Panic will panic with the given value if in development mode.
func Panic(v any) {
	panic(v)
}
