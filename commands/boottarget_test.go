package commands

import (
	"testing"

	"github.com/vapor-ware/synse-cli/client"
	"github.com/vapor-ware/synse-cli/utils"
)

func init() {
	client.Config("demo.vapor.io")
}

func TestListBootTarget(t *testing.T) {
	err := PrintListBootTarget()
	if err != nil {
		t.Error(err)
	}
}

func TestGetBootTarget(t *testing.T) {
	args := utils.GetDeviceArgs{RackID: "rack_1", BoardID: "40000001"}
	err := PrintGetBootTarget(args)
	if err != nil {
		t.Error(err)
	}
}

func TestSetBootTarget(t *testing.T) {
	args := utils.SetBootTargetArgs{utils.GetDeviceArgs{RackID: "rack_1", BoardID: "40000001"}, "no-override"}
	err := SetBootTarget(args)
	if err != nil {
		t.Error(err)
	}
}
