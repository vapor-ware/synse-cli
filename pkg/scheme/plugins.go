package scheme

// Plugin is the scheme for the Synse Server "plugin" endpoint response.
type Plugin struct {
	Name        string      `json:"name" yaml:"name"`
	Tag         string      `json:"tag" yaml:"tag"`
	Description string      `json:"description" yaml:"description"`
	Maintainer  string      `json:"maintainer" yaml:"maintainer"`
	VCS         string      `json:"vcs" yaml:"vcs"`
	Network     NetworkData `json:"network" yaml:"network"`
	Health      HealthData  `json:"health" yaml:"health"`
	Version     VersionData `json:"version" yaml:"version"`
}

// NetworkData describes network data.
type NetworkData struct {
	Protocol string `json:"protocol" yaml:"protocol"`
	Address  string `json:"address" yaml:"address"`
}

// HealthData describe health data.
type HealthData struct {
	Status    string      `json:"status" yaml:"status"`
	Message   string      `json:"message" yaml:"message"`
	Timestamp string      `json:"timestamp" yaml:"timestamp"`
	Checks    []CheckData `json:"checks" yaml:"checks"`
}

// VersionData describes version data.
type VersionData struct {
	Version    string `json:"plugin_version" yaml:"plugin_version"`
	SDKVersion string `json:"sdk_version" yaml:"sdk_version"`
	BuildDate  string `json:"build_date" yaml:"build_date"`
	GitCommit  string `json:"git_commit" yaml:"git_commit"`
	GitTag     string `json:"git_tag" yaml:"git_tag"`
	Arch       string `json:"arch" yaml:"arch"`
	OS         string `json:"os" yaml:"os"`
}

// CheckData describes checks data that is a part of health data.
type CheckData struct {
	Name      string `json:"name" yaml:"name"`
	Status    string `json:"status" yaml:"status"`
	Message   string `json:"message" yaml:"message"`
	Timestamp string `json:"timestamp" yaml:"timestamp"`
	Type      string `json:"type" yaml:"type"`
}

// ServerPluginOutput is the scheme for `server plugins` command pretty output.
type ServerPluginOutput struct {
	Tag      string `json:"tag" yaml:"tag"`
	Protocol string `json:"protocol" yaml:"protocol"`
	Address  string `json:"address" yaml:"address"`
	Version  string `json:"version" yaml:"version"`
	Status   string `json:"status" yaml:"status"`
}

// ServerPluginInfoOutput is the scheme for `server plugins info` command output.
// FIXME: This struct is the same as the Plugin above, except the Health field.
// This might create redundancy but definitely help with readability and
// consistency because it is separating out Synse Server's response and CLI's
// output.
// TODO: If that's the case, need to go through other files and model after this.
type ServerPluginInfoOutput struct {
	Name        string      `json:"name" yaml:"name"`
	Tag         string      `json:"tag" yaml:"tag"`
	Description string      `json:"description" yaml:"description"`
	Maintainer  string      `json:"maintainer" yaml:"maintainer"`
	VCS         string      `json:"vcs" yaml:"vcs"`
	Network     NetworkData `json:"network" yaml:"network"`
	Version     VersionData `json:"version" yaml:"version"`
}

// ServerPluginHealthOutput is the scheme for `server plugins health` command
// output.
type ServerPluginHealthOutput struct {
	Tag    string     `json:"tag" yaml:"tag"`
	Health HealthData `json:"health" yaml:"health"`
}
