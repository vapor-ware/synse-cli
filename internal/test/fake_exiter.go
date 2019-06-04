package test

import (
	"fmt"
	"io"
	"log"
)

// FakeExiter is an Exiter which can be used for testing.
type FakeExiter struct {
	Writer   io.Writer
	IsExited bool
	Code     int
}

func (exiter *FakeExiter) SetWriter(writer io.Writer) {
	exiter.Writer = writer
}

func (exiter *FakeExiter) Exit(code int) {
	exiter.Code = code
	exiter.IsExited = true
}

func (exiter *FakeExiter) Exitf(code int, format string, a ...interface{}) {
	_, err := fmt.Fprintf(exiter.Writer, format, a...)
	if err != nil {
		log.Fatal(err)
	}
	exiter.Exit(code)
}

func (exiter *FakeExiter) Err(err interface{}) {
	if err != nil {
		exiter.Fatal(err)
	}
}

func (exiter *FakeExiter) Fatal(msg interface{}) {
	exiter.Exitf(1, "Error: %s\n", msg)
}
