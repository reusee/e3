package e3

import (
	"fmt"
	"strings"
)

type Info struct {
	Info string
	Prev
}

func (i Info) Error() string {
	var b strings.Builder
	b.WriteString(i.Info)
	if i.Prev.Err != nil {
		b.WriteString("\n")
		b.WriteString(i.Prev.Err.Error())
	}
	return b.String()
}

func NewInfo(format string, args ...any) *Info {
	return &Info{
		Info: fmt.Sprintf(format, args...),
	}
}
