package scheme

// Read is the scheme for the Synse Server "read" endpoint response.
type Read struct {
	Type string              `json:"type"`
	Data map[string]ReadData `json:"data"`
}

// ReadData describes reading data.
type ReadData struct {
	Value     interface{} `json:"value"`
	Timestamp string      `json:"timestamp"`
	Unit      OutputUnit  `json:"unit"`
}
