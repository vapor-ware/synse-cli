// Package transformers defines a Transformer for CLI output. A Transformer
// is anything that can take a set of data and apply some constraints or
// processing to the data. Currently, there are two transformers: sorting
// and filtering. Each CLI command can have its own Transformer, but that
// Transformer is not required to define the constraints around transformations,
// so capabilities are command-specific.
package transformers

import (
	"fmt"
	"reflect"

	"github.com/vapor-ware/synse-cli/pkg/transformers/filters"
	"github.com/vapor-ware/synse-cli/pkg/transformers/sorters"
)

// Transformer supports transformations on a slice of items, such as
// sorting and filtering.
type Transformer struct {
	Sortable
	Filterable

	Items []interface{}
}

// Sort applies the LessFuncs configured for the Sortable over the
// Transformer's Items.
func (t *Transformer) Sort() {
	if len(t.less) > 0 {
		t.toSort = t.Items
		t.sort()
	}
}

// Filter applies the FilterFuncs configured for the Filterable over
// the Transformer's Items.
func (t *Transformer) Filter() {
	if len(t.filters) > 0 {
		t.toFilter = t.Items
		filtered := t.filter()
		t.Items = filtered
	}
}

// Apply will apply all of the transformations which are set with the
// Transformer. Any transformation which was not configured is skipped.
// Apply will execute transformations in the following order: sort, filter.
func (t *Transformer) Apply() {
	t.Sort()
	t.Filter()
}

// NewTransformer creates a new instance of a Transformer. It checks that the
// given items is actually a slice of items, and if so uses that as the
// Transformer Items. In general, it is recommended to use a command-specific
// constructor rather than this one.
func NewTransformer(items interface{}) (*Transformer, error) {
	var itemSlice []interface{}

	k := reflect.TypeOf(items).Kind()
	switch k {
	case reflect.Slice:
		v := reflect.ValueOf(items)
		for i := 0; i < v.Len(); i++ {
			itemSlice = append(itemSlice, v.Index(i).Interface())
		}
	default:
		return nil, fmt.Errorf("'items' should be a slice but is %v", k)
	}

	return &Transformer{
		Items: itemSlice,
	}, nil
}

// NewScanTransformer creates a new Transformer instance configured for the
// Synse Server scan command.
func NewScanTransformer(items interface{}) (*Transformer, error) {
	t, err := NewTransformer(items)
	if err != nil {
		return nil, err
	}
	t.sorterSet = sorters.ScanLess
	t.filterSet = filters.ScanFilters

	return t, nil
}
