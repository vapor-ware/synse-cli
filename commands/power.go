package commands
import (
  "os"
  //"fmt"
  "reflect"
  "strconv"

  "github.com/vapor-ware/vesh/client"
  //"github.com/vapor-ware/vesh/utils"

  "github.com/olekukonko/tablewriter"
)

const powerpath = "power/"
const device_id = "power"

type powerResponse struct {
  InputPower float64 `json:"input_power"`
  OverCurrent bool `json:"over_current"`
  PowerOK bool `json:"power_ok"`
  PowerStatus string `json:"power_status"`
}

func ListPower(vc *client.VeshClient) ([][]string, error) {
  scanResponse, _ := ScanOnly(vc) // Add error reporting
  scanResponsePtr := reflect.ValueOf(&scanResponse.Racks)
  scanResponseValuePtr := scanResponsePtr.Elem()
  fulltable := make([][]string, 0)
  totalruns := 0
  for i := 0; i < scanResponseValuePtr.Len(); i++ {
    boardsPtr := reflect.ValueOf(&scanResponse.Racks[i].Boards)
    boardsValuePtr := boardsPtr.Elem()
    for j := 0; j < boardsValuePtr.Len(); j++ {
      tablerow := make([]string, 0)
      tablerow = append(tablerow, scanResponse.Racks[i].Boards[j].BoardID)
      rack_id := scanResponse.Racks[i].RackID
      board_id := scanResponse.Racks[i].Boards[j].BoardID
      responseData := &powerResponse{}
      resp, err := vc.Sling.New().Path(powerpath).Path(rack_id + "/").Path(board_id + "/").Get(device_id).ReceiveSuccess(responseData) // Add error reporting
      if resp.StatusCode != 200 { // This is not what I meant by "error reporting"
        return nil, err
      }
      tablerow = append(tablerow, strconv.FormatFloat(responseData.InputPower, 'G', -1, 64))
      tablerow = append(tablerow, strconv.FormatBool(responseData.PowerOK))
      fulltable = append(fulltable, nil)
      fulltable[totalruns] = make([]string, 0)
      fulltable[totalruns] = append(fulltable[totalruns], tablerow...)
      totalruns ++
    }
  }
  return fulltable, nil
  //return nil, scanerr //fix with proper error
}

func PrintListPower(vc *client.VeshClient) error {
  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"Board", "Input Power", "Power Ok?"})
  table.SetBorder(false)
  powerList, _ := ListPower(vc) // Add error reporting
  table.AppendBulk(powerList)
  table.Render()
  return nil
}
