package filters

import "github.com/vapor-ware/synse-cli/pkg/scheme"

// typeFilter checks if a ScanDevice has the given type.
func typeFilter(i interface{}, v string) bool {
	return i.(*scheme.ScanDevice).Type == v
}

// rackFilter checks if a ScanDevice has the given rack id.
func rackFilter(i interface{}, v string) bool {
	return i.(*scheme.ScanDevice).Rack == v
}

// boardFilter checks if a ScanDevice has the given board id.
func boardFilter(i interface{}, v string) bool {
	return i.(*scheme.ScanDevice).Board == v
}

// ScanFilters is the collection of all filtering functions for the
// Synse Server scan command.
var ScanFilters = map[string]FilterFunc{
	"type":  typeFilter,
	"rack":  rackFilter,
	"board": boardFilter,
}
