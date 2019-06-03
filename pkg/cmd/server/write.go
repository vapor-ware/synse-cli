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

package server

import (
	"io"

	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	"github.com/vapor-ware/synse-client-go/synse/scheme"
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
		Write to a device managed by the Synse Server instance.

		A device's managing plugin defines the write ACTIONs which that device
		can support, as well as any requirements on the DATA. The DATA may not be
		required for all devices/actions.

		By default, this command executes writes asynchronously, returning
		information about the transaction generated for the write. This transaction
		can be checked later via 'synse server transaction'. If the --wait flag
		is specified, this command will wait until the write has completed and
		will display the final status of the write transaction, indicating error
		or success.

		The output of this command can be formatted as a table (default), as
		JSON, or as YAML. If specifying the output format, only one flag may
		be used. Using multiple output format flags will result in an error.

		For more information, see:
		<underscore>https://vapor-ware.github.io/synse-server/#write</>
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
			exitutil.Err(serverWriteSync(cmd.OutOrStdout(), device, action, data))
		} else {
			exitutil.Err(serverWriteAsync(cmd.OutOrStdout(), device, action, data))
		}
	},
}

func serverWriteAsync(out io.Writer, device, action, data string) error {
	client, err := utils.NewSynseHTTPClient()
	if err != nil {
		return err
	}

	response, err := client.WriteAsync(device, []scheme.WriteData{{
		Action: action,
		Data:   data,
	}})

	if len(response) == 0 {
		exitutil.Fatal("failed device write")
	}

	printer := utils.NewPrinter(out, flagJSON, flagYaml, flagNoHeader)
	printer.SetHeader("ID", "ACTION", "DATA", "DEVICE")
	printer.SetRowFunc(serverTransactionSummaryRowFunc)

	return printer.Write(response)
}

func serverWriteSync(out io.Writer, device, action, data string) error {
	client, err := utils.NewSynseHTTPClient()
	if err != nil {
		return err
	}

	response, err := client.WriteSync(device, []scheme.WriteData{{
		Action: action,
		Data:   data,
	}})

	if len(response) == 0 {
		exitutil.Fatal("failed device write")
	}

	printer := utils.NewPrinter(out, flagJSON, flagYaml, flagNoHeader)
	printer.SetHeader("ID", "STATUS", "MESSAGE", "CREATED", "UPDATED")
	printer.SetRowFunc(serverTransactionRowFunc)

	return printer.Write(response)
}
