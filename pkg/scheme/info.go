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
	Timestamp    string            `json:"timestamp"`
	UID          string            `json:"uid"`
	Type         string            `json:"type"`
	Model        string            `json:"model"`
	Manufacturer string            `json:"manufacturer"`
	Protocol     string            `json:"protocol"`
	Info         string            `json:"info"`
	Comment      string            `json:"comment"`
	Location     map[string]string `json:"location"`
	Output       []DeviceOutput    `json:"output"`
}

// DeviceOutput is the scheme for the output meta-info belonging to a device.
type DeviceOutput struct {
	Type      string      `json:"type"`
	DataType  string      `json:"data_type"`
	Precision int         `json:"precision,omitempty"`
	Unit      OutputUnit  `json:"unit,omitempty"`
	Range     OutputRange `json:"range,omitempty"`
}

// OutputUnit describes the unit for a given device output.
type OutputUnit struct {
	Name   string `json:"name,omitempty"`
	Symbol string `json:"symbol,omitempty"`
}

// OutputRange describes the min and max value range for a given device output.
type OutputRange struct {
	Min float64 `json:"min,omitempty"`
	Max float64 `json:"max,omitempty"`
}
