package client

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
)

func TestCheck1(t *testing.T) {
	var err error
	resp := http.Response{
		StatusCode: 200,
	}

	errScheme := new(scheme.Error)
	e := check(&resp, err, errScheme)
	test.ExpectNoError(t, e)
}

func TestCheck2(t *testing.T) {
	err := fmt.Errorf("test error")
	resp := http.Response{
		StatusCode: 200,
	}

	errScheme := new(scheme.Error)
	e := check(&resp, err, errScheme)
	test.ExpectExitCoderError(t, e)
}

func TestCheck3(t *testing.T) {
	var err error
	resp := http.Response{
		StatusCode: 500,
	}

	errScheme := new(scheme.Error)
	e := check(&resp, err, errScheme)
	test.ExpectExitCoderError(t, e)
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

	cli.OsExiter = func(int) {}

	for _, testCase := range testTable {
		r := makeURI(testCase.in...)
		if r != testCase.out {
			t.Errorf("makeURI(%v) => %s, want %s", testCase.in, r, testCase.out)
		}
	}
}

func TestGetVersioned(t *testing.T) {
	test.Setup()
	cli.OsExiter = func(int) {}

	type testData struct {
		Status string
	}

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc("/synse/2.0/foobar", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		test.ValidateFprint(t, w, `{"status": "ok"}`)
	})

	test.AddServerHost(server)

	out := &testData{}
	err := getVersioned("foobar", out)

	test.ExpectNoError(t, err)

	if out.Status != "ok" {
		t.Errorf("expected status 'ok', but was %s", out.Status)
	}
}

func TestGetUnversioned(t *testing.T) {
	test.Setup()
	cli.OsExiter = func(int) {}

	type testData struct {
		Status string
	}

	mux, server := test.UnversionedServer()
	defer server.Close()
	mux.HandleFunc("/synse/foobar", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		test.ValidateFprint(t, w, `{"status": "ok"}`)
	})

	test.AddServerHost(server)

	out := &testData{}
	err := getUnversioned("foobar", out)

	test.ExpectNoError(t, err)

	if out.Status != "ok" {
		t.Errorf("expected status 'ok', but was %s", out.Status)
	}
}

func TestPostVersioned(t *testing.T) {
	test.Setup()
	cli.OsExiter = func(int) {}

	type testData struct {
		Status string
	}

	mux, server := test.Server()
	defer server.Close()
	mux.HandleFunc("/synse/2.0/foobar", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected request to be POST, but was %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		test.ValidateFprint(t, w, `{"status": "ok"}`)
	})

	test.AddServerHost(server)

	out := &testData{}
	err := postVersioned("foobar", testData{}, out)

	test.ExpectNoError(t, err)

	if out.Status != "ok" {
		t.Errorf("expected status 'ok', but was %s", out.Status)
	}
}
