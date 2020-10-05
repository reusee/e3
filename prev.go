package e3

type Prev struct {
	Err error
}

func (p Prev) Unwrap() error {
	return p.Err
}

func (p *Prev) Wrap(err error) {
	p.Err = err
}
