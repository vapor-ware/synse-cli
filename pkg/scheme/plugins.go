package scheme

// Plugin is the scheme for the Synse Server "plugin" endpoint response.
type Plugin struct {
	Name    string `json:"name" yaml:"name"`
	Network string `json:"network" yaml:"network"`
	Address string `json:"address" yaml:"address"`
}
