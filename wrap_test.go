package e3

import (
	"io"
	"testing"
)

func TestWrap(t *testing.T) {
	err := Wrap(io.EOF, &Info{Info: "foo"})
	if !is(err, io.EOF) {
		t.Fatal()
	}
	if info := (Info{}); !as(err, &info) {
		t.Fatal()
	} else if info.Info != "foo" {
		t.Fatal()
	}
}

func TestWrapReturn(t *testing.T) {
	err := func() (err error) {
		defer WrapReturn(&err, &Info{Info: "foo"})
		return io.EOF
	}()
	if info := (Info{}); !as(err, &info) {
		t.Fatal()
	} else if info.Info != "foo" {
		t.Fatal()
	}

	err = func() (err error) {
		defer WrapReturn(nil)
		return io.EOF
	}()
	if err != io.EOF {
		t.Fatal()
	}

	err = func() (err error) {
		defer WrapReturn(&err, &Info{Info: "foo"})
		return nil
	}()
	if err != nil {
		t.Fatal()
	}
}
