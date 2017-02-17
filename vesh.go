package main

import (
	"os"

	"github.com/vapor-ware/vesh/commands"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "vesh"
	app.Usage = "Vapor Edge Shell"
	app.Version = "0.0.1"

	app.Commands = commands.Commands
	//app.CommandNotFound = commands.CommandNotFound

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
			EnvVar: "VESH_HOST",
			Name:   "host",
			Usage:  "Address of `Vapor Host`",
		},
	}

	app.Run(os.Args)

}
