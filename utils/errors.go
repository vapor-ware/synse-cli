package utils

import (
	"github.com/urfave/cli"
)

// CmdHandler is used to wrap a command action. If an error is returned from that
// action that already fulfils the ExitCoder interface, it is returned, otherwise
// the error is made into a new ExitError and returned. This way, the CLI will exit
// on command error.
func CmdHandler(c *cli.Context, err error) error {
	if err != nil {
		if exitErr, ok := err.(cli.ExitCoder); ok {
			return exitErr
		}
		return cli.NewExitError(err.Error(), 1)
	}
	return nil
}
