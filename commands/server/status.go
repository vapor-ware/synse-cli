package server

import (
	"fmt"
	"net/http"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/client"
	"github.com/vapor-ware/synse-cli/scheme"
	"github.com/vapor-ware/synse-cli/utils"
)

// testURI
const testURI = "test"

// statusCommand
var StatusCommand = cli.Command{
	Name:  "status",
	Usage: "get the status of the active Synse Server instance",
	Action: func(c *cli.Context) error {
		return utils.CommandHandler(c, cmdStatus(c))
	},
}

// cmdStatus
func cmdStatus(c *cli.Context) error {
	status := &scheme.TestStatus{}
	resp, err := client.NewUnversioned().Get(testURI).ReceiveSuccess(status)
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
