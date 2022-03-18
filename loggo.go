package loggo

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

const (
	CYAN    = "\x1b[36m"
	GREEN   = "\x1b[32m"
	MAGENTA = "\x1b[35m"
	RED     = "\x1b[31m"
	RESET   = "\x1b[0m"
	YELLOW  = "\x1b[33m"
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
			CYAN,
			RESET,
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
			CYAN,
			RESET,
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
			GREEN,
			RESET,
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
			GREEN,
			RESET,
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
			MAGENTA,
			RESET,
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
			MAGENTA,
			RESET,
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
			YELLOW,
			RESET,
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
			YELLOW,
			RESET,
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
			RED,
			RESET,
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
			RED,
			RESET,
			input)
	}
}
