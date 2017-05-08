package commands

import (
	"testing"

	"github.com/vapor-ware/synse-cli/client"
	"github.com/vapor-ware/synse-cli/utils"
)

func init() {
	client.Config("35.185.201.100")
}

func TestListTemp(t *testing.T) {
	err := PrintListTemp()
	if err != nil {
		t.Error(err)
	}
}

func TestGetTemp(t *testing.T) {
	args := utils.GetDeviceArgs{RackID: "rack_1", BoardID: "40000001"}
	err := PrintGetTemp(args)
	if err != nil {
		t.Error(err)
	}
}
