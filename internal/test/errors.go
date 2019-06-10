package test

import "github.com/pkg/errors"

var (
	// ErrFakeClient is an error that is returned when a
	// fake test client is configured to return an error.
	ErrFakeClient = errors.New("fake client err")
)
