// Package panic provides logic to use the configured panic behaviour.
package panic

import (
	"fmt"

	"github.com/crystalix007/panic/backend"
)

// Panic will panic with the behaviour configured by the build tags.
func Panic(v interface{}) {
	backend.Panic(v)
}

// Panicf will panic with the given format and values, as configured by the
// build tags.
func Panicf(format string, v ...interface{}) {
	backend.Panic(fmt.Sprintf(format, v...))
}
