package plugin

import (
	"os"
	"text/tabwriter"
	"text/template"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/utils"
	"github.com/vapor-ware/synse-server-grpc/go"
	"golang.org/x/net/context"
)

const (
	writeTmpl = "{{.ID}}\t{{.Action}}\t{{.Raw}}\n"
)

type writeOut struct {
	ID     string
	Action string
	Raw    string
}

var writeHeader = writeOut{
	ID:     "TRANSACTION ID",
	Action: "ACTION",
	Raw:    "RAW",
}

// pluginWriteCommand is a CLI sub-command for writing to a plugin
var pluginWriteCommand = cli.Command{
	Name:   "write",
	Usage:  "Write data directly to a plugin",
	Action: cmdWrite,
}

// cmdWrite is the action for pluginWriteCommand. It writes directly to
// the specified plugin.
func cmdWrite(c *cli.Context) error {
	err := utils.RequiresArgsInRange(4, 5, c)
	if err != nil {
		return err
	}

	rack := c.Args().Get(0)
	board := c.Args().Get(1)
	device := c.Args().Get(2)
	action := c.Args().Get(3)
	raw := c.Args().Get(4)

	var wd *synse.WriteData
	wd = &synse.WriteData{
		Action: action,
	}
	if raw != "" {
		wd.Raw = [][]byte{[]byte(raw)}
	}

	plugin, err := makeGrpcClient(c)
	if err != nil {
		return err
	}

	transactions, err := plugin.Write(context.Background(), &synse.WriteRequest{
		Device: device,
		Board:  board,
		Rack:   rack,
		Data:   []*synse.WriteData{wd},
	})
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 10, 1, 3, ' ', 0)

	tmpl, err := template.New("write").Parse(writeTmpl)
	err = tmpl.Execute(w, writeHeader)
	if err != nil {
		return err
	}

	for tid, ctx := range transactions.Transactions {
		raw := ""
		for _, i := range ctx.Raw {
			raw += string(i) + " "
		}

		out := writeOut{
			ID:     tid,
			Action: ctx.Action,
			Raw:    raw,
		}
		err = tmpl.Execute(w, out)
		if err != nil {
			return err
		}
	}

	w.Flush()
	return nil
}
