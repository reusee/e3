package e3

type Error interface {
	error
	Unwrap() error
}
