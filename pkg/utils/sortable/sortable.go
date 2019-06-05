package sortable

import "github.com/vapor-ware/synse-cli/pkg/config"

// ContextRecords implements sort.Interface based on the default
// of the Type, Name, and Address fields.
type ContextRecords []config.ContextRecord

func (r ContextRecords) Len() int {
	return len(r)
}

func (r ContextRecords) Less(i, j int) bool {
	if r[i].Type < r[j].Type {
		return true
	}
	if r[i].Type > r[j].Type {
		return false
	}
	if r[i].Name < r[j].Name {
		return true
	}
	if r[i].Name > r[j].Name {
		return false
	}
	return r[i].Context.Address < r[j].Context.Address
}

func (r ContextRecords) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
