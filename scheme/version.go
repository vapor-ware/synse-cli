package scheme

// Version is the scheme for the Synse Server "version" endpoint response.
type Version struct {
	Version    string `json:"version"`
	APIVersion string `json:"api_version"`
}
