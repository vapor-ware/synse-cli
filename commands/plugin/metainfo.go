package plugin

import (
	"io"
	"os"
	"text/tabwriter"
	"text/template"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-server-grpc/go"
	"golang.org/x/net/context"
)

const (
	metaTmpl = "{{.ID}}\t{{.Type}}\t{{.Model}}\t{{.Protocol}}\t{{.Rack}}\t{{.Board}}\n"
)

type scanOut struct {
	ID       string
	Type     string
	Model    string
	Protocol string
	Rack     string
	Board    string
}

var metaHeader = scanOut{
	ID:       "ID",
	Type:     "TYPE",
	Model:    "MODEL",
	Protocol: "PROTOCOL",
	Rack:     "RACK",
	Board:    "BOARD",
}

// pluginMetainfoCommand is a CLI sub-command for getting metainfo from a plugin.
var pluginMetainfoCommand = cli.Command{
	Name:   "meta",
	Usage:  "Get the metainformation from a plugin",
	Action: cmdMeta,
}

// cmdMeta is the action for pluginMetainfoCommand. It prints out the meta-information
// provided by the specified plugin.
func cmdMeta(c *cli.Context) error {
	plugin, err := makeGrpcClient(c)
	if err != nil {
		return err
	}

	stream, err := plugin.Metainfo(
		context.Background(),
		&synse.MetainfoRequest{},
	)

	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 10, 1, 3, ' ', 0)

	tmpl, err := template.New("meta").Parse(metaTmpl)
	err = tmpl.Execute(w, metaHeader)
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		metaInfo := scanOut{
			ID:       resp.Uid,
			Type:     resp.Type,
			Model:    resp.Model,
			Protocol: resp.Protocol,
			Rack:     resp.Location.Rack,
			Board:    resp.Location.Board,
		}

		err = tmpl.Execute(w, metaInfo)
		if err != nil {
			return err
		}
	}

	w.Flush()
	return nil
}
