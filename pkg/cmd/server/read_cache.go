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
	cmdReadCache.Flags().BoolVarP(&flagNoHeader, "no-header", "n", false, "do not print out column headers")
	cmdReadCache.Flags().BoolVarP(&flagJson, "json", "", false, "print output as JSON")
	cmdReadCache.Flags().BoolVarP(&flagYaml, "yaml", "", false, "print output as YAML")
	cmdReadCache.Flags().StringVarP(&flagStart, "start", "s", "", "timestamp specifying the starting bound for windowing")
	cmdReadCache.Flags().StringVarP(&flagEnd, "end", "e", "", "timestamp specifying the ending bound for windowing")
}

var cmdReadCache = &cobra.Command{
	Use:   "read-cache",
	Short: "Get cached readings for available devices",
	Long: utils.Doc(`
		Get cached readings for all devices available to the server.

		The readings returned by this command are cached by the plugin. A start
		and end bound can be provided to window the readings within a time
		period. It is recommended to bound the request start/end times to limit
		the potentially large number of readings that would be provided otherwise.

		The start and end bounding timestamps should be specified in FRC3339
		format. An invalidly formatted timestamp may render the bound ineffective.

		The output of this command can be formatted as a table (default), as
		JSON, or as YAML. If specifying the output format, only one flag may
		be used. Using multiple output format flags will result in an error.

		For more information, see:
		<underscore>https://vapor-ware.github.io/synse-server/#read-cache</>
	`),

	Run: func(cmd *cobra.Command, args []string) {
		exitutil.Err(serverReadCache(cmd.OutOrStdout()))
	},
}

func serverReadCache(out io.Writer) error {
	client, err := utils.NewSynseHTTPClient()
	if err != nil {
		return err
	}

	response, err := client.ReadCache(scheme.ReadCacheOptions{
		Start: flagStart,
		End:   flagEnd,
	})
	if err != nil {
		return err
	}

	if len(response) == 0 {
		exitutil.Exitf(0, "No readings found.")
	}

	printer := utils.NewPrinter(out, flagJson, flagYaml, flagNoHeader)
	printer.SetHeader("ID", "VALUE", "UNIT", "TYPE", "TIMESTAMP")
	printer.SetRowFunc(serverReadRowFunc)

	return printer.Write(response)
}
