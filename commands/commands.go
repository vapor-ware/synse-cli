package commands

import (
  "strconv" // I don't like having to use this here

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
    Usage: "Scan the infrastructure and display device summary",
    Action: func (c *cli.Context) error {
      req := client.New()
      _, err := Scan(req)
      if err != nil {
        return err
      }
      return nil
    },
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
            Action: func (c *cli.Context) error {
              req := client.New()
              err := ListHostnames(req)
              if err != nil {
                return err
              }
              return nil
            },
          },
          {
            Name: "get",
            Usage: "Get hostname for specific `device`",
            Category: "hostname",
            Flags: []cli.Flag{
              cli.BoolFlag{
                Name: "raw",
                Usage: "Only output a space separated list of hostnames and IP addresses",
              },
            },
            Action: func (c *cli.Context) error {
              req := client.New()
              if c.Args().Present() {
                err := PrintGetHostname(req, "rack_whatever", c.Args().Get(0), "system", c.Bool("raw")) //stop hardcoding this. Lookup?
                if err != nil {
                  return err
                }
                return nil
              }
              cli.ShowSubcommandHelp(c)
              return nil // Fix this. Restructure error checking and responses.
            },
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
            Action: func (c *cli.Context) error {
              req := client.New()
              if c.Args().Present() == false {
                err := PrintListPower(req)
                if err != nil {
                  return err
                }
                return nil
              }
              cli.ShowSubcommandHelp(c)
              return nil // Fix this. Restructure error checking and responses.
            },
          },
          {
            Name: "get",
            Usage: "Get power status for specific `device`",
            Category: "power",
            Action: func (c *cli.Context) error{
              req := client.New()
              if c.Args().Present() == true {
                err := PrintGetPower(req, c.Args().Get(0), c.Args().Get(1))
                if err != nil {
                  return err
                }
                return nil
              }
              cli.ShowSubcommandHelp(c)
              return nil // Fix this. Restructure error checking and responses.
            },
          },
          {
            Name: "set",
            Usage: "Change the power status `on/off/cycle`",
            Category: "power",
            Action: func (c *cli.Context) error {
              req := client.New()
              if c.Args().Present() == true {
                err := PrintSetPower(req, c.Args().Get(0), c.Args().Get(1), c.Args().Get(2)) // Consider breaking some of these out into flags
                if err != nil {
                  return err
                }
                return nil
              }
              cli.ShowSubcommandHelp(c)
              return nil // Fix this. Restructure error checking and responses.
            },
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
            Action: func (c *cli.Context) error {
              req := client.New()
              if c.Args().Present() != true {
                err := PrintListFan(req)
                if err != nil {
                  return err
                }
                return nil
              }
              cli.ShowSubcommandHelp(c)
              return nil // Fix this. Restructure error checking and responses.
            },
          },
          {
            Name: "get",
            Usage: "Get fan speed for specific `device`",
            Category: "fans",
            Action: func (c *cli.Context) error {
              req := client.New()
              if c.Args().Present() == true {
                err := PrintGetFan(req, c.Args().Get(0), c.Args().Get(1))
                if err != nil {
                  return err
                }
                return nil
              }
              cli.ShowSubcommandHelp(c)
              return nil // Fix this. Restructure error checking and responses.
            },
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
            Action: func (c *cli.Context) error {
              req := client.New()
              if c.Args().Present() != true {
                err := PrintListTemp(req)
                if err != nil {
                  return err
                }
                return nil
              }
              cli.ShowSubcommandHelp(c)
              return nil // Fix this. Restructure error checking and responses.
            },
          },
          {
            Name: "get",
            Usage: "Get temperature for specific `device`",
            Category: "temperature",
            Action: func (c *cli.Context) error {
              req := client.New()
              if c.Args().Present() == true {
                err := PrintGetTemp(req, c.Args().Get(0), c.Args().Get(1))
                if err != nil {
                  return err
                }
                return nil
              }
              cli.ShowSubcommandHelp(c)
              return nil // Fix this. Restructure error checking and responses.
            },
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
            Action: func (c *cli.Context) error {
              req := client.New()
              if c.Args().Present() == true {
                rack, _ := strconv.Atoi(c.Args().Get(0)) // This kind of thing should be done in the specific command
                board, _ := strconv.Atoi(c.Args().Get(1))// Ditto
                err := SetCurrentBootTarget(req, rack, board, c.Args().Get(2)) // Consider breaking some of these out into flags
                if err != nil {
                  return err
                }
                return nil
              }
              cli.ShowSubcommandHelp(c)
              return nil // Fix this. Restructure error checking and responses.
            },
          },
          {
            Name: "get",
            Usage: "Get current boot target for specific `device`",
            Category: "boot-target",
            Action: func (c *cli.Context) error {
              req := client.New()
              if c.Args().Present() == true {
                err := PrintGetCurrentBootTarget(req, c.Args().Get(0), c.Args().Get(1))
                if err != nil {
                  return err
                }
                return nil
              }
              cli.ShowSubcommandHelp(c)
              return nil // Fix this. Restructure error checking and responses.
            },
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
            Action: func (c *cli.Context) error {
              req := client.New()
              if c.Args().Present() != true {
                err := PrintListLights(req)
                if err != nil {
                  return err
                }
                return nil
              }
              cli.ShowSubcommandHelp(c)
              return nil // Fix this. Restructure error checking and responses.
            },
          },
          {
            Name: "get",
            Usage: "Get LED status for specific `device`",
            Category: "lights",
            Action: func (c *cli.Context) error {
              req := client.New()
              if c.Args().Present() == true {
                err := PrintGetLight(req, c.Args().Get(0), c.Args().Get(1))
                if err != nil {
                  return err
                }
                return nil
              }
              cli.ShowSubcommandHelp(c)
              return nil // Fix this. Restructure error checking and responses.
            },
          },
          {
            Name: "set",
            Usage: "Change the status for a specific LED `on/off/blink`",
            Category: "lights",
            Flags: []cli.Flag{
              cli.StringFlag{
                Name: "state",
                Usage: "Set state to `on/off/blink`",
              },
              cli.StringFlag{
                Name: "color",
                Usage: "Set LED to `color` (3 byte base 16 hex)",
              },
              cli.StringFlag{
                Name: "blink",
                Usage: "Set LED blink to `blink` or `steady`",
              },
            },
            Action: func (c *cli.Context) error {
              req := client.New()
              if c.Args().Present() == true {
                rack, _ := strconv.Atoi(c.Args().Get(0)) // This kind of thing should be done in the specific command
                board, _ := strconv.Atoi(c.Args().Get(1))// Ditto
                switch {
                case c.String("state") != "":
                  err := PrintSetLight(req, rack, board, c.String("state"), "state") // Consider breaking some of these out into flags
                  if err != nil {
                    return err
                  }
                  return nil
                case c.String("color") != "":
                  err := PrintSetLight(req, rack, board, c.String("color"), "color") // Consider breaking some of these out into flags
                  if err != nil {
                    return err
                  }
                  return nil
                case c.String("blink") != "":
                  err := PrintSetLight(req, rack, board, c.String("blink"), "blink") // Consider breaking some of these out into flags
                  if err != nil {
                    return err
                  }
                  return nil
                }
              }
              cli.ShowSubcommandHelp(c)
              return nil // Fix this. Restructure error checking and responses.
            },
          },
          {
            Name: "blink",
            Usage: "Blink specific `LED` (alias for '--blink true') (NOT YET IMPLEMENTED)",
            Category: "lights",
            Action: nil,
          },
          {
            Name: "color",
            Usage: "Set a specific `LED` to `color` (alias for '--color <hex>') (NOT YET IMPLEMENTED)", // Consider removing
            Category: "lights",
            Action: nil,
          },
        },
      },
      {
        Name: "location",
        Usage: "Get the physical location of a `device` (NOT YET IMPLEMENTED)",
        Category: "assets",
        Subcommands: []cli.Command{
          {
            Name: "set",
            Usage: "Set the geographic location for specific `device`",
            Category: "location",
            Action: nil,
          },
          {
            Name: "get",
            Usage: "Get current geographic location of a specific `device`",
            Category: "location",
            Action: nil,
          },
          {
            Name: "map",
            Usage: "Plot the geographic location of a specific `device` on a mapping service",
            Category: "location",
            Action: nil,
          },
        },
      },
      {
        Name: "find",
        Usage: "Blink the LEDs on a specific `device` for 10 seconds to locate it",
        Category: "assets",
        Action: nil,
      },
    },
  },
  {
    Name: "zones",
    Usage: "List available zones (NOT YET IMPLEMENTED)",
    //Action:, TBD
  },
  {
    Name: "racks",
    Usage: "List available racks within a given `zone` (or all zones if none is specified) (NOT YET IMPLEMENTED)",
    //Action:, TBD
  },
  {
    Name: "health",
    Usage: "Check health for a given `zone`, `rack`, or `device` (NOT YET IMPLEMENTED)",
    //Action:, TBD
  },
  {
    Name: "notifications",
    Usage: "List notifications for a given `zone`, `rack`, or `device` (NOT YET IMPLEMENTED)",
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
            //Destination: &clearNotificationsID,
          },
        },
      },
    },
  },
  {
    Name: "load",
    Usage: "Get the load by specific metric (NOT YET IMPLEMENTED)",
    //Action:, TBD
    Subcommands: []cli.Command{
      {
        Name: "power",
        Usage: "Show power consumption",
        Category: "load",
        Action: nil,
      },
      {
        Name: "memory",
        Usage: "Show memory usage",
        Category: "load",
        Action: nil,
      },
      {
        Name: "power",
        Usage: "Show temprature",
        Category: "load",
        Action: nil,
      },
      {
        Name: "cpu",
        Usage: "Show cpu usage",
        Category: "load",
        Action: nil,
      },
      {
        Name: "application",
        Usage: "Show load by application",
        Category: "load",
        Action: nil,
      },
    },
  },
  {
    Name: "provision",
    Usage: "Get (un)provisioned servers and provision new servers (NOT YET IMPLEMENTED)",
    Subcommands: []cli.Command{
      {
        Name: "new",
        Usage: "Provision unprovisioned servers",
        Category: "provision",
        Action: nil,
      },
      {
        Name: "deprovision",
        Usage: "deprovision previously provisioned servers",
        Category: "provision",
        Action: nil,
      },
      {
        Name: "list",
        Usage: "list provisioned or deprovisioned servers",
        Category: "provision",
        Action: nil,
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
  },
}
