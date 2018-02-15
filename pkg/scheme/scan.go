package scheme

import "strings"

// Scan is the scheme for the Synse Server "scan" endpoint response.
type Scan struct {
	Racks []rack `json:"racks"`
}

// rack describes a rack entry in the scan result.
type rack struct {
	ID     string  `json:"id"`
	Boards []board `json:"boards"`
}

// board describes a board entry in the scan result.
type board struct {
	ID      string   `json:"id"`
	Devices []device `json:"devices"`
}

// device describes a device entry in the scan result.
type device struct {
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
