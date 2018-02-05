package server

import (
	"fmt"
	"net/http"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/client"
	"github.com/vapor-ware/synse-cli/scheme"
	"github.com/vapor-ware/synse-cli/utils"
)

// versionURI
const versionURI = "version"

// versionCommand
var VersionCommand = cli.Command{
	Name:  "version",
	Usage: "get the version of the active Synse Server instance",
	Action: func(c *cli.Context) error {
		return utils.CommandHandler(c, cmdVersion(c))
	},
}

// cmdVersion
func cmdVersion(c *cli.Context) error {
	version := &scheme.Version{}
	resp, err := client.NewUnversioned().Get(versionURI).ReceiveSuccess(version)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return err
	}
	fmt.Printf("api version: %s\n", version.APIVersion)
	fmt.Printf("version:     %s\n", version.Version)
	return nil
}
