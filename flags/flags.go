package flags

import (
	"github.com/urfave/cli"
)

var debugFlag = cli.BoolFlag{
	Name:  "debug, d",
	Usage: "enable debug mode",
}

var Flags = []cli.Flag{
	debugFlag,
}
