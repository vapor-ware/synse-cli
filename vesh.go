// Vesh provides a cli to vapor.io infrastructure. Specifically it allows access
// to the openDCRE (http://www.vapor.io/opendcre/) REST API for running commands
// against connected infrastructure and devices.
//
// For usage information please see the help text or the README in this
// repository.
package main

import (
	"os"

	// "github.com/vapor-ware/vesh/client"
	"github.com/vapor-ware/vesh/commands"
	"github.com/vapor-ware/vesh/utils"

	"github.com/urfave/cli"
	log "github.com/Sirupsen/logrus"
)

// Main creates a new instance of cli.app (using https://github.com/urfave/cli)
// and sets the default configuration.
func main() {
	log.SetLevel(log.DebugLevel)

	app := cli.NewApp()
	app.Name = "vesh"
	app.Usage = "Vapor Edge Shell"
	app.Version = "0.0.1"
	app.Author = "Tim Fall <tim@vapor.io>"

	app.Commands = commands.Commands
	//app.CommandNotFound = commands.CommandNotFound
	app.EnableBashCompletion = true

	app.Before = func(cli *cli.Context) error {
		err := utils.ConstructConfig()
		return err
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "Enable debug mode",
			Destination: &utils.DebugFlag,
		},
		cli.StringFlag{
			EnvVar:      "VESH_CONFIG_FILE",
			Name:        "config, c",
			Usage:       "Path to config `file`",
			Destination: &utils.ConfigFilePath,
		},
		cli.StringFlag{
			EnvVar:      "VAPOR_HOST",
			Name:        "host",
			Value:       "demo.vapor.io", // This is temporary
			Usage:       "Address of `Vapor Host`",
			Destination: &utils.VaporHost,
		},
	}

	app.Run(os.Args)

}
