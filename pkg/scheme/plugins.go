package scheme

// Plugin is the scheme for the Synse Server "plugin" endpoint response.
type Plugin struct {
	Name        string                 `json:"name" yaml:"name"`
	Tag         string                 `json:"tag" yaml:"tag"`
	Description string                 `json:"description" yaml:"description"`
	Maintainer  string                 `json:"maintainer" yaml:"maintainer"`
	VCS         string                 `json:"vcs" yaml:"vcs"`
	Network     map[string]interface{} `json:"network" yaml:"network"`
	Health      map[string]interface{} `json:"health" yaml:"health"`
	Version     map[string]interface{} `json:"version" yaml:"version"`
}
