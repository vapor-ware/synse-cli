package scheme

// Read is the scheme for the Synse Server "read" endpoint response.
type Read struct {
	Type string              `json:"type"`
	Data map[string]readData `json:"data"`
}

// readData describes reading data.
type readData struct {
	Value     interface{} `json:"value"`
	Timestamp string      `json:"timestamp"`
	Unit      OutputUnit  `json:"unit"`
}
