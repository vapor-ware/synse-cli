package commands
import (
  "os"
  "fmt"
  "reflect"
  "strconv"

  "github.com/vapor-ware/vesh/client"
  "github.com/vapor-ware/vesh/utils"

  "github.com/olekukonko/tablewriter"
  "github.com/sethgrid/multibar"
)

const temperaturepath = "temperature/"
const temperaturedevicetype = "temperature"

type temperatureResponse struct {
  Health string `json:"health"`
  States []string `json:"states"`
  TemperatureC float64 `json:"temperature_c"`
}

func ListTemp(vc *client.VeshClient) ([][]string, error) {
  scanResponse, _ := utils.UtilScanOnly() // Add error reporting
  scanResponsePtr := reflect.ValueOf(&scanResponse.Racks)
  scanResponseValuePtr := scanResponsePtr.Elem()
  fulltable := make([][]string, 0)
  totalruns := 0
  totaltouched := 0
  progressBar, _ := multibar.New()
  go progressBar.Listen()
  polling := progressBar.MakeBar(utils.TotalElemsNum(), "Polling temperatures")
  for i := 0; i < scanResponseValuePtr.Len(); i++ {
    boardsPtr := reflect.ValueOf(&scanResponse.Racks[i].Boards)
    boardsValuePtr := boardsPtr.Elem()
    for j := 0; j < boardsValuePtr.Len(); j++ {
      board_id := scanResponse.Racks[i].Boards[j].BoardID
      devicePtr := reflect.ValueOf(&scanResponse.Racks[i].Boards[j].Devices)
      devicesValuePtr := devicePtr.Elem()
      for k := 0; k < devicesValuePtr.Len(); k++ {
        deviceTypePtr := reflect.ValueOf(&scanResponse.Racks[i].Boards[j].Devices[k].DeviceType)
        deviceTypeValuePtr := deviceTypePtr.Elem()
        totaltouched ++
        if deviceTypeValuePtr.String() == temperaturedevicetype { // This may need to be expanded to other types
          tablerow := make([]string, 0)
          rack_id := scanResponse.Racks[i].RackID // Maybe move this up to the "rack" loop
          device_id := scanResponse.Racks[i].Boards[j].Devices[k].DeviceID
          tablerow = append(tablerow, rack_id)
          tablerow = append(tablerow, board_id)
          tablerow = append(tablerow, device_id)
          polling(totaltouched)
          responseData := &temperatureResponse{}
          resp, err := vc.Sling.New().Path("read/").Path(temperaturepath).Path(rack_id + "/").Path(board_id + "/").Get(device_id).ReceiveSuccess(responseData) // Add error reporting
          if resp.StatusCode != 200 { // This is not what I meant by "error reporting"
            fmt.Println(vc)
            fmt.Println(resp)
            return nil, err
          }
          tablerow = append(tablerow, strconv.FormatFloat(responseData.TemperatureC, 'G', -1, 64))
          fulltable = append(fulltable, nil)
          fulltable[totalruns] = make([]string, 0)
          fulltable[totalruns] = append(fulltable[totalruns], tablerow...)
          polling(totalruns)
          totalruns ++
        }
      }
    }
  }
  return fulltable, nil
}

func PrintListTemp(vc *client.VeshClient) error {
  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"Rack", "Board", "Device", "Temperature in C"})
  table.SetBorder(false)
  table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
  table.SetCenterSeparator("|")
  table.SetAlignment(tablewriter.ALIGN_CENTER)
  table.SetAutoMergeCells(false)
  fmt.Println("Polling temperatures. This may take some time...")
  tempList, _ := ListTemp(vc) // Add error reporting
  //goterm.Flush() // Flush the current terminal buffer so we don't get weird print errors. I'm not satisfied with this, we shouldn't need it.
  table.AppendBulk(tempList)
  table.Render()
  return nil
}

func GetTemp(vc *client.VeshClient, rack_id, board_id string) ([][]string, error) {
  scanResponse, scanerr := utils.UtilScanOnly() // Add error reporting
  rackidint, _ := strconv.Atoi(rack_id)
  boardidint, _ := strconv.Atoi(board_id)
  devicePtr := reflect.ValueOf(&scanResponse.Racks[utils.RackIDtoElem(rackidint)].Boards[utils.BoardIDtoElem(boardidint)].Devices)
  deviceValuePtr := devicePtr.Elem()
  fulltable := make([][]string, 0)
  totalruns := 0
  for i := 0; i < deviceValuePtr.Len(); i++ {
    if scanResponse.Racks[utils.RackIDtoElem(rackidint)].Boards[utils.BoardIDtoElem(boardidint)].Devices[i].DeviceType == temperaturedevicetype {
      device_id := scanResponse.Racks[utils.RackIDtoElem(rackidint)].Boards[utils.BoardIDtoElem(boardidint)].Devices[i].DeviceID
      responseData := &temperatureResponse{}
      resp, err := vc.Sling.New().Path("read/").Path(temperaturepath).Path(rack_id + "/").Path(board_id + "/").Get(device_id).ReceiveSuccess(responseData) // Add error reporting
      if resp.StatusCode != 200 { // This is not what I meant by "error reporting"
        return nil, err
      }
      tablerow := make([]string, 0)
      tablerow = append(tablerow, rack_id, board_id, device_id, responseData.Health, strconv.FormatFloat(responseData.TemperatureC, 'G', -1, 64))
      tablerow = append(tablerow, responseData.States...)
      fulltable = append(fulltable, nil)
      fulltable[totalruns] = make([]string, 0)
      fulltable[totalruns] = append(fulltable[totalruns], tablerow...)
      totalruns ++
    }
  }
  return fulltable, scanerr
}

func PrintGetTemp(vc *client.VeshClient, rack_id, board_id string) error {
  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"Rack", "Board", "Device", "Health", "Temperature in C", "States"})
  table.SetBorder(false)
  table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
  table.SetCenterSeparator("|")
  table.SetAutoMergeCells(true)
  tempStatus, _ := GetTemp(vc, rack_id, board_id) // Add error reporting
  table.AppendBulk(tempStatus)
  table.Render()
  return nil
}
