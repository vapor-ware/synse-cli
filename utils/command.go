package utils

import (
	"fmt"

	"github.com/urfave/cli"
)

// NoArgs is used to specify that a command does not use any arguments.
const NoArgs = " "

// CmdHandler is used to wrap a command action. If an error is returned from that
// action that already fulfils the ExitCoder interface, it is returned, otherwise
// the error is made into a new ExitError and returned. This way, the CLI will exit
// on command error.
func CmdHandler(err error) error {
	if err != nil {
		if exitErr, ok := err.(cli.ExitCoder); ok {
			return exitErr
		}
		return cli.NewExitError(err.Error(), 1)
	}
	return nil
}

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
