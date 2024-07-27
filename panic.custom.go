//go:build GITHUB_COM_CRYSTALIX007_PANIC_CUSTOM

package panic

var panicImpl func(v any) = defaultPanicImpl

// Panic will panic with the custom implementation if in development mode.
func Panic(v any) {
	panicImpl(v)
}

// SetPanicImpl sets the custom panic implementation.
//
// WARNING: The implementation must terminate, and not infinitely recurse.
//
//	Calling Panic within the implementation is likely to cause a stack
//	overflow.
func SetPanicImpl(impl func(v any)) {
	panicImpl = impl
}

func defaultPanicImpl(v any) {
	panic(v)
}
