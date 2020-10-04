package e3

type thrownError struct {
	err error
	sig int64
}

func (t *thrownError) String() string { // NOCOVER
	return t.err.Error()
}

func New(
	makeErr Make,
) (
	check func(err error, args ...interface{}),
	catch func(errp *error, args ...interface{}),
) {

	check = func(err error, args ...interface{}) {
		if err != nil {
			if len(args) > 0 {
				err = makeErr(err, args...)
			}
			if _, ok := err.(Stacktrace); !ok {
				err = NewStacktrace(err)
			}
			panic(&thrownError{
				err: err,
			})
		}
	}

	catch = func(errp *error, args ...interface{}) {
		if errp == nil {
			return
		}
		if p := recover(); p != nil {
			if e, ok := p.(*thrownError); ok {
				if len(args) > 0 {
					e.err = makeErr(e.err, args...)
				}
				*errp = e.err
			} else {
				panic(p)
			}
		}
	}

	return
}
