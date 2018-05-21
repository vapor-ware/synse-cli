package scheme

// WriteTransaction is the scheme for the Synse Server "write" endpoint response.
type WriteTransaction struct {
	Context     WriteContext `json:"context"`
	Transaction string       `json:"transaction"`
}

// WriteContext describes the context returned with a write transaction.
type WriteContext struct {
	Action string   `json:"action"`
	Raw    []string `json:"raw"`
}

// WriteOutput defines the scheme for the data output by a "write" command.
type WriteOutput struct {
	Transaction string
	Action      string
	Raw         []string
}
