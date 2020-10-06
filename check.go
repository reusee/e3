package e3

type thrownError struct {
	err error
	sig int64
}

func (t *thrownError) String() string { // NOCOVER
	return t.err.Error()
}

func Check(err error, wrappers ...Wrapper) {
	if err != nil {
		if len(wrappers) > 0 {
			err = Wrap(err, wrappers...)
		}
		if _, ok := err.(*Stacktrace); !ok {
			err = Wrap(err, NewStacktrace())
		}
		panic(&thrownError{
			err: err,
		})
	}
}

func Catch(errp *error, wrappers ...Wrapper) {
	if errp == nil {
		return
	}
	if p := recover(); p != nil {
		if e, ok := p.(*thrownError); ok {
			if len(wrappers) > 0 {
				e.err = Wrap(e.err, wrappers...)
			}
			*errp = e.err
		} else {
			panic(p)
		}
	} else {
		err := *errp
		if err == nil {
			return
		}
		*errp = Wrap(err, wrappers...)
	}
}
