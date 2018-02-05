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

// infoURI
const infoURI = "info"

// infoCommand
var InfoCommand = cli.Command{
	Name:     "info",
	Usage:    "info",
	Category: "Synse Server Actions",
	Action: func(c *cli.Context) error {
		return utils.CommandHandler(c, cmdInfo(c))
	},
}

// cmdInfo
func cmdInfo(c *cli.Context) error {
	rack := c.Args().Get(0)
	board := c.Args().Get(1)
	device := c.Args().Get(2)
	if rack == "" {
		return cli.NewExitError("'info' requires at least 1 argument", 1)
	}

	var resp *http.Response
	var err error

	if board == "" {
		// No board is defined, so we are querying at the rack level.
		info := &scheme.RackInfo{}
		uri := fmt.Sprintf("%s/%s", infoURI, rack)
		resp, err = client.New().Get(uri).ReceiveSuccess(info)
		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			return err
		}

		out, err := yaml.Marshal(info)
		if err != nil {
			return err
		}
		fmt.Printf("%s", out)

	} else if device == "" {
		// Board is defined, but device is not, so we are querying at the board level.
		info := &scheme.BoardInfo{}
		uri := fmt.Sprintf("%s/%s/%s", infoURI, rack, board)
		resp, err = client.New().Get(uri).ReceiveSuccess(info)
		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			return err
		}

		out, err := yaml.Marshal(info)
		if err != nil {
			return err
		}
		fmt.Printf("%s", out)

	} else {
		// Rack, Board, Device is defined, so we are querying at the device level.
		info := &scheme.DeviceInfo{}
		uri := fmt.Sprintf("%s/%s/%s/%s", infoURI, rack, board, device)
		resp, err = client.New().Get(uri).ReceiveSuccess(info)
		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			return err
		}

		out, err := yaml.Marshal(info)
		if err != nil {
			return err
		}
		fmt.Printf("%s", out)
	}

	return nil
}
