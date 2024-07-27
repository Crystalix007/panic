// Package panic provides logic to panic when in development mode.
package panic

import "fmt"

// Panicf will panic with the given format and values if in development mode.
func Panicf(format string, v ...interface{}) {
	Panic(fmt.Sprintf(format, v...))
}
