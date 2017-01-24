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
        },
      },
      {
        Name: "power",
        Usage: "Manage power status",
        Category: "assets",
        Subcommands: []cli.Command{
          {
            Name: "list",
            Usage: "List power status",
            Category: "power",
            Action: cmdListPower(c *cli.Context),
          },
          {
            Name: "get",
            Usage: "Get power status for specific `device`",
            Category: "power",
            Action: cmdGetPower(c *cli.Context),
          },
          {
            Name: "set",
            Usage: "Change the power status `on/off/cycle`",
            Category: "power",
            Destination: &powerSet,
            Action: cmdSetPower(c *cli.Context, powerSet string),
          },
        }
      },
      {
        Name: "fan",
        Usage: "Manage fans",
        Category: "assets",
        Subcommands: []cli.Command{
          {
            Name: "list",
            Usage: "List fans speeds",
            Category: "fans",
            Action: cmdListFan(c *cli.Context),
          },
          {
            Name: "get",
            Usage: "Get fan speed for specific `device`",
            Category: "fans",
            Action: cmdGetFan(c *cli.Context),
          },
        },
      },
      {
        Name: "temperature",
        Usage: "Manage temperature",
        Category: "assets",
        Subcommands: []cli.Command{
          {
            Name: "list",
            Usage: "List temperatures",
            Category: "temperature",
            Action: cmdListTemp(c *cli.Context),
          },
          {
            Name: "get",
            Usage: "Get temperature for specific `device`",
            Category: "temperature",
            Action: cmdGetTemp(c *cli.Context),
          },
        },
      },
      {
        Name: "boot-target",
        Usage: "Get or change the boot target",
        Category: "assets",
        Subcommands: []cli.Command{
          {
            Name: "set",
            Usage: "Set the boot target for specific `device`. Can be `pxe` `hdd` or `no-override`",
            Category: "boot-target",
            Value: &bootTargetValue,
            Action: cmdSetBootTarget(c *cli.Context, bootTargetValue string),
          },
          {
            Name: "get",
            Usage: "Get current boot target for specific `device`",
            Category: "boot-target",
            Action: cmdGetBootTarget(c *cli.Context),
          },
        },
      },
      {
        Name: "lights",
        Usage: "Manage and change LED status",
        Category: "assets",
        Subcommands: []cli.Command{
          {
            Name: "list",
            Usage: "List LED status",
            Category: "lights",
            Action: cmdListLed(c *cli.Context),
          },
          {
            Name: "get",
            Usage: "Get LED status for specific `device`",
            Category: "lights",
            Action: cmdGetLed(c *cli.Context),
          },
          {
            Name: "set",
            Usage: "Change the status for a specific LED `on/off/blink`",
            Category: "lights",
            Destination: &ledSet,
            Action: cmdSetLed(c *cli.Context, ledSet string),
          },
          {
            Name: "blink",
            Usage: "Blink specific `LED`",
            Category: "lights",
            Action: cmdBlinkled(c *cli.Context),
          },
          {
            Name: "color",
            Usage: "Set a specific `LED` to `color`",
            Category: "lights",
            Destination: &ledColor,
            Action: cmdColorLed(c *cli.Context, ledColor string),
          },
        }
      },
      {
        Name: "location",
        Usage: "Get the physical location of a `device`",
        Category: "assets",
        Subcommands: []cli.Command{
          {
            Name: "set",
            Usage: "Set the geographic location for specific `device`",
            Category: "location",
            Value: &bootTargetValue,
            Action: cmdSetBootTarget(c *cli.Context, bootTargetValue string),
          },
          {
            Name: "get",
            Usage: "Get current geographic location of a specific `device`",
            Category: "location",
            Action: cmdGetBootTarget(c *cli.Context),
          },
          {
            Name: "map",
            Usage: "Plot the geographic location of a specific `device` on a mapping service",
            Category: "location",
            Action: cmdMapLocation(c *cli.Context),
          },
        },
      },
      {
        Name: "find"
        Usage: "Blink the LEDs on a specific `device` for 10 seconds to locate it",
        Category: "assets",
        Action: cmdFindDevice(c *cli.Context),
      },
    },
  }
}
