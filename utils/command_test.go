package utils

import (
	"flag"
	"testing"

	"fmt"

	"github.com/urfave/cli"
)

func TestRequiresArgsExactOkExact0(t *testing.T) {
	fs := &flag.FlagSet{}
	fs.Parse([]string{})
	ctx := cli.NewContext(nil, fs, nil)

	err := RequiresArgsExact(0, ctx)
	if err != nil {
		t.Error(err)
	}
}

func TestRequiresArgsExactOkExact1(t *testing.T) {
	fs := &flag.FlagSet{}
	fs.Parse([]string{"testarg"})
	ctx := cli.NewContext(nil, fs, nil)

	err := RequiresArgsExact(1, ctx)
	if err != nil {
		t.Error(err)
	}
}

func TestRequiresArgsInRangeOkMin(t *testing.T) {
	fs := &flag.FlagSet{}
	fs.Parse([]string{"testarg"})
	ctx := cli.NewContext(nil, fs, nil)

	err := RequiresArgsInRange(1, 3, ctx)
	if err != nil {
		t.Error(err)
	}
}

func TestRequiresArgsInRangeOkBetween(t *testing.T) {
	fs := &flag.FlagSet{}
	fs.Parse([]string{"testarg1", "testarg2"})
	ctx := cli.NewContext(nil, fs, nil)

	err := RequiresArgsInRange(1, 3, ctx)
	if err != nil {
		t.Error(err)
	}
}

func TestRequiresArgsInRangeOkMax(t *testing.T) {
	fs := &flag.FlagSet{}
	fs.Parse([]string{"testarg1", "testarg2", "testarg3"})
	ctx := cli.NewContext(nil, fs, nil)

	err := RequiresArgsInRange(1, 3, ctx)
	if err != nil {
		t.Error(err)
	}
}

func TestRequiresArgsInRangeErrMin(t *testing.T) {
	fs := &flag.FlagSet{}
	fs.Parse([]string{"testarg1", "testarg2", "testarg3"})
	ctx := cli.NewContext(nil, fs, nil)

	err := RequiresArgsInRange(5, 6, ctx)
	if err == nil {
		t.Error("expected error: arg count under min")
	}
}

func TestRequiresArgsInRangeErrMax(t *testing.T) {
	fs := &flag.FlagSet{}
	fs.Parse([]string{"testarg1", "testarg2", "testarg3"})
	ctx := cli.NewContext(nil, fs, nil)

	err := RequiresArgsInRange(1, 2, ctx)
	if err == nil {
		t.Error("expected error: arg count over max")
	}
}

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
