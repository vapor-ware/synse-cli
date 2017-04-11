package commands

import (
  "testing"

  "github.com/vapor-ware/vesh/client"
)

func init(){
  client.Config("demo.vapor.io")
}

func TestTestAPI(t *testing.T) {
  err := TestAPI()
  if err != nil {
    t.Error(err)
  }
}
