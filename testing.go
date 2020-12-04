package e3

import "testing"

func TestingFatalf(t *testing.T, format string, args ...any) WrapFunc {
	return func(err error) Error {
		t.Fatal(err)
		return NewInfo(format, args...)(err)
	}
}

func TestingFatal(t *testing.T) WrapFunc {
	return TestingFatalf(t, "testing Fatal")
}
