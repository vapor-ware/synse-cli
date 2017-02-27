package commands
import (
  "os"
  "fmt"
  "reflect"
  "strconv"

  "github.com/vapor-ware/vesh/client"
  "github.com/vapor-ware/vesh/utils"

  "github.com/olekukonko/tablewriter"
)

const lightspath = "led/"
const lightsdevicetype = "led"

type lightsResponse struct {
  State string `json:"led_state"`
  BlinkState string `json:"blink_state"`
  Color int16 `json:"color"`
}

func ListLights(vc *client.VeshClient) ([][]string, error) {
  scanResponse, _ := utils.UtilScanOnly() // Add error reporting
  scanResponsePtr := reflect.ValueOf(&scanResponse.Racks)
  scanResponseValuePtr := scanResponsePtr.Elem()
  fulltable := make([][]string, 0)
  totalruns := 0
  for i := 0; i < scanResponseValuePtr.Len(); i++ {
    boardsPtr := reflect.ValueOf(&scanResponse.Racks[i].Boards)
    boardsValuePtr := boardsPtr.Elem()
    for j := 0; j < boardsValuePtr.Len(); j++ {
      devicePtr := reflect.ValueOf(&scanResponse.Racks[i].Boards[j].Devices)
      devicesValuePtr := devicePtr.Elem()
      for k := 0; k < devicesValuePtr.Len(); k++ {
        deviceTypePtr := reflect.ValueOf(&scanResponse.Racks[i].Boards[j].Devices[k].DeviceType)
        deviceTypeValuePtr := deviceTypePtr.Elem()
        if deviceTypeValuePtr.String() == lightsdevicetype { // This may need to be expanded to other types
          tablerow := make([]string, 0)
          rack_id := scanResponse.Racks[i].RackID
          board_id := scanResponse.Racks[i].Boards[j].BoardID
          device_id := scanResponse.Racks[i].Boards[j].Devices[k].DeviceID
          tablerow = append(tablerow, rack_id)
          tablerow = append(tablerow, board_id)
          tablerow = append(tablerow, device_id)
          responseData := &lightsResponse{}
          resp, err := vc.Sling.New().Path(lightspath).Path(rack_id + "/").Path(board_id + "/").Get(device_id).ReceiveSuccess(responseData) // Add error reporting
          if resp.StatusCode != 200 { // This is not what I meant by "error reporting"
            fmt.Println(vc)
            fmt.Println(resp)
            return nil, err
          }
          tablerow = append(tablerow, responseData.State)
          fulltable = append(fulltable, nil)
          fulltable[totalruns] = make([]string, 0)
          fulltable[totalruns] = append(fulltable[totalruns], tablerow...)
          totalruns ++
        }
      }
    }
  }
  return fulltable, nil
}

func PrintListLights(vc *client.VeshClient) error {
  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"Rack", "Board", "Device", "LED State"})
  table.SetBorder(false)
  table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
  table.SetCenterSeparator("|")
  table.SetAutoMergeCells(false)
  fmt.Println("Polling light states. This may take some time...")
  lightList, _ := ListLights(vc) // Add error reporting
  table.AppendBulk(lightList)
  table.Render()
  return nil
}

func GetLight(vc *client.VeshClient, rack_id, board_id string) ([][]string, error) {
  scanResponse, scanerr := utils.UtilScanOnly() // Add error reporting
  rackidint, _ := strconv.Atoi(rack_id)
  boardidint, _ := strconv.Atoi(board_id)
  devicePtr := reflect.ValueOf(&scanResponse.Racks[utils.RackIDtoElem(rackidint)].Boards[utils.BoardIDtoElem(boardidint)].Devices)
  deviceValuePtr := devicePtr.Elem()
  fulltable := make([][]string, 0)
  totalruns := 0
  for i := 0; i < deviceValuePtr.Len(); i++ {
    if scanResponse.Racks[utils.RackIDtoElem(rackidint)].Boards[utils.BoardIDtoElem(boardidint)].Devices[i].DeviceType == lightsdevicetype {
      device_id := scanResponse.Racks[utils.RackIDtoElem(rackidint)].Boards[utils.BoardIDtoElem(boardidint)].Devices[i].DeviceID
      responseData := &lightsResponse{}
      resp, err := vc.Sling.New().Path(lightspath).Path(rack_id + "/").Path(board_id + "/").Get(device_id).ReceiveSuccess(responseData) // Add error reporting
      if resp.StatusCode != 200 { // This is not what I meant by "error reporting"
        return nil, err
      }
      tablerow := make([]string, 0)
      tablerow = append(tablerow, rack_id, board_id, device_id)
      tablerow = append(tablerow, responseData.State)
      fulltable = append(fulltable, nil)
      fulltable[totalruns] = make([]string, 0)
      fulltable[totalruns] = append(fulltable[totalruns], tablerow...)
      totalruns ++
    }
  }
  return fulltable, scanerr
}

func PrintGetLight(vc *client.VeshClient, rack_id, board_id string) error {
  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"Rack", "Board", "Device", "LED State"})
  table.SetBorder(false)
  table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
  table.SetCenterSeparator("|")
  table.SetAutoMergeCells(true)
  lightStatus, _ := GetLight(vc, rack_id, board_id) // Add error reporting
  table.AppendBulk(lightStatus)
  table.Render()
  return nil
}

func SetLight(vc *client.VeshClient, rack_id, board_id, light_status string) (string, error) {
  responseData := &lightsResponse{}
  resp, err := vc.Sling.New().Path(lightspath).Path(rack_id + "/").Path(board_id + "/").Path(device_id + "/").Get(light_status).ReceiveSuccess(responseData) // Add error reporting
  if resp.StatusCode != 200 { // This is not what I meant by "error reporting"
    return "", err
  }
  return responseData.State, err
}

func PrintSetLight(vc *client.VeshClient, rack_id int, board_id int, light_input, light_command string) error {
  switch light_command {
  case "state":
    light_action := fmt.Sprintf("%s", "status")
    status, err := SetLight(vc, strconv.Itoa(rack_id), strconv.Itoa(board_id), light_action)
    fmt.Println(status)
    return err
  case "color":
    light_action := fmt.Sprintf("status/%s/%s", light_command, light_input) // Might need this to be a nonstring input
    status, err := SetLight(vc, strconv.Itoa(rack_id), strconv.Itoa(board_id), light_action)
    fmt.Println(status)
    return err
  case "blink":
    light_action := fmt.Sprintf("status/%s/%s", "blink_state", light_input)
    status, err := SetLight(vc, strconv.Itoa(rack_id), strconv.Itoa(board_id), light_action)
    fmt.Println(status)
    return err
  }
  return nil // Add the correct error response
}
