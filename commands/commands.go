package commands

import (
  "github.com/urfave/cli"
  "github.com/vapor-ware/vesh/client"
)

var Commands = []cli.Command{
  {
    Name: "status",
    Aliases: []string{"stat"},
    Usage: "Get the status of the current deployment",
    Action: func (c *cli.Context) error {
      req := client.New()
      err := TestAPI(req)
      if err != nil {
        return err
      }
      return nil
    },
  },
  {
    Name: "scan",
    Usage: "Get the status of the current deployment",
    Action: func (c *cli.Context) error {
      req := client.New()
      err := Scan(req)
      if err != nil {
        return err
      }
      return nil
    },
  },
/*  {
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
            Action: cmdListHostname,
          },
          {
            Name: "get",
            Usage: "Get hostname for specific `device`",
            Category: "hostname",
            Action: cmdGetHostname,
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
            Action: cmdListPower,
          },
          {
            Name: "get",
            Usage: "Get power status for specific `device`",
            Category: "power",
            Action: cmdGetPower,
          },
          {
            Name: "set",
            Usage: "Change the power status `on/off/cycle`",
            Category: "power",
            Action: cmdSetPower,
          },
        },
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
            Action: cmdListFan,
          },
          {
            Name: "get",
            Usage: "Get fan speed for specific `device`",
            Category: "fans",
            Action: cmdGetFan,
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
            Action: cmdListTemp,
          },
          {
            Name: "get",
            Usage: "Get temperature for specific `device`",
            Category: "temperature",
            Action: cmdGetTemp,
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
            Action: cmdSetBootTarget,
          },
          {
            Name: "get",
            Usage: "Get current boot target for specific `device`",
            Category: "boot-target",
            Action: cmdGetBootTarget,
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
            Action: cmdListLed,
          },
          {
            Name: "get",
            Usage: "Get LED status for specific `device`",
            Category: "lights",
            Action: cmdGetLed,
          },
          {
            Name: "set",
            Usage: "Change the status for a specific LED `on/off/blink`",
            Category: "lights",
            Action: cmdSetLed,
          },
          {
            Name: "blink",
            Usage: "Blink specific `LED`",
            Category: "lights",
            Action: cmdBlinkled,
          },
          {
            Name: "color",
            Usage: "Set a specific `LED` to `color`",
            Category: "lights",
            Action: cmdColorLed,
          },
        },
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
            Action: cmdSetBootTarget,
          },
          {
            Name: "get",
            Usage: "Get current geographic location of a specific `device`",
            Category: "location",
            Action: cmdGetBootTarget,
          },
          {
            Name: "map",
            Usage: "Plot the geographic location of a specific `device` on a mapping service",
            Category: "location",
            Action: cmdMapLocation,
          },
        },
      },
      {
        Name: "find",
        Usage: "Blink the LEDs on a specific `device` for 10 seconds to locate it",
        Category: "assets",
        Action: cmdFindDevice,
      },
    },
  },
  {
    Name: "zones",
    Usage: "List available zones",
    //Action:, TBD
  },
  {
    Name: "racks",
    Usage: "List available racks within a given `zone` (or all zones if none is specified)",
    //Action:, TBD
  },
  {
    Name: "health",
    Usage: "Check health for a given `zone`, `rack`, or `device`",
    //Action:, TBD
  },
  {
    Name: "notifications",
    Usage: "List notifications for a given `zone`, `rack`, or `device`",
    //Action:, TBD
    Subcommands: []cli.Command{
      {
        Name: "clear",
        Usage: "Clear notifications (`all` or `id`)",
        Flags: []cli.Flag{
          cli.StringFlag{
            Name: "all",
            Usage: "Clear all notifications",
          },
          cli.StringFlag{
            Name: "id",
            Usage: "Clear notifications for a specific `id`",
            Destination: &clearNotificationsID,
          },
        },
      },
    },
  },
  {
    Name: "load",
    Usage: "Get the load by specific metric",
    //Action:, TBD
    Subcommands: []cli.Command{
      {
        Name: "power",
        Usage: "Show power consumption",
        Category: "load",
        Action: cmdShowPowerLoad,
      },
      {
        Name: "memory",
        Usage: "Show memory usage",
        Category: "load",
        Action: cmdShowMemoryLoad,
      },
      {
        Name: "power",
        Usage: "Show temprature",
        Category: "load",
        Action: cmdShowTempratureLoad,
      },
      {
        Name: "cpu",
        Usage: "Show cpu usage",
        Category: "load",
        Action: cmdShowCPULoad,
      },
      {
        Name: "application",
        Usage: "Show load by application",
        Category: "load",
        Action: cmdShowApplicationLoad,
      },
    },
  },
  {
    Name: "provision",
    Usage: "Get (un)provisioned servers and provision new servers",
    Subcommands: []cli.Command{
      {
        Name: "new",
        Usage: "Provision unprovisioned servers",
        Category: "provision",
        Action: cmdProvisionNew,
      },
      {
        Name: "deprovision",
        Usage: "deprovision previously provisioned servers",
        Category: "provision",
        Action: cmdDeprovision,
      },
      {
        Name: "list",
        Usage: "list provisioned or deprovisioned servers",
        Category: "provision",
        Action: cmdProvisionList,
        Flags: []cli.Flag{
          cli.StringFlag{
            Name: "provisioned",
            Usage: "list provisioned servers",
          },
          cli.StringFlag{
            Name: "unprovisioned",
            Usage: "list unprovisioned servers",
          },
        },
      },
    },
  },*/
}
