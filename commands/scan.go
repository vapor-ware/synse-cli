package commands
import (
  "os"
  "fmt"
  "net/http"
  "reflect"

  "github.com/vapor-ware/vesh/client"

  "github.com/olekukonko/tablewriter"
)

const scanpath = "scan"

type scanResponse struct {
  Racks []struct {
    Boards []struct {
      BoardID string `json:"board_id"`
      Devices []struct {
        DeviceID string `json:"device_id"`
        DeviceType string `json:"device_type"`
      } `json:"devices"`
    } `json:"boards"`
    RackID string `json:"rack_id"`
  } `json:"racks"`
}

func walkRacks(sr *scanResponse) {

}

func walkBoards(sr *scanResponse) {

}

func Scan(vc *client.VeshClient) error {
  status := &scanResponse{}
  resp, err := vc.Sling.New().Get(scanpath).ReceiveSuccess(status)
  if err != nil {
    return err
  }
  if resp.StatusCode != http.StatusOK {
    return err
  }
  fmt.Println("API reported status ok")
  otherthingy := status.Racks[0].Boards[2]
  stPtr := reflect.ValueOf(&otherthingy.Devices)
  stotherPtr := stPtr.Elem()
  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"Device ID", "Device Type"})
  data := make([][]string, stotherPtr.Len())
  for i := 0; i < stotherPtr.Len(); i ++ {
    data[i] = make([]string, 2)
    stPtr2 := reflect.ValueOf(&otherthingy.Devices[i])
    stotherPtr2 := stPtr2.Elem()
    for j := 0; j < stotherPtr2.NumField(); j++ {
      data[i][j] = stotherPtr2.Field(j).String()
    }
  }
  table.AppendBulk(data)
  table.Render()
  return nil
}

func writetable()  {

}
