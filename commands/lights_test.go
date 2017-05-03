package commands

import (
	"testing"

	"github.com/vapor-ware/synse-cli/client"
	"github.com/vapor-ware/synse-cli/utils"
)

func init() {
	client.Config("demo.vapor.io")
}

func TestListLights(t *testing.T) {
	err := PrintListLights()
	if err != nil {
		t.Error(err)
	}
}

func TestGetLight(t *testing.T) {
	args := utils.GetDeviceArgs{RackID: "rack_1", BoardID: "40000001"}
	err := PrintGetLight(args)
	if err != nil {
		t.Error(err)
	}
}

func TestSetLight(t *testing.T) {
	args := utils.SetLightsArgs{utils.GetDeviceArgs{RackID: "rack_1", BoardID: "40000001"}, "on", "000", "blink"}
	err := PrintSetLight(args)
	if err != nil {
		t.Error(err)
	}
}
