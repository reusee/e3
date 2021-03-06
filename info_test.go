package e3

import (
	"io"
	"testing"
)

func TestInfo(t *testing.T) {
	info := NewInfo("foo %s", "bar")(io.EOF)
	if info.Error() != "foo bar\nEOF" {
		t.Fatalf("got %s", info.Error())
	}
	if !is(info, io.EOF) {
		t.Fatal()
	}
}
