package utils

import (
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/asaskevich/govalidator"
)

type GetDeviceArgs struct {
	RackID string `valid:"ascii,required"`
	BoardID string `valid:"ascii,required"`
}

type SetPowerArgs struct {
	GetDeviceArgs
	Value string `valid:"in(on|off|cycle)"`
}

type SetBootTargetArgs struct {
	GetDeviceArgs
	Value string `valid:"in(pxe|hdd|no-override)"`
}

type SetLightsArgs struct {
	GetDeviceArgs
	State string `valid:"in(on|off|blink)"`
	Color string `valid:"hexcolor"`
	Blink string `valid:"in(blink|steady)"`
}

func InputValid(c *cli.Context, v interface{}) error {
	if _, err := govalidator.ValidateStruct(v); err != nil {
		cli.ShowSubcommandHelp(c)

		log.WithFields(log.Fields{
			"error": err,
		}).Error("invalid input")

		return err
	}
	return nil
}
