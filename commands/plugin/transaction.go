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
	transactionTmpl = "{{.ID}}\t{{.Status}}\t{{.State}}\t{{.Created}}\t{{.Updated}}\n"
)

type transactionOut struct {
	ID      string
	Status  string
	State   string
	Created string
	Updated string
}

var transactionHeader = transactionOut{
	ID:      "ID",
	Status:  "STATUS",
	State:   "STATE",
	Created: "CREATED",
	Updated: "UPDATED",
}

// pluginTransactionCommand is a CLI sub-command for getting transaction info from a plugin.
var pluginTransactionCommand = cli.Command{
	Name:   "transaction",
	Usage:  "Get transaction info from a plugin",
	Action: cmdTransaction,
}

// cmdTransaction is the action for pluginTransactionCommand. It prints out transaction info
// retrieved from the specified plugin.
func cmdTransaction(c *cli.Context) error {
	err := utils.RequiresArgsExact(1, c)
	if err != nil {
		return err
	}

	tid := c.Args().Get(0)

	plugin, err := makeGrpcClient(c)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 10, 1, 3, ' ', 0)

	tmpl, err := template.New("transaction").Parse(transactionTmpl)
	if err != nil {
		return err
	}
	err = tmpl.Execute(w, transactionHeader)
	if err != nil {
		return err
	}

	status, err := plugin.TransactionCheck(context.Background(), &synse.TransactionId{
		Id: tid,
	})
	if err != nil {
		return err
	}

	transaction := transactionOut{
		ID:      tid,
		Status:  status.Status.String(),
		State:   status.State.String(),
		Created: status.Created,
		Updated: status.Updated,
	}

	err = tmpl.Execute(w, transaction)
	if err != nil {
		return err
	}

	return w.Flush()
}
