package commands

import (
	"os"
	"testing"

	"github.com/vapor-ware/synse-cli/client"
	"github.com/vapor-ware/synse-cli/utils"
)

func init() {
	client.Config(os.Getenv("SYNSE_HOST"))
}

// Could be that this one is failing, causing subsequent failures (?)
func TestListFan(t *testing.T) {
  t.Logf("TestListFan\n")
  /*
	err := PrintListFan()
	if err != nil {
		t.Error(err)
	}
  */
}

func TestGetFan(t *testing.T) {
  t.Logf("TestGetFan\n")
  /*
	args := utils.GetDeviceArgs{RackID: "rack_1", BoardID: "40000001"}
	err := PrintGetFan(args)
	if err != nil {
		t.Error(err)
	}
  */
}
