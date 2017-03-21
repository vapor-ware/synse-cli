// Vesh provides a cli to vapor.io infrastructure. Specifically it allows access
// to the openDCRE (http://www.vapor.io/opendcre/) REST API for running commands
// against connected infrastructure and devices.
//
// For usage information please see the help text or the README in this
// repository.
package main

import (
	"os"

	"github.com/vapor-ware/vesh/client"
	"github.com/vapor-ware/vesh/commands"

	"github.com/urfave/cli"
)

// Main creates a new instance of cli.app (using https://github.com/urfave/cli)
// and sets the default configuration.
func main() {
	app := cli.NewApp()
	app.Name = "vesh"
	app.Usage = "Vapor Edge Shell"
	app.Version = "0.0.1"

	app.Commands = commands.Commands
	//app.CommandNotFound = commands.CommandNotFound
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "Enable debug mode",
		},
		cli.StringFlag{
			EnvVar: "VESH_CONFIG_FILE",
			Name:   "config, c",
			Usage:  "Path to config `file`",
		},
		cli.StringFlag{
			EnvVar:      "VESH_HOST",
			Name:        "host",
			Value:       "demo.vapor.io", // This is temporary
			Usage:       "Address of `Vapor Host`",
			Destination: &client.VeshHostPtr,
		},
	}

	app.Run(os.Args)

}
