package e3

type WrapFunc func(err error) Error

func Wrap(err error, fns ...WrapFunc) error {
	for _, fn := range fns {
		err = fn(err)
	}
	return err
}
