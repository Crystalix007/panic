# Panic

A library that allows different panic behaviours to be configured at runtime.

The intent is to allow stricter assertions in development, that may be relaxed
when running in production, although the reverse may also be configured.

## Usage

```go
package main

import "github.com/crystalix007/panic"

func main() {
	panic.Panic("test")
}
```

Then, this can be run / built with `-tag dev`:

```sh
$ go run -tag dev ./main.go
panic: test

goroutine 1 [running]:
github.com/crystalix007/panic.Panic(...)
        /Users/michaelkuc6/Documents/coding/panic/panic.dev.go:7
main.main()
        /private/tmp/testing/main.go:10 +0x30
exit status 2
```

or `-tag prod`:

```sh
$ go run -tags prod ./main.go
```

## Build flags

|                     Flag                         |                   Description                   |
|--------------------------------------------------|-------------------------------------------------|
| default or `GITHUB_COM_CRYSTALIX007_PANIC_PANIC` |     Panics and breaks execution on `panic`.     |
|   `prod` or `GITHUB_COM_CRYSTALIX007_PANIC_LOG`  |       Makes panics print a `slog` warning.      |
|      `GITHUB_COM_CRYSTALIX007_PANIC_SILENT`      |             Silently drops panics.              |
|      `GITHUB_COM_CRYSTALIX007_PANIC_CUSTOM`      | Allows setting a custom `panic` implementation. |

Using environment tags, like:

- `dev`
- `prod`

will also be sufficient, as `dev` will invoke the default behaviour, and `prod`
will enable `slog` logging.
