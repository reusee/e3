package e3

import "testing"

func TestingFatalf(t *testing.T, format string, args ...any) WrapFunc {
	t.Helper()
	return func(err error) Error {
		t.Helper()
		t.Fatal(err)
		return NewInfo(format, args...)(err)
	}
}

func TestingFatal(t *testing.T) WrapFunc {
	t.Helper()
	return TestingFatalf(t, "testing Fatal")
}
