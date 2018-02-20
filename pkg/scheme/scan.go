package scheme

import (
	"fmt"
	"sort"
	"strings"

	"github.com/urfave/cli"
)

// ScanResponse is the scheme for the Synse Server "scan" endpoint response.
// It is intended to be the structure that a scan gets marshaled into, but
// should be converted into a slice of `Scan` structs for all internal use.
type ScanResponse struct {
	Racks []struct {
		ID     string `json:"id" yaml:"id"`
		Boards []struct {
			ID      string `json:"id" yaml:"id"`
			Devices []struct {
				ID   string `json:"id" yaml:"id"`
				Type string `json:"type" yaml:"type"`
				Info string `json:"info" yaml:"info"`
			} `json:"devices" yaml:"devices"`
		} `json:"boards" yaml:"boards"`
	} `json:"racks" yaml:"racks"`
}

// ToScanDevices converts the ScanResponse result scheme to a slice of ScanDevice
// representations of the scan results, which makes it easier to sort, filter,
// format, and generally work with the scan results at a device-level.
func (s *ScanResponse) ToScanDevices() []*ScanDevice {
	var devices []*ScanDevice
	for _, rack := range s.Racks {
		for _, board := range rack.Boards {
			for _, device := range board.Devices {
				devices = append(devices, &ScanDevice{
					Rack:   rack.ID,
					Board:  board.ID,
					Device: device.ID,
					Info:   device.Info,
					Type:   device.Type,
				})
			}
		}
	}
	return devices
}

func rackOrder(p1, p2 interface{}) bool  { return p1.(*ScanDevice).Rack < p2.(*ScanDevice).Rack }
func boardOrder(p1, p2 interface{}) bool { return p1.(*ScanDevice).Board < p2.(*ScanDevice).Board }
func deviceOrder(p1, p2 interface{}) bool {
	return p1.(*ScanDevice).Device < p2.(*ScanDevice).Device
}
func typeOrder(p1, p2 interface{}) bool {
	return p1.(*ScanDevice).Type < p2.(*ScanDevice).Type
}

var scanSorters = map[string]LessFunc{
	"rack":   rackOrder,
	"board":  boardOrder,
	"device": deviceOrder,
	"type":   typeOrder,
}

func typeFilter(i interface{}, v string) bool {
	return i.(*ScanDevice).Type == v
}

func rackFilter(i interface{}, v string) bool {
	return i.(*ScanDevice).Rack == v
}

func boardFilter(i interface{}, v string) bool {
	return i.(*ScanDevice).Board == v
}

var scanFilters = map[string]FilterFunc{
	"type":  typeFilter,
	"rack":  rackFilter,
	"board": boardFilter,
}

func (s *ScanResponse) ToScanTransformer() *Transformer {
	devices := s.ToScanDevices()

	items := make([]interface{}, len(devices))
	for i, device := range devices {
		items[i] = device
	}
	t := &Transformer{
		Items: items,
	}
	t.sortOrder = scanSorters
	t.filterSet = scanFilters
	return t
}

// ScanDevice represents a single device from the "scan" output. A slice of
// ScanDevice is a flattened ScanResponse. This scheme is used internally to
// more easily sort, filter, and format scan results.
type ScanDevice struct {
	Rack   string `json:"rack" yaml:"rack"`
	Board  string `json:"board" yaml:"board"`
	Device string `json:"device" yaml:"device"`
	Info   string `json:"info" yaml:"info"`
	Type   string `json:"type" yaml:"type"`
}

// Transformer supports transformations on a slice of items, such as
// sorting and filtering.
type Transformer struct {
	Sortable
	Filterable

	Items []interface{}
}

func (t *Transformer) Sort() {
	t.toSort = t.Items
	t.sort()
}

func (t *Transformer) Filter() {
	if len(t.filters) > 0 {
		t.toFilter = t.Items
		filtered := t.filter()
		t.Items = filtered
	}
}

type FilterFunc func(i interface{}, v string) bool

type Filter struct {
	filterKey string
	filterVal string
}

func NewFilter(filterString string) (*Filter, error) {
	f := strings.Split(filterString, "=")
	if len(f) != 2 {
		return nil, fmt.Errorf("filter must be in the form 'key=value'")
	}
	return &Filter{
		filterKey: strings.ToLower(f[0]),
		filterVal: f[1],
	}, nil
}

type Filterable struct {
	// The slice of items to filter. When called via Filter from the
	// Transformer, it will be the Items of the transformer.
	toFilter []interface{}

	// The filters to use when filtering. These should be generated from
	// the values provided by the --filter flag.
	filters []*Filter

	// A map containing the possible FilterFunc filters for the Filterable.
	// Here, the key corresponds to the filter key passed in via the filter
	// flag for a command.
	filterSet map[string]FilterFunc
}

func (f *Filterable) filter() []interface{} {
	filtered := make([]interface{}, 0)
	for _, item := range f.toFilter {
		for _, filter := range f.filters {
			fn := f.filterSet[filter.filterKey]
			if fn != nil {
				if fn(item, filter.filterVal) {
					filtered = append(filtered, item)
				}
			}
		}
	}
	return filtered
}

func (f *Filterable) FiltersFromContext(c *cli.Context) error {
	filterString := c.String("filter")
	if filterString != "" {
		filter, err := NewFilter(filterString)
		if err != nil {
			return err
		}
		f.filters = append(f.filters, filter)
	}
	return nil
}

// LessFunc defines a function that compares two sortable items.
type LessFunc func(p1, p2 interface{}) bool

// Sortable implements the Sort interface, sorting the ScanDevice
// instances within.
type Sortable struct {
	// The slice of items to sort. When called via Sort from the
	// Transformer, it will be the Items of the Transformer.
	toSort []interface{}

	// The LessFuncs to sort by.
	less []LessFunc

	// A map containing all possible LessFunc sorters for the Sortable.
	sortOrder map[string]LessFunc
}

// Sort sorts the device slice according to the less functions. This is
// not exported, as it is intended to only be called by the Transformer
// that it extends.
func (s *Sortable) sort() {
	sort.Sort(s)
}

// OrderBy takes an array of strings that are keys into the
// Sortable sortOrder map. If a given key is not in the map, it
// is ignored. The order by which key strings are provided determines
// the order in which the LessFuncs are called when sorting.
func (s *Sortable) OrderBy(by ...string) {
	for _, key := range by {
		fn := s.sortOrder[key]
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
