package plugin

import (
	"io"
	"os"
	"text/tabwriter"
	"text/template"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/utils"
	"github.com/vapor-ware/synse-server-grpc/go"
	"golang.org/x/net/context"
)

const (
	readTmpl = "{{.ID}}\t{{.Type}}\t{{.Reading}}\n"
)

type readOut struct {
	ID      string
	Type    string
	Reading string
}

var readHeader = readOut{
	ID:      "ID",
	Type:    "TYPE",
	Reading: "READING",
}

// pluginReadCommand is a CLI sub-command for getting a reading from a plugin.
var pluginReadCommand = cli.Command{
	Name:   "read",
	Usage:  "Get a reading from a plugin",
	Action: cmdRead,
}

// cmdRead is the action for pluginReadCommand. It prints out a reading that was
// retrieved from the specified plugin.
func cmdRead(c *cli.Context) error { // nolint: gocyclo
	err := utils.RequiresArgsExact(3, c)
	if err != nil {
		return err
	}

	rack := c.Args().Get(0)
	board := c.Args().Get(1)
	device := c.Args().Get(2)

	plugin, err := makeGrpcClient(c)
	if err != nil {
		return err
	}

	stream, err := plugin.Read(context.Background(), &synse.ReadRequest{
		Device: device,
		Board:  board,
		Rack:   rack,
	})
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 10, 1, 3, ' ', 0)

	tmpl, err := template.New("read").Parse(readTmpl)
	if err != nil {
		return err
	}
	err = tmpl.Execute(w, readHeader)
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

		reading := readOut{
			ID:      device,
			Type:    resp.Type,
			Reading: resp.Value,
		}

		err = tmpl.Execute(w, reading)
		if err != nil {
			return err
		}

	}

	return w.Flush()
}
