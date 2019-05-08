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

	"github.com/vapor-ware/synse-cli/pkg/utils"
	"github.com/vapor-ware/synse-client-go/synse/scheme"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

var (
	start string
	end   string
)

var cmdReadCache = &cobra.Command{
	Use:   "read-cache",
	Short: "Get cached readings from all configured plugins",
	Long: heredoc.Doc(`
		Get the cached readings from all of the plugins registered with Synse Server.

		This command operates on all plugins, so it does not require any routing
		information to be specified. Start and end timestamps can be set in order
		to bound the reading data. It is suggested to use timestamp bounding to
		keep output manageable.

		The start and end timestamps should be formatted in RFC3339 or RFC3339Nano
		format. If no bounding timestamps are specified, all readings will be
		returned.

		For more information, see:
		https://vapor-ware.github.io/synse-server/#read-cached
	`),

	Run: func(cmd *cobra.Command, args []string) {
		utils.Err(serverReadCache(cmd.OutOrStdout()))
	},
}

func serverReadCache(out io.Writer) error {
	client, err := utils.NewSynseHTTPClient()
	if err != nil {
		return err
	}

	response, err := client.ReadCache(scheme.ReadCacheOptions{})
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
//cmd.PersistentFlags().StringVar(&start, "start", "", "the starting timestamp bound for the read cache data")
//cmd.PersistentFlags().StringVar(&end, "end", "", "the ending timestamp bound for the read cache data")
