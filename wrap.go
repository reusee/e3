package e3

import "reflect"

type Wrapper interface {
	error
	Wrap(error)
	Unwrap() error
}

func Wrap(
	err error,
	wrappers ...Wrapper,
) error {
	for _, wrapper := range wrappers {
		wrapper.Wrap(err)
		err = wrapper
	}
	v := reflect.ValueOf(err)
	return v.Interface().(error)
}
