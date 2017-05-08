package commands

import (
	"testing"

	"github.com/vapor-ware/synse-cli/client"
	"github.com/vapor-ware/synse-cli/utils"
)

func init() {
	client.Config("35.185.201.100")
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
