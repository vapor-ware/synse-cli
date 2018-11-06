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

// ReadCached is the scheme for the Synse Server "readcached" endpoint
// response. It is similar to ReadData, with the addtion of Kind and Location.
type ReadCached struct {
	// ReadData values.
	ReadData

	// Augmented vales.
	Kind     string         `json:"kind"`
	Location DeviceLocation `json:"location"`
}

// DeviceLocation describes the location of a device by providing the
// rack, board, and device ID which is used as routing info to that device.
type DeviceLocation struct {
	Rack   string `json:"rack"`
	Board  string `json:"board"`
	Device string `json:"device"`
}

// ReadCachedParams is the scheme for the Synse Server "readcached" endpoint
// parameters.
type ReadCachedParams struct {
	Start string `url:"start,omitempty"`
	End   string `url:"end,omitempty"`
}

// ReadCachedOutput defines the scheme for the data output by a "readcached" command.
type ReadCachedOutput struct {
	// FIXME: Embedded DeviceLocation doesn't work here for some reasons.
	// Have to define Rack, Board, Device again.
	Rack   string
	Board  string
	Device string

	Info      string
	Type      string
	Value     string
	Unit      string
	Timestamp string
}
