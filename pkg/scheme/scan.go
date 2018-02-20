package scheme

// ScanResponse is the scheme for the Synse Server "scan" endpoint response.
// It is intended to be the structure that a scan gets marshaled into, but
// should be converted into a slice of `Scan` structs for all internal use.
type ScanResponse struct {
	Racks []struct {
		ID     string `json:"id" yaml:"id"`
		Boards []struct {
			ID      string `json:"id" yaml:"id"`
			Devices []struct {
				ID   string `json:"id" yaml:"id"`
				Type string `json:"type" yaml:"type"`
				Info string `json:"info" yaml:"info"`
			} `json:"devices" yaml:"devices"`
		} `json:"boards" yaml:"boards"`
	} `json:"racks" yaml:"racks"`
}

// ToScanDevices converts the ScanResponse result scheme to a slice of ScanDevice
// representations of the scan results, which makes it easier to sort, filter,
// format, and generally work with the scan results at a device-level.
// FIXME - this should move to utils?
func (s *ScanResponse) ToScanDevices() []*ScanDevice {
	var devices []*ScanDevice
	for _, rack := range s.Racks {
		for _, board := range rack.Boards {
			for _, device := range board.Devices {
				devices = append(devices, &ScanDevice{
					Rack:   rack.ID,
					Board:  board.ID,
					Device: device.ID,
					Info:   device.Info,
					Type:   device.Type,
				})
			}
		}
	}
	return devices
}

// ScanDevice represents a single device from the "scan" output. A slice of
// ScanDevice is a flattened ScanResponse. This scheme is used internally to
// more easily sort, filter, and format scan results.
type ScanDevice struct {
	Rack   string `json:"rack" yaml:"rack"`
	Board  string `json:"board" yaml:"board"`
	Device string `json:"device" yaml:"device"`
	Info   string `json:"info" yaml:"info"`
	Type   string `json:"type" yaml:"type"`
}
