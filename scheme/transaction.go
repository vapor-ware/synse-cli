package scheme

// Transaction
type Transaction struct {
	Id      string       `json:"id"`
	Context WriteContext `json:"context"`
	State   string       `json:"state"`
	Created string       `json:"created"`
	Updated string       `json:"updated"`
	Message string       `json:"message"`
}
