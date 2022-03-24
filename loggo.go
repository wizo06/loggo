package loggo

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

const (
	cyan    = "\x1b[36m"
	green   = "\x1b[32m"
	magenta = "\x1b[35m"
	red     = "\x1b[31m"
	reset   = "\x1b[0m"
	yellow  = "\x1b[33m"
)

// Outputs to stdout with an [INFO] tag
func Info(input string, a ...any) {
	t := time.Now()
	_, offsetInSeconds := t.Zone()
	offsetInHours := offsetInSeconds / 3600
	_, file, lineNumber, ok := runtime.Caller(1)
	if ok {
		fullPath := strings.Split(file, "/")
		resultingString := fmt.Sprintf("[%d.%d.%d|%d:%d:%d|UTC%d] [%s:%d] [%sINFO%s] %s\n",
			t.Year(),
			t.Month(),
			t.Day(),
			t.Hour(),
			t.Minute(),
			t.Second(),
			offsetInHours,
			fullPath[len(fullPath)-1],
			lineNumber,
			cyan,
			reset,
			input)
		fmt.Fprintf(os.Stdout, resultingString, a...)
	} else {
		fmt.Fprintf(os.Stdout, input, a...)
	}
}

// Outputs to stdout with a [SUCCESS] tag
func Success(input string, a ...any) {
	t := time.Now()
	_, offsetInSeconds := t.Zone()
	offsetInHours := offsetInSeconds / 3600
	_, file, lineNumber, ok := runtime.Caller(1)
	if ok {
		fullPath := strings.Split(file, "/")
		resultingString := fmt.Sprintf("[%d.%d.%d|%d:%d:%d|UTC%d] [%s:%d] [%sSUCCESS%s] %s\n",
			t.Year(),
			t.Month(),
			t.Day(),
			t.Hour(),
			t.Minute(),
			t.Second(),
			offsetInHours,
			fullPath[len(fullPath)-1],
			lineNumber,
			green,
			reset,
			input)
		fmt.Fprintf(os.Stdout, resultingString, a...)
	} else {
		fmt.Fprintf(os.Stdout, input, a...)
	}
}

// Outputs to stdout with a [DEBUG] tag
func Debug(input string, a ...any) {
	t := time.Now()
	_, offsetInSeconds := t.Zone()
	offsetInHours := offsetInSeconds / 3600
	_, file, lineNumber, ok := runtime.Caller(1)
	if ok {
		fullPath := strings.Split(file, "/")
		resultingString := fmt.Sprintf("[%d.%d.%d|%d:%d:%d|UTC%d] [%s:%d] [%sDEBUG%s] %s\n",
			t.Year(),
			t.Month(),
			t.Day(),
			t.Hour(),
			t.Minute(),
			t.Second(),
			offsetInHours,
			fullPath[len(fullPath)-1],
			lineNumber,
			magenta,
			reset,
			input)
		fmt.Fprintf(os.Stdout, resultingString, a...)
	} else {
		fmt.Fprintf(os.Stdout, input, a...)
	}
}

// Outputs to stderr with a [WARNING] tag
func Warning(input string, a ...any) {
	t := time.Now()
	_, offsetInSeconds := t.Zone()
	offsetInHours := offsetInSeconds / 3600
	_, file, lineNumber, ok := runtime.Caller(1)
	if ok {
		fullPath := strings.Split(file, "/")
		resultingString := fmt.Sprintf("[%d.%d.%d|%d:%d:%d|UTC%d] [%s:%d] [%sWARNING%s] %s\n",
			t.Year(),
			t.Month(),
			t.Day(),
			t.Hour(),
			t.Minute(),
			t.Second(),
			offsetInHours,
			fullPath[len(fullPath)-1],
			lineNumber,
			yellow,
			reset,
			input)
		fmt.Fprintf(os.Stderr, resultingString, a...)
	} else {
		fmt.Fprintf(os.Stderr, input, a...)
	}
}

// Outputs to stderr with an [ERROR] tag
func Error(input string, a ...any) {
	t := time.Now()
	_, offsetInSeconds := t.Zone()
	offsetInHours := offsetInSeconds / 3600
	_, file, lineNumber, ok := runtime.Caller(1)
	if ok {
		fullPath := strings.Split(file, "/")
		resultingString := fmt.Sprintf("[%d.%d.%d|%d:%d:%d|UTC%d] [%s:%d] [%sERROR%s] %s\n",
			t.Year(),
			t.Month(),
			t.Day(),
			t.Hour(),
			t.Minute(),
			t.Second(),
			offsetInHours,
			fullPath[len(fullPath)-1],
			lineNumber,
			red,
			reset,
			input)
		fmt.Fprintf(os.Stderr, resultingString, a...)
	} else {
		fmt.Fprintf(os.Stderr, input, a...)
	}
}
