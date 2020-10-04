package e3

import (
	"io"
	"regexp"
	"testing"
)

func TestCheck(t *testing.T) {
	ce, he := New(Default)

	err := func() (err error) {
		defer he(&err)
		ce(io.EOF)
		return
	}()
	if !is(err, io.EOF) {
		t.Fatal()
	}

	err = func() (err error) {
		defer he(&err)
		ce(io.EOF, "foo %s", "bar")
		return
	}()
	if !is(err, io.EOF) {
		t.Fatal()
	}
	ok, e := regexp.MatchString(
		"foo bar\nEOF",
		err.Error(),
	)
	if e != nil {
		t.Fatal(e)
	}
	if !ok {
		t.Fatalf("got %s", err.Error())
	}

	err = func() (err error) {
		defer he(&err, "foo %s", "bar")
		ce(io.EOF)
		return
	}()
	if !is(err, io.EOF) {
		t.Fatal()
	}
	ok, e = regexp.MatchString(
		"foo bar\n> at .*check_test.go:[0-9]+.*\n-.*\n-.*\n-.*\nEOF",
		err.Error(),
	)
	if e != nil {
		t.Fatal(e)
	}
	if !ok {
		t.Fatalf("got %s", err.Error())
	}

	func() {
		defer func() {
			p := recover()
			te, ok := p.(*thrownError)
			if !ok {
				t.Fatal()
			}
			if !is(te.err, io.EOF) {
				t.Fatal()
			}
		}()
		func() (err error) {
			defer he(nil)
			ce(io.EOF)
			return
		}()
	}()

	func() {
		defer func() {
			p := recover()
			if p != 42 {
				t.Fatal()
			}
		}()
		func() (err error) {
			defer he(&err)
			panic(42)
		}()
	}()

}
