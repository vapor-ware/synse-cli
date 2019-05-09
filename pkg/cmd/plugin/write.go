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
	"encoding/json"
	"io"

	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	synse "github.com/vapor-ware/synse-server-grpc/go"
	"gopkg.in/yaml.v2"
)

func init() {
	cmdWrite.Flags().BoolVarP(&flagNoHeader, "no-header", "n", false, "do not print out column headers")
	cmdWrite.Flags().BoolVarP(&flagJson, "json", "", false, "print output as JSON")
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
		if flagJson && flagYaml {
			utils.Err("cannot use multiple formatting flags at once")
		}

		device := args[0]
		action := args[1]
		data := ""
		if len(args) == 3 {
			data = args[2]
		}

		if flagWait {
			utils.Err(pluginWriteSync(cmd.OutOrStdout(), device, action, data))
		} else {
			utils.Err(pluginWriteAsync(cmd.OutOrStdout(), device, action, data))
		}
	},
}

func pluginWriteAsync(out io.Writer, device, action, data string) error {
	conn, client, err := utils.NewSynseGrpcClient()
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
		// TODO: on no reading, should it print a message "no readings",
		//   should it print nothing, or should it just print header info
		//   with no rows?
	}

	// Format output
	// FIXME: there is probably a way to clean this up / generalize this, but
	//   that can be done later.
	if flagJson {
		o, err := json.MarshalIndent(txns, "", "  ")
		if err != nil {
			return err
		}
		_, err = out.Write(append(o, '\n'))
		return err

	} else if flagYaml {
		o, err := yaml.Marshal(txns)
		if err != nil {
			return err
		}
		_, err = out.Write(o)
		return err

	} else {
		w := utils.NewTabWriter(out)
		defer w.Flush()

		if !flagNoHeader {
			if err := printTransactionInfoHeader(w); err != nil {
				return err
			}
		}

		for _, t := range txns {
			if err := printTransactionInfoRow(w, t); err != nil {
				return err
			}
		}
	}
	return nil
}

func pluginWriteSync(out io.Writer, device, action, data string) error {
	conn, client, err := utils.NewSynseGrpcClient()
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
		// TODO: on no reading, should it print a message "no readings",
		//   should it print nothing, or should it just print header info
		//   with no rows?
	}

	// Format output
	// FIXME: there is probably a way to clean this up / generalize this, but
	//   that can be done later.
	if flagJson {
		o, err := json.MarshalIndent(txns, "", "  ")
		if err != nil {
			return err
		}
		_, err = out.Write(append(o, '\n'))
		return err

	} else if flagYaml {
		o, err := yaml.Marshal(txns)
		if err != nil {
			return err
		}
		_, err = out.Write(o)
		return err

	} else {
		w := utils.NewTabWriter(out)
		defer w.Flush()

		if !flagNoHeader {
			if err := printTransactionStatusHeader(w); err != nil {
				return err
			}
		}

		for _, t := range txns {
			if err := printTransactionStatusRow(w, t); err != nil {
				return err
			}
		}
	}
	return nil
}
