package scheme

import "strings"

// Scan is the scheme for the Synse Server "scan" endpoint response.
type Scan struct {
	Racks []Rack `json:"racks"`
}

// ToInternalScan converts the Scan result scheme to a list of InternalScan
// representations of the scan results, which makes it easier to sort, filter,
// format, and generally work with the scan results at a device-level.
func (s *Scan) ToInternalScan() []*InternalScan {
	var devices []*InternalScan
	for _, rack := range s.Racks {
		for _, board := range rack.Boards {
			for _, device := range board.Devices {
				devices = append(devices, &InternalScan{
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

// Rack describes a rack entry in the scan result.
type Rack struct {
	ID     string  `json:"id"`
	Boards []Board `json:"boards"`
}

// Board describes a board entry in the scan result.
type Board struct {
	ID      string   `json:"id"`
	Devices []Device `json:"devices"`
}

// Device describes a device entry in the scan result.
type Device struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Info string `json:"info"`
}

// InternalScan is a representation of a Scan result that is used internally.
// FIXME (etd): not sure if this belongs in scheme. this can be used be formatting,
//   sorting, and filtering. Other commands will need something similar. Will sort this
//   out when I tackle generalizing of sorting/filtering.
type InternalScan struct {
	Rack   string
	Board  string
	Device string
	Info   string
	Type   string
}

// ID generates the ID of the device by joining the rack, board, and device.
func (device *InternalScan) ID() string {
	return strings.Join([]string{
		device.Rack,
		device.Board,
		device.Device,
	}, "-")
}
