package utils
import (
  "net/http"

  "github.com/vapor-ware/vesh/client"
)

// I DON'T LIKE THIS AT ALL

const Scanpath = "scan"

type scanResponse struct {
  Racks []struct {
    Boards []struct {
      BoardID string `json:"board_id"`
      Hostnames []string `json:"hostnames"`
      IPAddresses []string `json:"ip_addresses"`
      Devices []struct {
        DeviceID string `json:"device_id"`
        DeviceInfo string `json:"device_info"`
        DeviceType string `json:"device_type"`
      } `json:"devices"`
    } `json:"boards"`
    RackID string `json:"rack_id"`
  } `json:"racks"`
}


func UtilScanOnly() (*scanResponse, error) {
  vc := client.New()
  status := &scanResponse{}
  resp, err := vc.Sling.New().Get(Scanpath).ReceiveSuccess(status)
  if err != nil {
    return nil, err
  }
  if resp.StatusCode != http.StatusOK {
    return status, err
  }
  return status, nil
}
