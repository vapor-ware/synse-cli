package main

import (
  "fmt"
  "os"
  "net/http"
  "github.com/urfave/cli"
)

var Commands = []cli.Command{
  {
    Name: "status",
    Aliases: []string{"stat"},
    Usage: "Get the status of the current deployment",
    Action:, //TBD
  },
  {
    Name: "assets",
    Usage: "Manage and get information about physical devices",
    Subcommands: []cli.Command{
      {
        Name: "hostname",
        Usage: "Manage hostnames",
        Category: "assets",
        Subcommands: []cli.Command{
          {
            Name: "list",
            Usage: "List hostnames",
            Category: "hostname",
            Action: cmdListHostname(c *cli.Context),
          },
          {
            Name: "get",
            Usage: "Get hostname for specific `device`",
            Category: "hostname",
            Action: cmdGetHostname(c *cli.Context),
          },
        }
      }
    }
  }
}
