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
	"encoding/json"
	"io"

	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	"github.com/vapor-ware/synse-client-go/synse/scheme"
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
			utils.Err(serverWriteSync(cmd.OutOrStdout(), device, action, data))
		} else {
			utils.Err(serverWriteAsync(cmd.OutOrStdout(), device, action, data))
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

	if len(*response) == 0 {
		// TODO: on no reading, should it print a message "no readings",
		//   should it print nothing, or should it just print header info
		//   with no rows?
	}

	// Format output
	// FIXME: there is probably a way to clean this up / generalize this, but
	//   that can be done later.
	if flagJson {
		o, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return err
		}
		_, err = out.Write(append(o, '\n'))
		return err

	} else if flagYaml {
		o, err := yaml.Marshal(response)
		if err != nil {
			return err
		}
		_, err = out.Write(o)
		return err

	} else {
		w := utils.NewTabWriter(out)
		defer w.Flush()

		if !flagNoHeader {
			if err := printTransactionSummaryHeader(w); err != nil {
				return err
			}
		}

		for _, r := range *response {
			if err := printTransactionSummaryRow(w, &r); err != nil {
				return err
			}
		}
	}
	return nil
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

	if len(*response) == 0 {
		// TODO: on no reading, should it print a message "no readings",
		//   should it print nothing, or should it just print header info
		//   with no rows?
	}

	// Format output
	// FIXME: there is probably a way to clean this up / generalize this, but
	//   that can be done later.
	if flagJson {
		o, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return err
		}
		_, err = out.Write(append(o, '\n'))
		return err

	} else if flagYaml {
		o, err := yaml.Marshal(response)
		if err != nil {
			return err
		}
		_, err = out.Write(o)
		return err

	} else {
		w := utils.NewTabWriter(out)
		defer w.Flush()

		if !flagNoHeader {
			if err := printTransactionHeader(w); err != nil {
				return err
			}
		}

		for _, t := range *response {
			if err := printTransactionRow(w, &t); err != nil {
				return err
			}
		}
	}
	return nil
}
