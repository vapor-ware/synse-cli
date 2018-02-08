package client

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/urfave/cli"
)

func TestCheck1(t *testing.T) {
	var err error
	resp := http.Response{
		StatusCode: 200,
	}

	e := check(&resp, err)
	if e != nil {
		t.Errorf("expected nil but got error: %v", e)
	}
}

func TestCheck2(t *testing.T) {
	err := fmt.Errorf("test error")
	resp := http.Response{
		StatusCode: 200,
	}

	e := check(&resp, err)
	if e == nil {
		t.Error("expected error, but got nil")
	}
	_, ok := e.(cli.ExitCoder)
	if !ok {
		t.Error("expected error to fulfil cli.ExitCoder interface, but does not")
	}
}

func TestCheck3(t *testing.T) {
	var err error
	resp := http.Response{
		StatusCode: 500,
	}

	e := check(&resp, err)
	if e == nil {
		t.Error("expected error, but got nil")
	}
	_, ok := e.(cli.ExitCoder)
	if !ok {
		t.Error("expected error to fulfil cli.ExitCoder interface, but does not")
	}
}

func TestMakeURI(t *testing.T) {
	var testTable = []struct {
		in  []string
		out string
	}{
		{
			in:  []string{""},
			out: "",
		},
		{
			in:  []string{"foo"},
			out: "foo",
		},
		{
			in:  []string{"foo", "bar"},
			out: "foo/bar",
		},
		{
			in:  []string{"foo", "bar", "baz/"},
			out: "foo/bar/baz/",
		},
		{
			in:  []string{"foo/", "/bar"},
			out: "foo///bar",
		},
	}

	for _, testCase := range testTable {
		r := MakeURI(testCase.in...)
		if r != testCase.out {
			t.Errorf("MakeURI(%v) => %s, want %s", testCase.in, r, testCase.out)
		}
	}
}

func TestDoGet(t *testing.T) {
	// TODO (etd) -- need to figure out how to properly mock the HTTP call here
}

func TestDoGetUnversioned(t *testing.T) {
	// TODO (etd) -- need to figure out how to properly mock the HTTP call here
}
