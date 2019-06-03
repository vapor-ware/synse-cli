// Synse CLI
// Copyright (c) 2019 Vapor IO
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package plugin

import (
	"context"
	"io"

	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	synse "github.com/vapor-ware/synse-server-grpc/go"
)

func init() {
	cmdWrite.Flags().BoolVarP(&flagNoHeader, "no-header", "n", false, "do not print out column headers")
	cmdWrite.Flags().BoolVarP(&flagJSON, "json", "", false, "print output as JSON")
	cmdWrite.Flags().BoolVarP(&flagYaml, "yaml", "", false, "print output as YAML")
	cmdWrite.Flags().BoolVarP(&flagWait, "wait", "w", false, "wait for the write to complete")
}

var cmdWrite = &cobra.Command{
	Use:   "write DEVICE ACTION [DATA]",
	Short: "Write to a specified device",
	Long: utils.Doc(`
		Write to a device managed by the plugin.

		The plugin defines the write ACTIONs which that device supports, as well
		as any requirements on the DATA. The DATA may not be required for all 
		devices/actions.

		By default, this command executes writes asynchronously, returning
		information about the transaction generated for the write. This transaction
		can be checked later via 'synse server transaction'. If the --wait flag
		is specified, this command will wait until the write has completed and
		will display the final status of the write transaction, indicating error
		or success.

		The output of this command can be formatted as a table (default), as
		JSON, or as YAML. If specifying the output format, only one flag may
		be used. Using multiple output format flags will result in an error.
	`),
	Args: cobra.RangeArgs(2, 3),
	Run: func(cmd *cobra.Command, args []string) {
		// Error out if multiple output formats are specified.
		if flagJSON && flagYaml {
			exitutil.Err("cannot use multiple formatting flags at once")
		}

		device := args[0]
		action := args[1]
		data := ""
		if len(args) == 3 {
			data = args[2]
		}

		if flagWait {
			exitutil.Err(pluginWriteSync(cmd.OutOrStdout(), device, action, data))
		} else {
			exitutil.Err(pluginWriteAsync(cmd.OutOrStdout(), device, action, data))
		}
	},
}

func pluginWriteAsync(out io.Writer, device, action, data string) error {
	conn, client, err := utils.NewSynseGrpcClient(flagContext, flagTLSCert)
	if err != nil {
		return err
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := client.WriteAsync(ctx, &synse.V3WritePayload{
		Selector: &synse.V3DeviceSelector{
			Id: device,
		},
		Data: []*synse.V3WriteData{{
			Action: action,
			Data:   []byte(data),
		}},
	})
	if err != nil {
		return err
	}

	var txns []*synse.V3WriteTransaction
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		txns = append(txns, resp)
	}

	if len(txns) == 0 {
		exitutil.Fatal("failed device write")
	}

	printer := utils.NewPrinter(out, flagJSON, flagYaml, flagNoHeader)
	printer.SetHeader("TRANSACTION", "ACTION", "DATA", "DEVICE")
	printer.SetRowFunc(pluginTransactionInfoRowFunc)

	return printer.Write(txns)
}

func pluginWriteSync(out io.Writer, device, action, data string) error {
	conn, client, err := utils.NewSynseGrpcClient(flagContext, flagTLSCert)
	defer conn.Close()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := client.WriteSync(ctx, &synse.V3WritePayload{
		Selector: &synse.V3DeviceSelector{
			Id: device,
		},
		Data: []*synse.V3WriteData{{
			Action: action,
			Data:   []byte(data),
		}},
	})
	if err != nil {
		return err
	}

	var txns []*synse.V3TransactionStatus
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		txns = append(txns, resp)
	}

	if len(txns) == 0 {
		exitutil.Fatal("failed device write")
	}

	printer := utils.NewPrinter(out, flagJSON, flagYaml, flagNoHeader)
	printer.SetHeader("ID", "STATUS", "MESSAGE", "CREATED", "UPDATED")
	printer.SetRowFunc(pluginTransactionStatusRowFunc)

	return printer.Write(txns)
}
