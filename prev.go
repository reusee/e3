package e3

type Prev struct {
	Err error
}

func (p Prev) Unwrap() error {
	return p.Err
}

func (p Prev) String(prefix string) string {
	if p.Err == nil {
		return ""
	}
	return prefix + p.Err.Error()
}
