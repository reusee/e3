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
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Interface().(error)
}
