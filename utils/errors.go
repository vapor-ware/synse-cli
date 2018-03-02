package utils

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/client"
)

// DeviceNotFoundErr takes in a specific device and queries the endpoint for that
// device directly to return the error generated. It attempts to determine the
// correct path to use from the DeviceType, but this is not always correct and can
// cause overlaps.
func DeviceNotFoundErr(res Result) error {
	failure := &client.ErrorResponse{}
	path := fmt.Sprintf("%s/%s/%s", res.RackID, res.BoardID, res.DeviceType) // FIXME: Do the lookup by device_type or device_id
	// FIXME: NOOOOOOOOOOOOOOOO
	var prefix string
	switch res.DeviceType {
	case "fan":
		prefix = "fan/"
	case "system":
		prefix = "power/"
	case "led":
		prefix = res.DeviceType + "/"
	case "temperature":
		prefix = "temperature/"
	}
	resp, err := client.New().Path(prefix).Get(path).Receive(nil, failure) // FIXME: The path with break if device_id doesn't match devicepath
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Error Code: %+v\nError: %+v", failure.HTTPCode, failure.Message)
	}
	return nil
}

// CommandHandler wraps the error logger to log results from a called command.
func CommandHandler(c *cli.Context, err error) error {
	if err != nil {
		log.WithFields(log.Fields{
			// TODO: Add full options and flags as fields.
			"command": c.Command.Name,
		}).Error(err)
		return err
	}
	return nil
}
