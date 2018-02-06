package scheme

// WriteTransaction is the scheme for the Synse Server "write" endpoint response.
type WriteTransaction struct {
	Context     writeContext `json:"context"`
	Transaction string       `json:"transaction"`
}

// writeContext describes the context returned with a write transaction.
type writeContext struct {
	Action string   `json:"action"`
	Raw    []string `json:"raw"`
}
