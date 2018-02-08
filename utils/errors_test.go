package utils

import (
	"fmt"
	"testing"

	"github.com/urfave/cli"
)

func TestCmdHandlerCliErr(t *testing.T) {
	err := cli.NewExitError("test cli error", 1)
	e := CmdHandler(err)
	if e == nil {
		t.Error("expected error but got nil")
	}
	_, ok := e.(cli.ExitCoder)
	if !ok {
		t.Error("expected returned error to be of type cli.ExitCoder but was not")
	}
}

func TestCmdHandlerRegErr(t *testing.T) {
	err := fmt.Errorf("test regular error")
	e := CmdHandler(err)
	if e == nil {
		t.Error("expected error but got nil")
	}
	_, ok := e.(cli.ExitCoder)
	if !ok {
		t.Error("expected returned error to be of type cli.ExitCoder but was not")
	}
}

func TestCmdHandlerNoErr(t *testing.T) {
	var err error
	e := CmdHandler(err)
	if e != nil {
		t.Errorf("expected nil, but got %v", e)
	}
}
