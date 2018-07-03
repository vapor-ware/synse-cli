package scheme

// DevicesOutput is the scheme for `plugin devices` pretty output.
type DevicesOutput struct {
	ID     string
	Kind   string
	Plugin string
	Info   string
	Rack   string
	Board  string
}
