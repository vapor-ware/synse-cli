package server

import (
	"fmt"
	"net/http"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/client"
	"github.com/vapor-ware/synse-cli/scheme"
	"github.com/vapor-ware/synse-cli/utils"
)

// versionBase is the base URI for the "version" route.
const versionBase = "version"

// VersionCommand is the CLI command for Synse Server's "version" API route.
var VersionCommand = cli.Command{
	Name:     "version",
	Usage:    "Get the version of the active host",
	Category: "Synse Server Actions",
	Action: func(c *cli.Context) error {
		return utils.CommandHandler(c, cmdVersion(c))
	},
}

// cmdVersion is the action for the VersionCommand. It makes an "version" request
// against the active Synse Server instance.
func cmdVersion(c *cli.Context) error {
	version := &scheme.Version{}
	resp, err := client.NewUnversioned().Get(versionBase).ReceiveSuccess(version)
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
