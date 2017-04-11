package commands

import (
  "testing"

  "github.com/vapor-ware/vesh/utils"
  "github.com/vapor-ware/vesh/client"
)

func init(){
  client.Config("demo.vapor.io")
}

func TestListHostnames(t *testing.T) {
  err := ListHostnames()
  if err != nil {
    t.Error(err)
  }
}

func TestGetHostname(t *testing.T) {
  args := utils.GetDeviceArgs{RackID: "rack_1", BoardID: "40000001"}
  err := PrintGetHostname(args)
  if err != nil {
    t.Error(err)
  }
}
