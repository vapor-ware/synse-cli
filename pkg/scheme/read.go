package scheme

// Read is the scheme for the Synse Server "read" endpoint response.
type Read struct {
	Kind string     `json:"kind" yaml:"kind"`
	Data []ReadData `json:"data" yaml:"data"`
}

// ReadData describes reading data.
type ReadData struct {
	Value     interface{} `json:"value" yaml:"value"`
	Timestamp string      `json:"timestamp" yaml:"timestamp"`
	Type      string      `json:"type" yaml:"type"`
	Info      string      `json:"info" yaml:"info"`
	Unit      OutputUnit  `json:"unit" yaml:"unit"`
}

// ReadOutput defines the scheme for the data output by a "read" command.
type ReadOutput struct {
	Info      string `json:"info" yaml:"info"`
	Type      string `json:"type" yaml:"type"`
	Value     string `json:"value" yaml:"value"`
	Unit      string `json:"unit" yaml:"unit"`
	Timestamp string `json:"timestamp" yaml:"timestamp"`
}
