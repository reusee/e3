package e3

import (
	"io"
	"regexp"
	"testing"
)

func TestMakeFuncs(t *testing.T) {
	fn := Make(NewInfo).WithStacktrace()
	e := fn(io.EOF, "foo %s", "bar")
	ok, err := regexp.MatchString(
		`> at .*make_test.go:[0-9]+ github.com/reusee/e3.TestMakeFuncs
-.*
-.*
foo bar
EOF`,
		e.Error(),
	)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatalf("got %s", e.Error())
	}
	if !is(e, io.EOF) {
		t.Fatal()
	}
}
