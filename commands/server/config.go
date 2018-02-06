package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/scheme"
	"github.com/vapor-ware/synse-cli/utils"
)

// configBase is the base URI for the "config" route.
const configBase = "config"

// ConfigCommand is the CLI command for Synse Server's "config" API route.
var ConfigCommand = cli.Command{
	Name:     "config",
	Usage:    "Get the configuration for the active host",
	Category: "Synse Server Actions",
	Action: func(c *cli.Context) error {
		return utils.CmdHandler(c, cmdConfig(c))
	},
}

// cmdConfig is the action for the ConfigCommand. It makes an "config" request
// against the active Synse Server instance.
func cmdConfig(c *cli.Context) error {
	cfg := &scheme.Config{}
	err := utils.DoGet(configBase, cfg)
	if err != nil {
		return err
	}

	return utils.AsYAML(cfg)
}
