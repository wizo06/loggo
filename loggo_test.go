package loggo

import "testing"

type MyStruct struct {
	ID   string
	Name string
}

func TestAll(t *testing.T) {
	Info("hello world %s %+v", "foo", &MyStruct{ID: "123", Name: "bar"})
	Success("hello world")
	Debug("hello world")
	Warning("hello world")
	Error("hello world")
}
