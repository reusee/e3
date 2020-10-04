package e3

import (
	"fmt"
	"runtime"
	"strings"
)

type Stacktrace struct {
	Frames []Frame
	Prev   error
}

type Frame struct {
	File     string
	Line     int
	Function string
}

func (s Stacktrace) Unwrap() error {
	return s.Prev
}

func (s Stacktrace) Error() string {
	var b strings.Builder
	for i, frame := range s.Frames {
		if i == 0 {
			b.WriteString("> at ")
		} else {
			b.WriteString("\n-    ")
		}
		b.WriteString(fmt.Sprintf("%s:%d %s", frame.File, frame.Line, frame.Function))
	}
	if s.Prev != nil {
		b.WriteString("\n")
		b.WriteString(s.Prev.Error())
	}
	return b.String()
}

func NewStacktrace(prev error) error {
	stacktrace := Stacktrace{
		Prev: prev,
	}
	numPCs := 32
	for {
		pcs := make([]uintptr, numPCs)
		n := runtime.Callers(1, pcs)
		if n == len(pcs) { // NOCOVER
			numPCs *= 2
			continue
		}
		pcs = pcs[:n]
		frames := runtime.CallersFrames(pcs)
		for {
			frame, more := frames.Next()
			if strings.HasPrefix(frame.Function, "github.com/reusee/e3.") &&
				!strings.HasPrefix(frame.Function, "github.com/reusee/e3.Test") {
				// internal funcs
				continue
			}
			stacktrace.Frames = append(stacktrace.Frames, Frame{
				File:     frame.File,
				Line:     frame.Line,
				Function: frame.Function,
			})
			if !more {
				break
			}
		}
		break
	}
	return stacktrace
}

func (n Make) WithStacktrace() Make {
	return func(prev error, args ...any) error {
		return NewStacktrace(n(prev, args...))
	}
}
