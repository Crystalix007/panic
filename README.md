# Panic

A library that allows different panic behaviours to be configured at runtime.

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
| `dev` or `GITHUB_COM_CRYSTALIX007_PANIC_ENABLED` | Panics and breaks execution on `panic`.         |
|      `GITHUB_COM_CRYSTALIX007_PANIC_CUSTOM`      | Allows setting a custom `panic` implementation. |

Note that the `prod` tag actually indicates production mode by the _absence_ of
the `dev` tag. You could equally not tag it at all.
