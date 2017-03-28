// Commands provides the skeleton structure for the
// commands, subcmmands, and flags available to the cli. It also provides
// basic input parsing and error checking.
/*
Commands defines the commands, subcommands, and flags used in app.Cli to form
the structure of the CLI. Definitions, usage strings, help text, and flags are
deligated to app.Cli. The `Action:` field gives the function called when each
command is run.

Commands are broken up into separate files by category matching the category listed
in their descriptions. Top level commands accessible to the user, and matching
the definitions given below are included in the commands package in this directory.
Unless otherwise stated, each command definition should contain the following:

	command:

		- Data struct:

			This data struct usually forms the definition of returned data from a
			querying function call. Typically it matches the `json` fields returned
			from the HTTP GET request. When a request is made, the response is stored
			in a copy of this struct and pointers are used by the command functions
			to access the data.

		- Listing function:

			Most commands have some form of a "list" function that returns most (or all)
			of the elements being queried for that the given backend contains. For
			example running `assets fan list` will return information on all fans on
			all boards and racks.

		- Geting/Setting function(s):

			Most commands also contain a specific "get" function that returns
			information about a specific device. Typically, depending on the type of
			device, and how nested it is, a specific device is given by the rack and
			board id on which it is located (as well as the `device_type` that
			corresponds to the device being queried). These commands may or may not
			have the advantage of not requiring a full device list to be built by
			walking the tree, thereby saving return time.
			If the field in question allows bi-directional interaction a second
			form of this function may be present. This "set" function allows the value
			fields within this device to be set. Like the "get" command it takes a
			specifier, usually in the form of a rack and board id. It may also take
			one or more values to be set. Which values are possible, and in what order
			they should appear are typically given in the "Usage:" string for the
			specific command.

		- Printing functions:

			Unless otherwise specified, the functions for each command do not print
			their output. Each command should have an accompanying "print" function
			that will take the output of the corresponding command and format it
			properly, then print it to stdout. A tablewriter is typically used to
			organize multiple rows of data. Unless otherwise specified tables are
			formatted as close to Markdown table format as possible. Some functions
			may also have a `--raw` output mode, which will print the output with
			minimal formatting. By separating printing from a command itself, commands
			can be used internally without presenting to the user.

Typically the `Action:` field below runs a wrapper function that does minimal
input and error parsing before calling the associated "print" function. The
"print" function is responsible for printing output to the user, only errors
(if any) are returned to the calling function.

NOTE: Some commands may use a progress bar to specify progress during long
queries (usually walking a device tree). This is experimental and not implemented
everywhere.

Unless otherwise specified all errors should be fatal since each command
is stateless and called once during each run.
*/
package commands

import (
	"strconv" // I don't like having to use this here
	"fmt"

	"github.com/vapor-ware/vesh/client"
	"github.com/vapor-ware/vesh/utils"

	"github.com/urfave/cli"

)

// Commands provides the global list of commands to app.cli.
// Definitions, usage information, and executed functions are given.
var Commands = []cli.Command{
	{
		Name:    "status",
		Aliases: []string{"stat"},
		Usage:   "Get the status of the current deployment",
		Action: func(c *cli.Context) error {
			req := client.New()
			err := TestAPI(req)
			if err != nil {
				fmt.Println(err)
				return err
			}
			return nil
		},
	},
	{
		Name:  "scan",
		Usage: "Scan the infrastructure and display device summary",
		Action: func(c *cli.Context) error {
			req := client.New()
			err := Scan(req)
			fmt.Println(err)
			if err != nil {
				fmt.Println(err)
				return err
			}
			return nil
		},
	},
	{
		Name:  "assets",
		Usage: "Manage and get information about physical devices",
		Subcommands: []cli.Command{
			{
				Name:     "hostname",
				Usage:    "Manage hostnames",
				Category: "assets",
				Subcommands: []cli.Command{
					{
						Name:     "list",
						Usage:    "List hostnames",
						Category: "hostname",
						Action: func(c *cli.Context) error {
							req := client.New()
							err := ListHostnames(req)
							if err != nil {
								fmt.Println(err)
								return err
							}
							return nil
						},
					},
					{
						Name:      "get",
						Usage:     "Get hostname for specific `device`",
						ArgsUsage: "<rack id> <board id>",
						Category:  "hostname",
						Action: func(c *cli.Context) error {
							req := client.New()
							format := []string{"%s", "%x"}
							errFormat := utils.InputCheckFormat(c, format)
							if c.Args().Present() && errFormat == nil {
								err := PrintGetHostname(req, c.Args().Get(0), c.Args().Get(1))
								if err != nil {
									fmt.Println(err)
									return err
								}
								return nil
							}
							if errFormat != nil {
								fmt.Println(errFormat)
								cli.ShowSubcommandHelp(c)
							}
							return errFormat // Fix this. Restructure error checking and responses.
						},
					},
				},
			},
			{
				Name:     "power",
				Usage:    "Manage power status",
				Category: "assets",
				Subcommands: []cli.Command{
					{
						Name:     "list",
						Usage:    "List power status",
						Category: "power",
						Action: func(c *cli.Context) error {
							req := client.New()
							if c.Args().Present() == false {
								err := PrintListPower(req)
								if err != nil {
									fmt.Println(err)
									return err
								}
								return nil
							}
							cli.ShowSubcommandHelp(c)
							return nil // Fix this. Restructure error checking and responses.
						},
					},
					{
						Name:      "get",
						Usage:     "Get power status for specific `device`",
						ArgsUsage: "<rack id> <board_id>",
						Category:  "power",
						Action: func(c *cli.Context) error {
							req := client.New()
							if c.Args().Present() == true {
								err := PrintGetPower(req, c.Args().Get(0), c.Args().Get(1))
								if err != nil {
									fmt.Println(err)
									return err
								}
								return nil
							}
							cli.ShowSubcommandHelp(c)
							return nil // Fix this. Restructure error checking and responses.
						},
					},
					{
						Name:      "set",
						Usage:     "Change the power status `on/off/cycle`",
						ArgsUsage: "<on/off/cycle>",
						Category:  "power",
						Action: func(c *cli.Context) error {
							req := client.New()
							if c.Args().Present() == true {
								err := PrintSetPower(req, c.Args().Get(0), c.Args().Get(1), c.Args().Get(2)) // Consider breaking some of these out into flags
								if err != nil {
									fmt.Println(err)
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
				Name:     "fan",
				Usage:    "Manage fans",
				Category: "assets",
				Subcommands: []cli.Command{
					{
						Name:     "list",
						Usage:    "List fans speeds",
						Category: "fans",
						Action: func(c *cli.Context) error {
							req := client.New()
							if c.Args().Present() != true {
								err := PrintListFan(req)
								if err != nil {
									fmt.Println(err)
									return err
								}
								return nil
							}
							cli.ShowSubcommandHelp(c)
							return nil // Fix this. Restructure error checking and responses.
						},
					},
					{
						Name:      "get",
						Usage:     "Get fan speed for specific `device`",
						ArgsUsage: "<rack id> <board id>",
						Category:  "fans",
						Action: func(c *cli.Context) error {
							req := client.New()
							if c.Args().Present() == true {
								err := PrintGetFan(req, c.Args().Get(0), c.Args().Get(1))
								if err != nil {
									fmt.Println(err)
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
				Name:     "temperature",
				Usage:    "Manage temperature",
				Category: "assets",
				Subcommands: []cli.Command{
					{
						Name:     "list",
						Usage:    "List temperatures",
						Category: "temperature",
						Action: func(c *cli.Context) error {
							req := client.New()
							if c.Args().Present() != true {
								err := PrintListTemp(req)
								if err != nil {
									fmt.Println(err)
									return err
								}
								return nil
							}
							cli.ShowSubcommandHelp(c)
							return nil // Fix this. Restructure error checking and responses.
						},
					},
					{
						Name:      "get",
						Usage:     "Get temperature for specific `device`",
						ArgsUsage: "<rack id> <board id>",
						Category:  "temperature",
						Action: func(c *cli.Context) error {
							req := client.New()
							if c.Args().Present() == true {
								err := PrintGetTemp(req, c.Args().Get(0), c.Args().Get(1))
								if err != nil {
									fmt.Println(err)
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
				Name:     "boot-target",
				Usage:    "Get or change the boot target",
				Category: "assets",
				Subcommands: []cli.Command{
					{
						Name:      "set",
						Usage:     "Set the boot target for specific `device`. Can be `pxe` `hdd` or `no-override`",
						ArgsUsage: "<rack id> <board id> <pxe/hdd/no-override>",
						Category:  "boot-target",
						Action: func(c *cli.Context) error {
							req := client.New()
							if c.Args().Present() == true {
								rack, _ := strconv.Atoi(c.Args().Get(0))                       // This kind of thing should be done in the specific command
								board, _ := strconv.Atoi(c.Args().Get(1))                      // Ditto
								err := SetCurrentBootTarget(req, rack, board, c.Args().Get(2)) // Consider breaking some of these out into flags
								if err != nil {
									fmt.Println(err)
									return err
								}
								return nil
							}
							cli.ShowSubcommandHelp(c)
							return nil // Fix this. Restructure error checking and responses.
						},
					},
					{
						Name:      "get",
						Usage:     "Get current boot target for specific `device`",
						ArgsUsage: "<rack id> <board id>",
						Category:  "boot-target",
						Action: func(c *cli.Context) error {
							req := client.New()
							if c.Args().Present() == true {
								err := PrintGetCurrentBootTarget(req, c.Args().Get(0), c.Args().Get(1))
								if err != nil {
									fmt.Println(err)
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
				Name:     "lights",
				Usage:    "Manage and change LED status",
				Category: "assets",
				Subcommands: []cli.Command{
					{
						Name:     "list",
						Usage:    "List LED status",
						Category: "lights",
						Action: func(c *cli.Context) error {
							req := client.New()
							if c.Args().Present() != true {
								err := PrintListLights(req)
								if err != nil {
									fmt.Println(err)
									return err
								}
								return nil
							}
							cli.ShowSubcommandHelp(c)
							return nil // Fix this. Restructure error checking and responses.
						},
					},
					{
						Name:      "get",
						Usage:     "Get LED status for specific `device`",
						ArgsUsage: "<rack id> <board id>",
						Category:  "lights",
						Action: func(c *cli.Context) error {
							req := client.New()
							if c.Args().Present() == true {
								err := PrintGetLight(req, c.Args().Get(0), c.Args().Get(1))
								if err != nil {
									fmt.Println(err)
									return err
								}
								return nil
							}
							cli.ShowSubcommandHelp(c)
							return nil // Fix this. Restructure error checking and responses.
						},
					},
					{
						Name:     "set",
						Usage:    "Change the status for a specific LED `on/off/blink`",
						Category: "lights",
						Flags: []cli.Flag{
							cli.StringFlag{
								Name:  "state",
								Usage: "Set state to `on/off/blink`",
							},
							cli.StringFlag{
								Name:  "color",
								Usage: "Set LED to `color` (3 byte base 16 hex)",
							},
							cli.StringFlag{
								Name:  "blink",
								Usage: "Set LED blink to `blink` or `steady`",
							},
						},
						Action: func(c *cli.Context) error {
							req := client.New()
							if c.Args().Present() == true && c.NArg() == 2 {
								rack, _ := strconv.Atoi(c.Args().Get(0))  // This kind of thing should be done in the specific command
								board, _ := strconv.Atoi(c.Args().Get(1)) // Ditto
								switch {
								case c.IsSet("state") == true:
									err := PrintSetLight(req, rack, board, c.String("state"), "state") // Consider breaking some of these out into flags
									if err != nil {
										fmt.Println(err)
										return err
									}
									return nil
								case c.IsSet("color"):
									err := PrintSetLight(req, rack, board, c.String("color"), "color") // Consider breaking some of these out into flags
									if err != nil {
										fmt.Println(err)
										return err
									}
									return nil
								case c.IsSet("blink"):
									err := PrintSetLight(req, rack, board, c.String("blink"), "blink") // Consider breaking some of these out into flags
									if err != nil {
										fmt.Println(err)
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
						Name:      "blink",
						Usage:     "Blink specific `LED` (alias for '--blink true') (NOT YET IMPLEMENTED)",
						ArgsUsage: "<rack id> <board id> <device id> <true/false>",
						Category:  "lights",
						Action:    nil,
					},
					{
						Name:      "color",
						Usage:     "Set a specific `LED` to `color` (alias for '--color <hex>') (NOT YET IMPLEMENTED)", // Consider removing
						ArgsUsage: "<rack id> <board id> <device id> <color>",
						Category:  "lights",
						Action:    nil,
					},
				},
			},
			{
				Name:     "location",
				Usage:    "Get the physical location of a `device` (NOT YET IMPLEMENTED)",
				Category: "assets",
				Subcommands: []cli.Command{
					{
						Name:     "set",
						Usage:    "Set the geographic location for specific `device`",
						Category: "location",
						Action:   nil,
					},
					{
						Name:     "get",
						Usage:    "Get current geographic location of a specific `device`",
						Category: "location",
						Action:   nil,
					},
					{
						Name:     "map",
						Usage:    "Plot the geographic location of a specific `device` on a mapping service",
						Category: "location",
						Action:   nil,
					},
				},
			},
			{
				Name:     "find",
				Usage:    "Blink the LEDs on a specific `device` for 10 seconds to locate it",
				Category: "assets",
				Action:   nil,
			},
		},
	},
	{
		Name:  "zones",
		Usage: "List available zones (NOT YET IMPLEMENTED)",
		//Action:, TBD
	},
	{
		Name:  "racks",
		Usage: "List available racks within a given `zone` (or all zones if none is specified) (NOT YET IMPLEMENTED)",
		//Action:, TBD
	},
	{
		Name:  "health",
		Usage: "Check health for a given `zone`, `rack`, or `device` (NOT YET IMPLEMENTED)",
		//Action:, TBD
	},
	{
		Name:  "notifications",
		Usage: "List notifications for a given `zone`, `rack`, or `device` (NOT YET IMPLEMENTED)",
		//Action:, TBD
		Subcommands: []cli.Command{
			{
				Name:  "clear",
				Usage: "Clear notifications (`all` or `id`)",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "all",
						Usage: "Clear all notifications",
					},
					cli.StringFlag{
						Name:  "id",
						Usage: "Clear notifications for a specific `id`",
						//Destination: &clearNotificationsID,
					},
				},
			},
		},
	},
	{
		Name:  "load",
		Usage: "Get the load by specific metric (NOT YET IMPLEMENTED)",
		//Action:, TBD
		Subcommands: []cli.Command{
			{
				Name:     "power",
				Usage:    "Show power consumption",
				Category: "load",
				Action:   nil,
			},
			{
				Name:     "memory",
				Usage:    "Show memory usage",
				Category: "load",
				Action:   nil,
			},
			{
				Name:     "power",
				Usage:    "Show temprature",
				Category: "load",
				Action:   nil,
			},
			{
				Name:     "cpu",
				Usage:    "Show cpu usage",
				Category: "load",
				Action:   nil,
			},
			{
				Name:     "application",
				Usage:    "Show load by application",
				Category: "load",
				Action:   nil,
			},
		},
	},
	{
		Name:  "provision",
		Usage: "Get (un)provisioned servers and provision new servers (NOT YET IMPLEMENTED)",
		Subcommands: []cli.Command{
			{
				Name:     "new",
				Usage:    "Provision unprovisioned servers",
				Category: "provision",
				Action:   nil,
			},
			{
				Name:     "deprovision",
				Usage:    "deprovision previously provisioned servers",
				Category: "provision",
				Action:   nil,
			},
			{
				Name:     "list",
				Usage:    "list provisioned or deprovisioned servers",
				Category: "provision",
				Action:   nil,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "provisioned",
						Usage: "list provisioned servers",
					},
					cli.StringFlag{
						Name:  "unprovisioned",
						Usage: "list unprovisioned servers",
					},
				},
			},
		},
	},
}
