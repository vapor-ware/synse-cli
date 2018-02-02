package scheme

// Read
type Read struct {
	Type string          `json:"type"`
	Data map[string]Data `json:"data"`
}

// Data
type Data struct {
	Value     interface{} `json:"value"`
	Timestamp string      `json:"timestamp"`
	Unit      Unit        `json:"unit"`
}

// Unit
type Unit struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
