package e3

import (
	"io"
	"testing"
)

func TestInfo(t *testing.T) {
	info := NewInfo(io.EOF, "foo %s", "bar")
	if info.Error() != "foo bar\nEOF" {
		t.Fatalf("got %s", info.Error())
	}
	if !is(info, io.EOF) {
		t.Fatal()
	}
}
