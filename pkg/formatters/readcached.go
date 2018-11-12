package formatters

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// pterryReadCached the pretty output format for `readcache` requests
	prettyReadCached = "{{.Location.Rack}}\t{{.Location.Board}}\t{{.Location.Device}}\t{{.Info}}\t{{.Type}}\t{{.Value}}\t{{.Unit}}\t{{.Timestamp}}\n"
)

// newReadCachedFormat is the handler for readcached commands that is used by the
// Formatter to add new read data to the format context.
func newReadCachedFormat(data interface{}) (interface{}, error) {
	read, ok := data.(scheme.ReadCached)
	if !ok {
		return nil, fmt.Errorf("formatter data %T not of type scheme.ReadCached", data)
	}

	return scheme.ReadCachedOutput{
		Location: scheme.DeviceLocation{
			Rack:   read.Location.Rack,
			Board:  read.Location.Board,
			Device: read.Location.Device,
		},
		Info:      read.Info,
		Type:      read.Type,
		Value:     fmt.Sprintf("%v", read.Value),
		Unit:      read.Unit.Symbol,
		Timestamp: utils.ParseTimestamp(read.Timestamp),
	}, nil
}

// NewReadCachedFormatter creates a new instance of a Formatter configured
// for the readcached command.
func NewReadCachedFormatter(c *cli.Context) *Formatter {
	f := NewFormatter(c, newReadCachedFormat)
	f.Template = prettyReadCached
	f.Decoder = &scheme.ReadCachedOutput{}

	return f
}
