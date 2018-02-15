package scheme

// TestStatus is the scheme for the Synse Server "test" endpoint response.
type TestStatus struct {
	Status    string `json:"status" yaml:"status"`
	Timestamp string `json:"timestamp" yaml:"timestamp"`
}
