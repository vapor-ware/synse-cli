package flags

import (
	"github.com/urfave/cli"
)

// OutputFlag is the flag for setting the output format of a command.
var OutputFlag = cli.StringFlag{
	Name:  "output, o",
	Value: "yaml",
	Usage: "set the output format of the command",
}

// FilterFlag is the flag for setting a filter on a command's output.
var FilterFlag = cli.StringFlag{
	Name:  "filter, f",
	Usage: "set a filter for the output results",
}

// HostFlag is the flag for setting the hostname for a plugin.
var HostFlag = cli.StringFlag{
	Name:  "tcp, t",
	Usage: "set the hostname/ip[:port] for a plugin",
}

// UnixFlag is the flag for setting the unix socket path for a plugin.
var UnixFlag = cli.StringFlag{
	Name:  "unix, u",
	Usage: "set the unix socket path for a plugin",
}
