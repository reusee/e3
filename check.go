package e3

type thrownError struct {
	err error
	sig int64
}

func (t *thrownError) String() string { // NOCOVER
	return t.err.Error()
}

func Check(err error, fns ...WrapFunc) {
	if err == nil {
		return
	}
	for _, fn := range fns {
		err = fn(err)
	}
	if _, ok := err.(*Stacktrace); !ok {
		err = NewStacktrace()(err)
	}
	panic(&thrownError{
		err: err,
	})
}

func Catch(errp *error, fns ...WrapFunc) {
	if errp == nil {
		return
	}
	var err error
	if p := recover(); p != nil {
		if e, ok := p.(*thrownError); ok {
			err = e.err
		} else {
			panic(p)
		}
	} else {
		if *errp != nil {
			err = *errp
		}
	}
	if err == nil {
		return
	}
	for _, fn := range fns {
		err = fn(err)
	}
	*errp = err
}

var C = Check

var T = Catch
