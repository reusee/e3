package e3

type Prev struct {
	error
}

func (p Prev) Unwrap() error {
	return p.error
}

func (p *Prev) Wrap(err error) {
	p.error = err
}
