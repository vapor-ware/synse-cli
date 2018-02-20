package sorters

import "github.com/vapor-ware/synse-cli/pkg/scheme"

// rackOrder compares two ScanDevices by rack id.
func rackOrder(p1, p2 interface{}) bool {
	return p1.(*scheme.ScanDevice).Rack < p2.(*scheme.ScanDevice).Rack
}

// boardOrder compares two ScanDevices by board id.
func boardOrder(p1, p2 interface{}) bool {
	return p1.(*scheme.ScanDevice).Board < p2.(*scheme.ScanDevice).Board
}

// deviceOrder compares two ScanDevices by device id.
func deviceOrder(p1, p2 interface{}) bool {
	return p1.(*scheme.ScanDevice).Device < p2.(*scheme.ScanDevice).Device
}

// typeOrder compares two ScanDevices by type.
func typeOrder(p1, p2 interface{}) bool {
	return p1.(*scheme.ScanDevice).Type < p2.(*scheme.ScanDevice).Type
}

// ScanLess is the collection of all Less comparison functions for the
// Synse Server scan command.
var ScanLess = map[string]LessFunc{
	"rack":   rackOrder,
	"board":  boardOrder,
	"device": deviceOrder,
	"type":   typeOrder,
}
