package scheme

// Error is the scheme for Synse Server error responses (e.g. 500, 404).
type Error struct {
	Code        int    `json:"http_code"`
	ErrorID     int    `json:"error_id"`
	Description string `json:"description"`
	Timestamp   string `json:"timestamp"`
	Context     string `json:"context"`
}
