package utils

import (
	log "github.com/Sirupsen/logrus"
	"github.com/asaskevich/govalidator"
	"github.com/urfave/cli"
)

// GetDeviceArgs contains the required arguments for the `get` command.
type GetDeviceArgs struct {
	RackID  string `valid:"ascii,required"`
	BoardID string `valid:"ascii,required"`
}

// SetPowerArgs contains the required arguments for the `power set` command.
type SetPowerArgs struct {
	GetDeviceArgs
	Value string `valid:"in(on|off|cycle)"`
}

// SetBootTargetArgs contains the required arguments for the `boot-target set` command.
type SetBootTargetArgs struct {
	GetDeviceArgs
	Value string `valid:"in(pxe|hdd|no-override)"`
}

// SetLightsArgs contains the required arguments for the `lights set` command.
type SetLightsArgs struct {
	GetDeviceArgs
	State string `valid:"in(on|off|blink)"`
	Color string `valid:"hexcolor"`
	Blink string `valid:"in(blink|steady)"`
}

// InputValid takes in arguments passed to the CLI and does loose comparison
// against expect value types to determine if they are acceptable values. Type
// matching is limited as the range of accepted types is small.
func InputValid(c *cli.Context, v interface{}) error {
	if _, err := govalidator.ValidateStruct(v); err != nil {
		err := cli.ShowSubcommandHelp(c)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("invalid input")
		}
		return err
	}
	return nil
}
