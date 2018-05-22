package scheme

// ListHostOutput defines the scheme for the data output by a "list host" command.
type ListHostOutput struct {
	Active  bool   `json:"active"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
