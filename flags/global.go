package flags

import (
	"github.com/urfave/cli"
)

// debugFlag is the flag for setting debug mode.
var debugFlag = cli.BoolFlag{
	Name:  "debug, d",
	Usage: "enable debug mode",
}

// GlobalFlags is a list of flags globally available to the CLI.
var GlobalFlags = []cli.Flag{
	debugFlag,
}
