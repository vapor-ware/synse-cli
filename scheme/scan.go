package scheme

// Scan
type Scan struct {
	Racks []Rack `json:"racks"`
}

// Rack
type Rack struct {
	Id     string  `json:"id"`
	Boards []Board `json:"boards"`
}

// Board
type Board struct {
	Id      string   `json:"id"`
	Devices []Device `json:"devices"`
}

// Device
type Device struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Info string `json:"info"`
}
