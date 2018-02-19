package scheme

import (
	"sort"
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



func rackOrder(p1, p2 interface{}) bool { return p1.(*ScanDevice).Rack < p2.(*ScanDevice).Rack }
func boardOrder(p1, p2 interface{}) bool { return p1.(*ScanDevice).Board < p2.(*ScanDevice).Board }
func deviceOrder(p1, p2 interface{}) bool {
	return p1.(*ScanDevice).Device < p2.(*ScanDevice).Device
}
func typeOrder(p1, p2 interface{}) bool {
	return p1.(*ScanDevice).Type < p2.(*ScanDevice).Type
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


var scanSorters = map[string]LessFunc{
	"rack": rackOrder,
	"board": boardOrder,
	"device": deviceOrder,
	"type": typeOrder,
}


func (s *ScanResponse) ToScanTransformer() *Transformer {
	devices := s.ToScanDevices()

	items := make([]interface{}, len(devices))
	for i, device := range devices {
		items[i] = device
	}
	t :=  &Transformer{
		Items: items,
	}
	t.sortOrder = scanSorters
	return t
}




type Transformer struct {
	Sortable

	Items []interface{}

}

func (t *Transformer) Sort() {
	t.toSort = t.Items
	t.sort()
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

type LessFunc func(p1, p2 interface{}) bool

// ScanDevices implements the Sort interface, sorting the ScanDevice instances
// within.
type Sortable struct {
	toSort []interface{}
	less []LessFunc
	sortOrder map[string]LessFunc
}

// Sort sorts the device slice according to the less functions.
func (s *Sortable) sort() {
	sort.Sort(s)
}

// OrderBy
func (s *Sortable) OrderBy(by ...string) {
	for _, key := range by {
		fn := s.sortOrder[key]
		if fn != nil {
			s.less = append(s.less, fn)
		}
	}
}

// Len
func (s *Sortable) Len() int {
	return len(s.toSort)
}

// Swap
func (s *Sortable) Swap(i, j int) {
	s.toSort[i], s.toSort[j] = s.toSort[j], s.toSort[i]
}

// Less
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