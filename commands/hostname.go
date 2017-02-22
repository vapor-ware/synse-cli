package commands
import (
  "os"
  "fmt"
  "reflect"

  "github.com/vapor-ware/vesh/client"
  //"github.com/vapor-ware/vesh/utils"

  "github.com/olekukonko/tablewriter"
)

const hostnamepath = "host_info/"

type hostnameResponse struct {
  Hostname []string `json:"hostnames"`
  IPAddress []string `json:"ip_addresses"`
}

func ListHostnames(vc *client.VeshClient) error {
  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"Hostnames", "IP Addesses", "Board ID",})
  table.SetBorder(false)
  //table.SetFooter([]string{"", "", "I hate Thomas"})
  //status := &hostname{}
  scanresult, scanerr := ScanOnly(vc)
  racksPtr := reflect.ValueOf(&scanresult.Racks)
  racksValuePtr := racksPtr.Elem()
  racks := make([]string, 0)
  for i := 0; i < racksValuePtr.Len(); i++ {
    rackIDPtr := reflect.ValueOf(&scanresult.Racks[i].RackID)
    rackIDValuePtr := rackIDPtr.Elem()
    racks = append(racks, rackIDValuePtr.String())
    boardsPtr := reflect.ValueOf(&scanresult.Racks[i].Boards)
    boardsValuePtr := boardsPtr.Elem()
    for n := 0; n < boardsValuePtr.Len(); n++ {
      boardsIDPtr := reflect.ValueOf(&scanresult.Racks[i].Boards[n].BoardID)
      boardsIDValuePtr := boardsIDPtr.Elem()
      boards := make([]string, 0)
      boards = append(boards, boardsIDValuePtr.String())
      hostnamePtr := reflect.ValueOf(&scanresult.Racks[i].Boards[n].Hostnames)
      hostnameValuePtr := hostnamePtr.Elem()
      ipaddressPtr := reflect.ValueOf(&scanresult.Racks[i].Boards[n].IPAddresses)
      ipaddressValuePtr := ipaddressPtr.Elem()
      hostnames := make([]string, 0)
      hostnames = append(hostnames, hostnameValuePtr.String())
      ipaddresses := make([]string, 0)
      ipaddresses = append(ipaddresses, ipaddressValuePtr.String())
      tablerow := make([]string, 0)
      for l := range scanresult.Racks[i].Boards[n].Hostnames {
        tablerow = append(tablerow, scanresult.Racks[i].Boards[n].Hostnames[l])
      }
      for m := range scanresult.Racks[i].Boards[n].IPAddresses {
        tablerow = append(tablerow, scanresult.Racks[i].Boards[n].IPAddresses[m])
      }
      tablerow = append(tablerow, boards[0])
      table.Append(tablerow)
    }
  }
  //fmt.Println(len(racks))
  table.Render()
  return scanerr //fix this return
}

func GetHostname(vc *client.VeshClient, rack_id, board_id, device_id string) ([]string, error) {
  responseData := &hostnameResponse{}
  _, err := vc.Sling.New().Path(hostnamepath).Path(rack_id + "/").Path(board_id + "/").Get(device_id).ReceiveSuccess(responseData)
  tableline := make([]string, 0)
  tableline = append(tableline, responseData.Hostname...)
  tableline = append(tableline, responseData.IPAddress...)
  if err != nil {
    return nil, err
  }
  return tableline, err
}

func PrintGetHostname(vc *client.VeshClient, rack_id, board_id, device_id string, raw bool) error {
  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"Hostnames", "IP Addesses"})
  table.SetBorder(false)
  tablerow, err := GetHostname(vc, rack_id, board_id, device_id)
  if err != nil {
    return err
  }
  if raw == true {
    fmt.Println(tablerow)
    return nil
  }
  table.Append(tablerow)
  table.Render()
  return nil

}
