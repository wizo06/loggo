package loggo

import "testing"

type Foo struct {
	Bar string
	Baz int
}

func Test(t *testing.T) {
	log := New(Config{
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
