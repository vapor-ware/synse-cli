package commands
import (
  "fmt"
  "net/http"

  "github.com/vapor-ware/vesh/client"
)

const testpath = "test"

type APIStatus struct {
  Status string `json:"status"`
}

func TestAPI(vc *client.VeshClient) error {
  status := &APIStatus{}
  resp, err := vc.Sling.New().Get(testpath).ReceiveSuccess(status)
  if err != nil {
    return err
  }
  if resp.StatusCode != http.StatusOK {
    return err
  }
  fmt.Println("API reported status ok")
  return nil
}
