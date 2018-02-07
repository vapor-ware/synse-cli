package utils

import (
	"flag"
	"testing"

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
