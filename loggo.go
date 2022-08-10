package loggo

import (
	"fmt"
	"os"
)

type Log struct {
	PrintHostname               bool
	PrintUNIXTimestamp          bool
	PrintHumanReadableTimestamp bool
	StackDepth                  int
	PrintFileName               bool
	PrintFunctionName           bool
	PrintLineNumber             bool
	PrintLogLevel               bool
}

type Config struct {
	PrintHostname               bool
	PrintUNIXTimestamp          bool
	PrintHumanReadableTimestamp bool
	StackDepth                  int
	PrintFileName               bool
	PrintFunctionName           bool
	PrintLineNumber             bool
	PrintLogLevel               bool
}

func New(c Config) *Log {
	return &Log{
		PrintHostname:               c.PrintHostname,
		PrintUNIXTimestamp:          c.PrintUNIXTimestamp,
		PrintHumanReadableTimestamp: c.PrintHumanReadableTimestamp,
		StackDepth:                  c.StackDepth,
		PrintFileName:               c.PrintFileName,
		PrintFunctionName:           c.PrintFunctionName,
		PrintLineNumber:             c.PrintLineNumber,
		PrintLogLevel:               c.PrintLogLevel,
	}
}

func (log *Log) Info(a ...any) {
	s := createPrintString(
		log.PrintHostname,
		log.PrintUNIXTimestamp,
		log.PrintHumanReadableTimestamp,
		log.StackDepth,
		log.PrintFileName,
		log.PrintFunctionName,
		log.PrintLineNumber,
		log.PrintLogLevel,
		loglevel_info,
	)
	fmt.Fprintln(os.Stdout, s, fmt.Sprintf("%+v", a))
}

func (log *Log) Success(a ...any) {
	s := createPrintString(
		log.PrintHostname,
		log.PrintUNIXTimestamp,
		log.PrintHumanReadableTimestamp,
		log.StackDepth,
		log.PrintFileName,
		log.PrintFunctionName,
		log.PrintLineNumber,
		log.PrintLogLevel,
		loglevel_success,
	)
	fmt.Fprintln(os.Stdout, s, fmt.Sprintf("%+v", a))
}

func (log *Log) Debug(a ...any) {
	s := createPrintString(
		log.PrintHostname,
		log.PrintUNIXTimestamp,
		log.PrintHumanReadableTimestamp,
		log.StackDepth,
		log.PrintFileName,
		log.PrintFunctionName,
		log.PrintLineNumber,
		log.PrintLogLevel,
		loglevel_debug,
	)
	fmt.Fprintln(os.Stdout, s, fmt.Sprintf("%+v", a))
}

func (log *Log) Warn(a ...any) {
	s := createPrintString(
		log.PrintHostname,
		log.PrintUNIXTimestamp,
		log.PrintHumanReadableTimestamp,
		log.StackDepth,
		log.PrintFileName,
		log.PrintFunctionName,
		log.PrintLineNumber,
		log.PrintLogLevel,
		loglevel_warn,
	)
	fmt.Fprintln(os.Stderr, s, fmt.Sprintf("%+v", a))
}

func (log *Log) Error(a ...any) {
	s := createPrintString(
		log.PrintHostname,
		log.PrintUNIXTimestamp,
		log.PrintHumanReadableTimestamp,
		log.StackDepth,
		log.PrintFileName,
		log.PrintFunctionName,
		log.PrintLineNumber,
		log.PrintLogLevel,
		loglevel_error,
	)
	fmt.Fprintln(os.Stderr, s, fmt.Sprintf("%+v", a))
}
