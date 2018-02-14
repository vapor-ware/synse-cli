package scheme

// Transaction is the scheme for the Synse Server "transaction" endpoint response.
type Transaction struct {
	ID      string       `json:"id"`
	Context WriteContext `json:"context"`
	State   string       `json:"state"`
	Status  string       `json:"status"`
	Created string       `json:"created"`
	Updated string       `json:"updated"`
	Message string       `json:"message"`
}
