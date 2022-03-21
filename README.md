# loggo
Better fmt.Println() for Go

Inspired by [@wizo06/logger](https://www.npmjs.com/package/@wizo06/logger)

# Installation

```console
$ go get github.com/wizo06/loggo
```

# Quickstart

Filename `loggo_test.go` with content:

```go
package loggo

import "github.com/wizo06/loggo"

func main() {
	loggo.Info("hello world %s %s", "foo", "bar")
	loggo.Success("hello world")
	loggo.Debug("hello world")
	loggo.Warning("hello world")
	loggo.Error("hello world")
}
```

Output:

```console
[2022.3.21|2:19:26|UTC-4] [loggo_test.go:6] [INFO] hello world foo bar
[2022.3.21|2:19:26|UTC-4] [loggo_test.go:7] [SUCCESS] hello world
[2022.3.21|2:19:26|UTC-4] [loggo_test.go:8] [DEBUG] hello world
[2022.3.21|2:19:26|UTC-4] [loggo_test.go:9] [WARNING] hello world
[2022.3.21|2:19:26|UTC-4] [loggo_test.go:10] [ERROR] hello world
```

# Force refresh in pkg.go.dev

```console
$ curl proxy.golang.org/github.com/wizo06/loggo/@latest
```