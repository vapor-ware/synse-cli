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
	"gopkg.in/yaml.v2"
)

func init() {
	cmdStatus.Flags().BoolVarP(&flagNoHeader, "no-header", "n", false, "do not print out column headers")
	cmdStatus.Flags().BoolVarP(&flagJson, "json", "", false, "print output as JSON")
	cmdStatus.Flags().BoolVarP(&flagYaml, "yaml", "", false, "print output as YAML")
}

var cmdStatus = &cobra.Command{
	Use:   "status",
	Short: "Get the status of the server instance",
	Long: utils.Doc(`
		Get the connectivity status for the configured Synse Server instance.

		The output of this command can be formatted as a table (default), as
		JSON, or as YAML. If specifying the output format, only one flag may
		be used. Using multiple output format flags will result in an error.

		For more information, see:
		<underscore>https://vapor-ware.github.io/synse-server/#test</>
	`),
	Run: func(cmd *cobra.Command, args []string) {
		// Error out if multiple output formats are specified.
		if flagJson && flagYaml {
			utils.Err("cannot use multiple formatting flags at once")
		}

		utils.Err(serverStatus(cmd.OutOrStdout()))
	},
}

func serverStatus(out io.Writer) error {
	client, err := utils.NewSynseHTTPClient()
	if err != nil {
		return err
	}

	response, err := client.Status()
	if err != nil {
		return err
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
			if err := printStatusHeader(w); err != nil {
				return err
			}
		}

		if err := printStatusRow(w, response); err != nil {
			return err
		}
	}
	return nil
}
