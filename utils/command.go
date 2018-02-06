package utils

import (
	"github.com/urfave/cli"
	"fmt"
)

func RequiresArgsInRange(min, max int, c *cli.Context) error {
	if c.NArg() < min || c.NArg() > max {
		return cli.NewExitError(
			fmt.Sprintf("command '%v' requires between %d and %d arguments, %d given", c.Command.Name, min, max, c.NArg()),
			1,
		)
	}
	return nil
}


func RequiresArgsExact(count int, c *cli.Context) error {
	if c.NArg() != count {
		return cli.NewExitError(
			fmt.Sprintf("command '%v' requires exactly %d arguments, %d given", c.Command.Name, count, c.NArg()),
			1,
		)
	}
	return nil
}