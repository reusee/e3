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
