package transformers

import (
	"sort"

	"strings"

	"github.com/vapor-ware/synse-cli/pkg/transformers/sorters"
)

// Sortable implements the Sort interface in order to sort the
// Transformer Items.
type Sortable struct {
	// The slice of items to sort. When called via Sort from the
	// Transformer, it will be the Items of the Transformer.
	toSort []interface{}

	// The LessFuncs to sort by. This is populated via the `OrderBy`
	// function.
	less []sorters.LessFunc

	// A map containing all possible LessFunc sorters for the Sortable.
	sorterSet map[string]sorters.LessFunc
}

// sort sorts the device slice according to the less functions. This is
// not exported, as it is intended to only be called by the Transformer
// that it extends.
func (s *Sortable) sort() {
	sort.Sort(s)
}

// OrderBy takes a string that specifies the key(s) to sort by. If
// multiple keys are used, they are separated by commas, e.g. "key1,key2".
// The keys are used in the sortOrder map to get the corresponding sort
// function. If a given key is not in the map, it is ignored. The order by
// which the keys are provided determines the order in which the LessFuncs
// are executed when sorting.
func (s *Sortable) OrderBy(by string) {
	keys := strings.Split(by, ",")
	for _, key := range keys {
		fn := s.sorterSet[key]
		if fn != nil {
			s.less = append(s.less, fn)
		}
	}
}

// Len is part of sort.Interface.
func (s *Sortable) Len() int {
	return len(s.toSort)
}

// Swap is part of sort.Interface.
func (s *Sortable) Swap(i, j int) {
	s.toSort[i], s.toSort[j] = s.toSort[j], s.toSort[i]
}

// Less is part of sort.Interface. Here, it iterates over the configured
// LessFunc(s) until it finds a comparison between two items. If no difference
// is found, the results of the last LessFunc that is called are used as the
// results of the Less comparison.
func (s *Sortable) Less(i, j int) bool {
	p, q := s.toSort[i], s.toSort[j]

	var k int
	for k = 0; k < len(s.less)-1; k++ {
		less := s.less[k]
		switch {
		case less(p, q):
			// p is less than q
			return true
		case less(q, p):
			// q is less than p
			return false
		}
		// Otherwise, p == q, in which case we want to
		// continue on to the next comparison if there is one.
	}
	// If we make it here, all comparisons were 'equal', so we will
	// just arbitrarily return the result of the final comparison.
	return s.less[k](p, q)
}
