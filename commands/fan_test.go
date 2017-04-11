package commands

import (
	"testing"

	"github.com/vapor-ware/vesh/client"
	"github.com/vapor-ware/vesh/utils"
)

func init() {
	client.Config("demo.vapor.io")
}

func TestListFan(t *testing.T) {
	err := PrintListFan()
	if err != nil {
		t.Error(err)
	}
}

func TestGetFan(t *testing.T) {
	args := utils.GetDeviceArgs{RackID: "rack_1", BoardID: "40000001"}
	err := PrintGetFan(args)
	if err != nil {
		t.Error(err)
	}
}
