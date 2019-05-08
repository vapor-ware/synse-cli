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

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	"github.com/vapor-ware/synse-client-go/synse/scheme"
)

var (
	filters []string
	sort    string
)

var cmdScan = &cobra.Command{
	Use:   "scan",
	Short: "Enumerate all devices",
	Long: heredoc.Doc(`
		Enumerate all devices available to Synse Server.

		Scan results can be sorted and filtered. This is particularly useful
		when a Synse Server instance manages many devices.

		Sorting is done via the '--sort' flag. The value for the flag specifies
		the fields that should be sorted. Multiple fields can be specified in
		a comma-separated string, e.g. "rack,board". The first field will is the
		primary sort key, the second field is the secondary sort key, etc.
		The following fields support sorting:
 		 - rack   - device
 		 - board  - type

		Filtering is done via the '--filter' flag. The value for the flag is
		a string in the format "KEY=VALUE", where KEY is the field to filter
		by, and VALUE is the desired value to filter against. The '--filter'
		flag can be used multiple times to specify multiple filters. The
		following fields support filtering:
 		 - rack   - type
		 - board

		Some examples:
		* Show only LED devices sorted by their rack, board, and device ids:
		  sysne server scan --sort "rack,board,device" --filter "type=led"

		* Show only temperature and pressure devices:
		  synse server scan --filter "type=temperature" --filter "type=pressure"

		For more information, see:
		https://vapor-ware.github.io/synse-server/#scan
	`),

	Run: func(cmd *cobra.Command, args []string) {
		utils.Err(serverScan(cmd.OutOrStdout()))
	},
}

func serverScan(out io.Writer) error {
	client, err := utils.NewSynseHTTPClient()
	if err != nil {
		return err
	}

	response, err := client.Scan(scheme.ScanOptions{})
	if err != nil {
		return err
	}

	// TODO: figure out output formatting
	o, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		return err
	}
	_, err = out.Write(append(o, '\n'))
	return err
}

// Add flag options to the command.
// FIXME
//cmd.PersistentFlags().StringArrayVar(&filters, "filter", []string{}, "set filter(s) for the output results")
//cmd.PersistentFlags().StringVar(&sort, "sort", "", "set the sorting constraints")
