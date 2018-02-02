package scheme

type RackInfo struct {
	Rack   string   `json:"rack"`
	Boards []string `json:"boards"`
}

type BoardInfo struct {
	Board    string            `json:"board"`
	Location map[string]string `json:"location"`
	Devices  []string          `json:"devices"`
}

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

type DeviceOutput struct {
	Type      string      `json:"type"`
	DataType  string      `json:"data_type"`
	Precision int         `json:"precision"`
	Unit      OutputUnit  `json:"unit"`
	Range     OutputRange `json:"range"`
}

type OutputUnit struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type OutputRange struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}
