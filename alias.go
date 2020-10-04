package e3

import (
	"errors"
	"fmt"
)

type (
	any = interface{}
)

var (
	is = errors.Is
	pt = fmt.Printf
)
