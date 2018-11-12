package scheme

// DevicesOutput is the scheme for `plugin devices` pretty output.
type DevicesOutput struct {
	ID     string `json:"id"`
	Kind   string `json:"kind"`
	Plugin string `json:"plugin"`
	Info   string `json:"info"`
	Rack   string `json:"rack"`
	Board  string `json:"board"`
}
