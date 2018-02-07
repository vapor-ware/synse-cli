package utils

import (
	"fmt"

	"github.com/urfave/cli"
)

// NoArgs is used to specify that a command does not use any arguments.
const NoArgs = " "

// RequiresArgsInRange checks the number of arguments passed to the command against the
// specified expected minimum and maximum number of supported arguments. If the number
// of arguments is out of the specified range, an error is returned.
func RequiresArgsInRange(min, max int, c *cli.Context) error {
	if c.NArg() < min || c.NArg() > max {
		return cli.NewExitError(
			fmt.Sprintf("command '%v' requires between %d and %d arguments, %d given", c.Command.Name, min, max, c.NArg()),
			1,
		)
	}
	return nil
}

// RequiresArgsExact checks the number of arguments passed to the command against the
// specified exact count of expected supported arguments. If the number of arguments
// do not match, an error is returned.
func RequiresArgsExact(count int, c *cli.Context) error {
	if c.NArg() != count {
		return cli.NewExitError(
			fmt.Sprintf("command '%v' requires exactly %d arguments, %d given", c.Command.Name, count, c.NArg()),
			1,
		)
	}
	return nil
}
