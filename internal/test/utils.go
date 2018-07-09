package test

import (
	"fmt"
	"io"
	"testing"

	"github.com/urfave/cli"
)

// ExpectNoError is a test utility to verify that the given error is nil.
func ExpectNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}
}

// ExpectExitCoderError is a test utility to verify that the given error is
// not nil and fulfils the cli.ExitCoder interface.
func ExpectExitCoderError(t *testing.T, err error) {
	if err == nil {
		t.Error("expected error, but got nil")
	}
	_, ok := err.(cli.ExitCoder)
	if !ok {
		t.Error("expected error to fulfill cli.ExitCoder interface, but does not")
	}
}

// ValidateFprint is a test utility to validate the returned error of
// fmt.Fprint.
func ValidateFprint(t *testing.T, w io.Writer, a interface{}) {
	_, err := fmt.Fprint(w, a)
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}
}
