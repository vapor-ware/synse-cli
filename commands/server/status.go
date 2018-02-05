package server

import (
	"fmt"
	"net/http"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/client"
	"github.com/vapor-ware/synse-cli/scheme"
	"github.com/vapor-ware/synse-cli/utils"
)

// testBase is the base URI for the "test" route.
const testBase = "test"

// StatusCommand is the CLI command for Synse Server's "test" API route.
var StatusCommand = cli.Command{
	Name:     "status",
	Usage:    "Get the status of the active host",
	Category: "Synse Server Actions",
	Action: func(c *cli.Context) error {
		return utils.CommandHandler(c, cmdStatus(c))
	},
}

// cmdStatus is the action for the StatusCommand. It makes an "status" request
// against the active Synse Server instance.
func cmdStatus(c *cli.Context) error {
	status := &scheme.TestStatus{}
	resp, err := client.NewUnversioned().Get(testBase).ReceiveSuccess(status)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return err
	}
	fmt.Printf("status:    %s\n", status.Status)
	fmt.Printf("timestamp: %s\n", status.Timestamp)
	return nil
}
