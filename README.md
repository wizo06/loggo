# loggo
Better fmt.Println() for Go

Inspired by [@wizo06/logger](https://www.npmjs.com/package/@wizo06/logger)

# Installation

```console
$ go get github.com/wizo06/loggo
```

# Quick start

```go
package main

import "github.com/wizo06/loggo"

type Foo struct {
	Bar string
	Baz int
}

func main() {
	log := loggo.New(loggo.Config{
		PrintHostname:               true,
		PrintUNIXTimestamp:          true,
		PrintHumanReadableTimestamp: true,
		StackDepth:                  3,
		PrintFileName:               true,
		PrintFunctionName:           true,
		PrintLineNumber:             true,
		PrintLogLevel:               true,
	})

	log.Info("hello world")
	log.Success("hello world")
	log.Debug("hello world")
	log.Warn("hello world")
	log.Error("hello world")
	log.Info("hello", "world", "and", "friends")
	log.Info(Foo{})
}
```

Output:

```console
[hanabira] [1660118384|2022.8.10|3:59:44|UTC-4] [main.go:main:22] [info] [hello world]
[hanabira] [1660118384|2022.8.10|3:59:44|UTC-4] [main.go:main:23] [success] [hello world]
[hanabira] [1660118384|2022.8.10|3:59:44|UTC-4] [main.go:main:24] [debug] [hello world]
[hanabira] [1660118384|2022.8.10|3:59:44|UTC-4] [main.go:main:25] [warn] [hello world]
[hanabira] [1660118384|2022.8.10|3:59:44|UTC-4] [main.go:main:26] [error] [hello world]
[hanabira] [1660118384|2022.8.10|3:59:44|UTC-4] [main.go:main:27] [info] [hello world and friends]
[hanabira] [1660118384|2022.8.10|3:59:44|UTC-4] [main.go:main:28] [info] [{Bar: Baz:0}]
```

# Full Output Format

```
[<hostname>] [<unix timestamp>|<human readable timestamp>] [<file name>:<function_name>:<line number>] [<log level>]
```

# Available styling

```go
loggo.BLACK
loggo.RED
loggo.GREEN
loggo.YELLOW
loggo.BLUE
loggo.MAGENTA
loggo.CYAN
loggo.WHITE

loggo.RESET
loggo.BRIGHT
loggo.DIM
loggo.UNDERSCORE
loggo.BLINK
loggo.REVERSE
loggo.HIDDEN
```

## Example

```go
log.Info("Downloading", UNDERSCORE, "movie.mp4", RESET)
```

# Force refresh in pkg.go.dev

```console
$ TAG="v1.0.0"
$ git tag $TAG
$ git push origin $TAG
$ GOPROXY=proxy.golang.org go list -m github.com/wizo06/loggo@$TAG
$ curl proxy.golang.org/github.com/wizo06/loggo/@latest
```