package server

import "github.com/urfave/cli"

// NewHostsCommand
func NewServerCommands() []cli.Command {
	return []cli.Command{
		StatusCommand,
	}
}
