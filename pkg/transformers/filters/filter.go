package filters

import (
	"fmt"
	"strings"
)

// FilterFunc is the type for a function that provides filtering constraints.
type FilterFunc func(i interface{}, v string) bool

// Filter represents a filter key/value pair specified from the CLI via the
// --filter flag. The value of the flag should follow the form "key=value", which
// is parsed into the corresponding fields of this struct.
type Filter struct {
	Key   string
	Value string
}

// NewFilter creates a new instance of a Filter given the value taken from the
// CLI --filter flag.
func NewFilter(filterString string) (*Filter, error) {
	f := strings.Split(filterString, "=")
	if len(f) != 2 {
		return nil, fmt.Errorf("filter must be in the form 'key=value'")
	}
	return &Filter{
		Key:   strings.ToLower(f[0]),
		Value: f[1],
	}, nil
}
