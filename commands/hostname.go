package commands
import (
  "os"
  //"fmt"
  "reflect"

  "github.com/vapor-ware/vesh/client"
  //"github.com/vapor-ware/vesh/utils"

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
  //status := &hostname{}
  scanresult, scanerr := Scan(vc)
  scanPtr := reflect.ValueOf(&scanresult.Racks)
  racksPtr := scanPtr.Elem()
  racks := make([]string, 0)
  boards := make([]string, 0)
  for i := 0; i < racksPtr.Len(); i++ {
    rackIDPtr := reflect.ValueOf(&scanresult.Racks[i].RackID)
    rackIDValuePtr := rackIDPtr.Elem()
    racks = append(racks, rackIDValuePtr.String())
  }
  for j := range racks {
    for k := range boards {
      hostnamePtr := reflect.ValueOf(&scanresult.Racks[j].Boards[k].Hostnames)
      hostnameValuePtr := hostnamePtr.Elem()
      ipaddressPtr := reflect.ValueOf(&scanresult.Racks[j].Boards[k].IPAddresses)
      ipaddressValuePtr := ipaddressPtr.Elem()
      hostnames := make([]string, 0)
      hostnames = append(hostnames, hostnameValuePtr.Field(j).String())
      ipaddresses := make([]string, 0)
      ipaddresses = append(ipaddresses, ipaddressValuePtr.Field(j).String())
      tablerow := make([]string, 0)
      for l := range hostnames {
        tablerow = append(tablerow, hostnames[l])
      }
      for m := range ipaddresses {
        tablerow = append(tablerow, ipaddresses[m])
      }
      table.Append(tablerow)
    }
  }
  table.Render()
  return scanerr //fix this return
}
