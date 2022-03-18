package loggo

import "testing"

func TestAll(t *testing.T) {
	Info("hello world")
	Success("hello world")
	Debug("hello world")
	Warning("hello world")
	Error("hello world")
}
