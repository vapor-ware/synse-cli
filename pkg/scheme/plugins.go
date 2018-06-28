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
