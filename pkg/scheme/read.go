package scheme

// Read is the scheme for the Synse Server "read" endpoint response.
type Read struct {
	Type string     `json:"type"`
	Data []ReadData `json:"data"`
}

// ReadData describes reading data.
type ReadData struct {
	Value     interface{} `json:"value"`
	Timestamp string      `json:"timestamp"`
	Unit      OutputUnit  `json:"unit"`
}

// ReadOutput defines the scheme for the data output by a "read" command.
type ReadOutput struct {
	Type      string `json:"type"`
	Value     string `json:"value"`
	Unit      string `json:"unit"`
	Timestamp string `json:"timestamp"`
}
