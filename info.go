package e3

import (
	"fmt"
	"strings"
)

type Info struct {
	Prev
	Info string
}

func (i *Info) Error() string {
	var b strings.Builder
	b.WriteString(i.Info)
	b.WriteString(i.Prev.String("\n"))
	return b.String()
}

var _ Error = new(Info)

func NewInfo(format string, args ...any) WrapFunc {
	return func(err error) Error {
		info := &Info{
			Info: fmt.Sprintf(format, args...),
		}
		info.Prev.Err = err
		return info
	}
}

var WrapInfo = NewInfo

var WithInfo = NewInfo
