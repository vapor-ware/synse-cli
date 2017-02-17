package commands
import (
  "os"
  //"fmt"
  "reflect"

  "github.com/vapor-ware/vesh/client"
  "github.com/vapor-ware/vesh/utils"

  "github.com/olekukonko/tablewriter"
)

const hostnamepath = "boot_target/"

type hostname struct {
  Hostname []string `json:"hostnames"`
  IPAddress []string `json:"ip_addresses"`
}

func ListHostnames(vc *client.VeshClient) error {
  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"Board ID", "Hostnames", "IP Addesses"})
  table.SetBorder(false)
  status := &hostname{}
  scanresult, scanerr := Scan(vc)
  scanPtr:= reflect.ValueOf(&scanresult.Racks)
  racksPtr := scanPtr.Elem()
  racks := make([]string, 0)
  boards := make([]string, 0)
  for i := 0; i < racksPtr.Len(); i++ {
    racks := append(racks, racksPtr.Field(i).String())
  }
  for j := range racks {
    for k := range boards {
      boardsPtr := reflect.ValueOf(&scanresult.Racks[j].Boards[k].Hostnames)
      boardsValuePtr := boardsPtr.Elem()
      hostnames := make([]string, 0)
      hostnames = append(hostnames, boardsValuePtr.Field(j).String())
      ipaddresstostring, err := utils.StructToSingleSlice(&scanresult.Racks[j].Boards[k].IPAddresses)
      //ipaddresses := make([]string, 0)
      //ipaddresses = append(ipaddresses, ipaddresstostring)
      for l := range ipaddresstostring {
        fulltablerow := append(hostnames, ipaddresstostring[l]...)
      }
      table.Append(fulltablerow)
    }
  }
  table.Render()
  return nil
}
