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

type MyStruct struct {
	ID   string
	Name string
}

func main() {
	loggo.Info("hello world %s %+v", "foo", &MyStruct{ID: "123", Name: "bar"})
	loggo.Success("hello world")
	loggo.Debug("hello world")
	loggo.Warning("hello world")
	loggo.Error("hello world")
}
```

Output:

```console
[2022.3.24|18:26:4|UTC-4] [loggo_test.go:11] [INFO] hello world foo &{ID:123 Name:bar}
[2022.3.24|18:26:4|UTC-4] [loggo_test.go:12] [SUCCESS] hello world
[2022.3.24|18:26:4|UTC-4] [loggo_test.go:13] [DEBUG] hello world
[2022.3.24|18:26:4|UTC-4] [loggo_test.go:14] [WARNING] hello world
[2022.3.24|18:26:4|UTC-4] [loggo_test.go:15] [ERROR] hello world
```

# Force refresh in pkg.go.dev

```console
$ git tag v0.1.0
$ git push origin v0.1.0
$ GOPROXY=proxy.golang.org go list -m github.com/wizo06/loggo@v0.1.0
$ curl proxy.golang.org/github.com/wizo06/loggo/@latest
```