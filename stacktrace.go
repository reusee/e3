package e3

import (
	"fmt"
	"runtime"
	"strings"
)

type Stacktrace struct {
	Frames []Frame
	Prev
}

type Frame struct {
	File     string
	Line     int
	Function string
}

func (s *Stacktrace) Error() string {
	var b strings.Builder
	for i, frame := range s.Frames {
		if i == 0 {
			b.WriteString("> at ")
		} else {
			b.WriteString("\n-    ")
		}
		b.WriteString(fmt.Sprintf("%s:%d %s", frame.File, frame.Line, frame.Function))
	}
	b.WriteString(s.Prev.String("\n"))
	return b.String()
}

var _ Error = new(Stacktrace)

func NewStacktrace() WrapFunc {
	return func(err error) Error {
		stacktrace := new(Stacktrace)
		stacktrace.Prev.Err = err
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
}

var WithStacktrace = NewStacktrace

var WrapStacktrace = NewStacktrace

var WithStack = NewStacktrace

var WrapStack = NewStacktrace
