package e3

import (
	"io"
	"testing"
)

func TestInfo(t *testing.T) {
	info := Wrap(io.EOF, NewInfo("foo %s", "bar"))
	if info.Error() != "foo bar\nEOF" {
		t.Fatalf("got %s", info.Error())
	}
	if !is(info, io.EOF) {
		t.Fatal()
	}
}
