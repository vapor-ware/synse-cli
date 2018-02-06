package scheme

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
