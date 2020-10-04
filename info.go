package e3

import (
	"fmt"
	"strings"
)

type Info struct {
	Info string
	Prev error
}

func (i Info) Unwrap() error {
	return i.Prev
}

func (i Info) Error() string {
	var b strings.Builder
	b.WriteString(i.Info)
	if i.Prev != nil {
		b.WriteString("\n")
		b.WriteString(i.Prev.Error())
	}
	return b.String()
}

func NewInfo(prev error, args ...any) error {
	return Info{
		Info: fmt.Sprintf(args[0].(string), args[1:]...),
		Prev: prev,
	}
}
