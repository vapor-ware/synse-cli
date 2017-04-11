package commands

import (
  "testing"

  "github.com/vapor-ware/vesh/utils"
  "github.com/vapor-ware/vesh/client"
)

func init(){
  client.Config("demo.vapor.io")
}

func TestListPower(t *testing.T) {
  err := PrintListPower()
  if err != nil {
    t.Error(err)
  }
}

func TestGetPower(t *testing.T) {
  args := utils.GetDeviceArgs{RackID: "rack_1", BoardID: "40000001"}
  err := PrintGetPower(args)
  if err != nil {
    t.Error(err)
  }
}

func TestSetPower(t *testing.T) {
  args := utils.SetPowerArgs{utils.GetDeviceArgs{RackID: "rack_1", BoardID: "40000001"}, "cycle"}
  err := PrintSetPower(args)
  if err != nil {
    t.Error(err)
  }
}
