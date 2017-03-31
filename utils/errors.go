package utils

import (
	"fmt"
	"net/http"
	"errors"

	"github.com/vapor-ware/vesh/client"
)

func DeviceNotFoundErr(res Result) error {
	vc := client.New()
	failure := &client.ErrorResponse{}
	path := fmt.Sprintf("%s/%s/%s", res.RackID, res.BoardID, res.DeviceType) // FIXME: Do the lookup by device_type or device_id
	// FIXME: NOOOOOOOOOOOOOOOO
	var prefix string
	switch res.DeviceType {
	case "fan_speed":
		prefix = "fan/"
	case "system":
		prefix = "power/"
	case "led":
		prefix = res.DeviceType + "/"
	case "temperature":
		prefix = "temperature/"
	}
	resp, err := vc.Sling.New().Path(prefix).Get(path).Receive(nil, failure) // FIXME: The path with break if device_id doesn't match devicepath
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("Error Code: %+v\nError: %+v", failure.HttpCode, failure.Message))
	}
	return nil
}
