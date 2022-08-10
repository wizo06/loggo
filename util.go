package loggo

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func createPrintString(host, unix, human bool, stackDepth int, file, function, line, level bool, logLevel LogLevel) string {
	msg := []string{}

	if host {
		name, _ := os.Hostname()
		msg = append(msg, fmt.Sprintf("[%s]", name))
	}

	if unix || human {
		msg = append(msg, fmt.Sprintf("[%s]", getTimestamp(unix, human)))
	}

	if file {
		msg = append(msg, fmt.Sprintf("[%s]", getCaller(stackDepth, line, function)))
	}

	if level {
		msg = append(msg, fmt.Sprintf("[%s]", formatLogLeveL(logLevel)))
	}

	return strings.Join(msg, " ")
}

func getTimestamp(unix, human bool) string {
	now := time.Now()
	_, offset := now.Zone()
	humanReadableTimestamp := fmt.Sprintf(
		"%d.%d.%d|%d:%d:%d|UTC%d",
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second(),
		offset/3600,
	)

	if unix && human {
		return fmt.Sprintf("%d|%s", now.Unix(), humanReadableTimestamp)
	}

	if unix {
		return strconv.Itoa(int(now.Unix()))
	}

	return humanReadableTimestamp
}

func getCaller(stackDepth int, line, function bool) string {
	pc, file, lineNumber, _ := runtime.Caller(stackDepth)
	fullFuncName := runtime.FuncForPC(pc).Name()
	baseName := path.Base(fullFuncName)
	funcName := strings.Join(strings.Split(baseName, ".")[1:], ".")

	if line && function {
		// file.go:functionName:123
		return fmt.Sprintf("%s:%s:%d", path.Base(file), funcName, lineNumber)
	}

	if line && !function {
		// file.go:123
		return fmt.Sprintf("%s:%d", path.Base(file), lineNumber)
	}

	if !line && function {
		// file.go:functionName
		return fmt.Sprintf("%s:%s", path.Base(file), funcName)
	}

	// file.go
	return path.Base(file)
}

type LogLevel string

const (
	loglevel_info    LogLevel = "info"
	loglevel_success LogLevel = "success"
	loglevel_debug   LogLevel = "debug"
	loglevel_warn    LogLevel = "warn"
	loglevel_error   LogLevel = "error"
)

func formatLogLeveL(logLevel LogLevel) string {
	if logLevel == loglevel_info {
		return fmt.Sprintf("%s%s%s", CYAN, logLevel, RESET)
	}

	if logLevel == loglevel_success {
		return fmt.Sprintf("%s%s%s", GREEN, logLevel, RESET)
	}

	if logLevel == loglevel_debug {
		return fmt.Sprintf("%s%s%s", MAGENTA, logLevel, RESET)
	}

	if logLevel == loglevel_warn {
		return fmt.Sprintf("%s%s%s", YELLOW, logLevel, RESET)
	}

	if logLevel == loglevel_error {
		return fmt.Sprintf("%s%s%s", RED, logLevel, RESET)
	}

	return fmt.Sprintf("%s%s%s", CYAN, loglevel_info, RESET)
}
