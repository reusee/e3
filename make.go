package e3

type Make func(
	prev error,
	args ...any,
) error

var Default = Make(NewInfo)
