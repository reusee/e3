package e3

import "reflect"

type Wrapper interface {
	error
	Wrap(error)
	Unwrap() error
}

func Wrap(
	err error,
	targets ...Wrapper,
) error {
	for _, target := range targets {
		target.Wrap(err)
		err = target
	}
	v := reflect.ValueOf(err)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Interface().(error)
}

func WrapReturn(p *error, args ...Wrapper) {
	if p == nil {
		return
	}
	err := *p
	if err == nil {
		return
	}
	*p = Wrap(err, args...)
}
