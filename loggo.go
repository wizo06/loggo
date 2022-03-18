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

func Info(input string) {
	t := time.Now()
	_, offsetInSeconds := t.Zone()
	offsetInHours := offsetInSeconds / 3600
	_, file, lineNumber, ok := runtime.Caller(1)
	if ok {
		fullPath := strings.Split(file, "/")
		fmt.Fprintf(os.Stdout, "[%d.%d.%d|%d:%d:%d|UTC%d] [%s:%d] [%sINFO%s] %s\n",
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
	} else {
		fmt.Fprintf(os.Stdout, "[%d.%d.%d|%d:%d:%d|UTC%d] [%sINFO%s] %s\n",
			t.Year(),
			t.Month(),
			t.Day(),
			t.Hour(),
			t.Minute(),
			t.Second(),
			offsetInHours,
			cyan,
			reset,
			input)
	}
}

func Success(input string) {
	t := time.Now()
	_, offsetInSeconds := t.Zone()
	offsetInHours := offsetInSeconds / 3600
	_, file, lineNumber, ok := runtime.Caller(1)
	if ok {
		fullPath := strings.Split(file, "/")
		fmt.Fprintf(os.Stdout, "[%d.%d.%d|%d:%d:%d|UTC%d] [%s:%d] [%sSUCCESS%s] %s\n",
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
	} else {
		fmt.Fprintf(os.Stdout, "[%d.%d.%d|%d:%d:%d|UTC%d] [%sSUCCESS%s] %s\n",
			t.Year(),
			t.Month(),
			t.Day(),
			t.Hour(),
			t.Minute(),
			t.Second(),
			offsetInHours,
			green,
			reset,
			input)
	}
}

func Debug(input string) {
	t := time.Now()
	_, offsetInSeconds := t.Zone()
	offsetInHours := offsetInSeconds / 3600
	_, file, lineNumber, ok := runtime.Caller(1)
	if ok {
		fullPath := strings.Split(file, "/")
		fmt.Fprintf(os.Stdout, "[%d.%d.%d|%d:%d:%d|UTC%d] [%s:%d] [%sDEBUG%s] %s\n",
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
	} else {
		fmt.Fprintf(os.Stdout, "[%d.%d.%d|%d:%d:%d|UTC%d] [%sDEBUG%s] %s\n",
			t.Year(),
			t.Month(),
			t.Day(),
			t.Hour(),
			t.Minute(),
			t.Second(),
			offsetInHours,
			magenta,
			reset,
			input)
	}
}

func Warning(input string) {
	t := time.Now()
	_, offsetInSeconds := t.Zone()
	offsetInHours := offsetInSeconds / 3600
	_, file, lineNumber, ok := runtime.Caller(1)
	if ok {
		fullPath := strings.Split(file, "/")
		fmt.Fprintf(os.Stderr, "[%d.%d.%d|%d:%d:%d|UTC%d] [%s:%d] [%sWARNING%s] %s\n",
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
	} else {
		fmt.Fprintf(os.Stderr, "[%d.%d.%d|%d:%d:%d|UTC%d] [%sWARNING%s] %s\n",
			t.Year(),
			t.Month(),
			t.Day(),
			t.Hour(),
			t.Minute(),
			t.Second(),
			offsetInHours,
			yellow,
			reset,
			input)
	}
}

func Error(input string) {
	t := time.Now()
	_, offsetInSeconds := t.Zone()
	offsetInHours := offsetInSeconds / 3600
	_, file, lineNumber, ok := runtime.Caller(1)
	if ok {
		fullPath := strings.Split(file, "/")
		fmt.Fprintf(os.Stderr, "[%d.%d.%d|%d:%d:%d|UTC%d] [%s:%d] [%sERROR%s] %s\n",
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
	} else {
		fmt.Fprintf(os.Stderr, "[%d.%d.%d|%d:%d:%d|UTC%d] [%sERROR%s] %s\n",
			t.Year(),
			t.Month(),
			t.Day(),
			t.Hour(),
			t.Minute(),
			t.Second(),
			offsetInHours,
			red,
			reset,
			input)
	}
}
