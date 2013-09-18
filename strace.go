package sys

import (
	"fmt"
	"github.com/mgutz/ansi"
	"io"
	"os"
	"runtime"
	"strings"
	"syscall"
	"time"
)

var DEBUG, TRACE string
var straceOut io.Writer
var debug, trace bool
var depth int

func init() {
	depth, DEBUG, TRACE = 0, "DEBUG", "TRACE"

	if straceOut == nil {
		straceOut = os.Stderr
	}

	if val, found := syscall.Getenv("DEBUG"); found && val != "" {
		debug = true
	} else {
		debug = false
	}

	if val, found := syscall.Getenv("TRACE"); found && val != "" {
		trace = true
	} else {
		trace = false
	}
}

func Interactive() {
	reset := ansi.ColorCode("reset")
	DEBUG = ansi.Color("DEBUG", "red") + reset
	TRACE = ansi.Color("TRACE", "cyan") + reset
}

func Output(writer io.Writer) {
	straceOut = writer
}

func DebugNil() {
	if debug {
		pc, _, _, _ := runtime.Caller(1)
		funcName := runtime.FuncForPC(pc).Name()
		fmt.Fprintf(straceOut, "%v: %v(nil) -> nil\n", DEBUG, funcName)
	}
}

func Debug(input, output string) error {
	if debug {
		pc, _, _, _ := runtime.Caller(1)
		funcName := runtime.FuncForPC(pc).Name()
		fmt.Fprintf(straceOut, "%v: %v(%v) -> (%v)\n",
			DEBUG, funcName, input, output)
	}
	return nil
}

func Trace() error {
	if trace {
		pc, _, _, _ := runtime.Caller(1)
		funcName := runtime.FuncForPC(pc).Name()
		fmt.Fprintf(straceOut, "%v:%s->%v()@%v\n",
			TRACE, strings.Repeat(" ", depth),
			funcName, time.Now())
		depth += 1
	}
	return nil
}

func UnTrace(_ error) {
	if trace {
		pc, _, _, _ := runtime.Caller(1)
		funcName := runtime.FuncForPC(pc).Name()
		depth -= 1
		fmt.Fprintf(straceOut, "%v:%s<-%v()@%v\n",
			TRACE, strings.Repeat(" ", depth),
			funcName, time.Now())
	}
	return
}
