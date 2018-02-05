package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/utils"
	"github.com/vapor-ware/synse-cli/scheme"
	"github.com/vapor-ware/synse-cli/client"
	"net/http"
	"fmt"
	"gopkg.in/yaml.v2"
)

// configURI
const configURI = "config"

// configCommand
var ConfigCommand = cli.Command{
	Name:    "config",
	Usage:   "get the configuration of the active Synse Server instance",
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
