package scheme

// RackInfo is the scheme for the Synse Server "info" endpoint response
// for a request at the rack level.
type RackInfo struct {
	Rack   string   `json:"rack"`
	Boards []string `json:"boards"`
}

// BoardInfo is the scheme for the Synse Server "info" endpoint response
// for a request at the board level.
type BoardInfo struct {
	Board    string            `json:"board"`
	Location map[string]string `json:"location"`
	Devices  []string          `json:"devices"`
}

// DeviceInfo is the scheme for the Synse Server "info" endpoint response
// for a request at the device level.
type DeviceInfo struct {
	Timestamp string            `json:"timestamp"`
	UID       string            `json:"uid"`
	Kind      string            `json:"kind"`
	Metadata  map[string]string `json:"metadata"`
	Plugin    string            `json:"plugin"`
	Info      string            `json:"info"`
	Location  map[string]string `json:"location"`
	Output    []DeviceOutput    `json:"output"`
}

// DeviceOutput is the scheme for the output meta-info belonging to a device.
type DeviceOutput struct {
	Name          string     `json:"name"`
	Type          string     `json:"type"`
	Precision     int        `json:"precision,omitempty"`
	ScalingFactor float64    `json:"scaling_factor,omitempty" yaml:"scaling_factor,omitempty"`
	Unit          OutputUnit `json:"unit,omitempty"`
}

// OutputUnit describes the unit for a given device output.
type OutputUnit struct {
	Name   string `json:"name,omitempty"`
	Symbol string `json:"symbol,omitempty"`
}
