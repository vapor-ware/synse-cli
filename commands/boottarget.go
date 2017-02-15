package commands
import (


  "github.com/vapor-ware/vesh/client"

  "github.com/olekukonko/tablewriter"
)

const bootpath = "boot_target/"
var bootdevicetype = "system"

type boottargetresponse struct {
  Target string `json:"target"`
  status string `json:"status"`
}

func GetCurrentBootTarget(rack_id int, board_id int) error {
  client := &boottargetresponse{}
  resp, err := vc.Sling.New().Path(bootpath).Path(rack_id + "/").Path(board_id + "/").Get(bootdevicetype).ReceiveSuccess(status)
  if err != nil {
    return err
  }
  if resp.StatusCode != http.StatusOK {
    return err
  }
  fmt.Println(resp.Target)
  return nil
}

func SetCurrentBootTarget(rack_id int, board_id int, boot_target string) error {
  client := &boottargetresponse{}
  resp, err := vc.Sling.New().Path(bootpath).Path(rack_id + "/").Path(board_id + "/").Path(bootdevicetype).Get(boot_target).ReceiveSuccess(status)
  if err != nil {
    return err
  }
  if resp.StatusCode != http.StatusOK {
    return err
  }
  fmt.Println(resp.Target, resp.status)
  return nil
}
