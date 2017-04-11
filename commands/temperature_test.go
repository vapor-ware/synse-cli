package commands

import (
  "testing"

  "github.com/vapor-ware/vesh/utils"
  "github.com/vapor-ware/vesh/client"
)

func init(){
  client.Config("demo.vapor.io")
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
