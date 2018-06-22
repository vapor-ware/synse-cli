package scheme

// DevicesOutput dines the scheme for the data output by a "plugin devices" command.
type DevicesOutput struct {
	ID     string
	Kind   string
	Plugin string
	Info   string
	Rack   string
	Board  string
}
