package scheme

// ReadCached is the scheme for the Synse Server "readcached" endpoint
// response. It is similar to ReadData, with the addtion of Kind and Location.
type ReadCached struct {
	// ReadData values.
	ReadData

	// Augmented vales.
	Kind     string         `json:"kind"`
	Location DeviceLocation `json:"location"`
}

// ReadCachedParams is the scheme for the Synse Server "readcached" endpoint
// parameters.
type ReadCachedParams struct {
	Start string `url:"start,omitempty"`
	End   string `url:"end,omitempty"`
}

// ReadCachedOutput defines the scheme for the data output by a "readcached" command.
type ReadCachedOutput struct {
	Location  DeviceLocation `json:"location" yaml:"location"`
	Info      string         `json:"info" yaml:"info"`
	Type      string         `json:"type" yaml:"type"`
	Value     string         `json:"value" yaml:"value"`
	Unit      string         `json:"unit" yaml:"unit"`
	Timestamp string         `json:"timestamp" yaml:"timestamp"`
}
