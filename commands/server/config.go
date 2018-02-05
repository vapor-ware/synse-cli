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

// configURI
const configURI = "config"

// configCommand
var ConfigCommand = cli.Command{
	Name:     "config",
	Usage:    "get the configuration of the active Synse Server instance",
	Category: "Synse Server Actions",
	Action: func(c *cli.Context) error {
		return utils.CommandHandler(c, cmdConfig(c))
	},
}

// cmdConfig
func cmdConfig(c *cli.Context) error {
	cfg := &scheme.Config{}
	resp, err := client.New().Get(configURI).ReceiveSuccess(cfg)
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
