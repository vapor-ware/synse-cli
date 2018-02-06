package server

import (
	"fmt"
	"net/http"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/client"
	"github.com/vapor-ware/synse-cli/scheme"
	"github.com/vapor-ware/synse-cli/utils"
	"gopkg.in/yaml.v2"
)

// configBase is the base URI for the "config" route.
const configBase = "config"

// ConfigCommand is the CLI command for Synse Server's "config" API route.
var ConfigCommand = cli.Command{
	Name:     "config",
	Usage:    "Get the configuration for the active host",
	Category: "Synse Server Actions",
	Action: func(c *cli.Context) error {
		return utils.CommandHandler(c, cmdConfig(c))
	},
}

// cmdConfig is the action for the ConfigCommand. It makes an "config" request
// against the active Synse Server instance.
func cmdConfig(c *cli.Context) error {
	cfg := &scheme.Config{}
	resp, err := client.New().Get(configBase).ReceiveSuccess(cfg)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return err
	}

	out, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	fmt.Printf("%s", out)
	return nil
}
