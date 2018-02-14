package flags

import (
	"github.com/urfave/cli"
)

// debugFlag is the flag for setting debug mode.
var debugFlag = cli.BoolFlag{
	Name:  "debug, d",
	Usage: "enable debug mode",
}

// configFlag is the flag for displaying CLI configuration.
var configFlag = cli.BoolFlag{
	Name:  "config",
	Usage: "display the current CLI configuration",
}

// GlobalFlags is a list of flags globally available to the CLI.
var GlobalFlags = []cli.Flag{
	debugFlag,
	configFlag,
}
