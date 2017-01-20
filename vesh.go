package main

import (
  "fmt"
  "os"
  "net/http"
  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "vesh"
  app.Usage = "Vapor Edge Shell"
  app.Version = "0.0.1"
  app.Run(os.Args)

  app.Commands = commands.Commands
  app.CommandNotFound = commands.CommandNotFound

  app.Flags = []cli.Flag {
    cli.BoolFlag{
      Name: "debug, d",
      Usage: "Enable debug mode",
    },
    cli.StringFlag{
      EnvVar: "VESH_CONFIG_FILE",
      Name: "config, c",
      Usage: "Path to config file",
    },
  }

}
