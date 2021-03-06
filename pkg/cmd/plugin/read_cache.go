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
	"sort"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	"github.com/vapor-ware/synse-cli/pkg/utils/exit"
	synse "github.com/vapor-ware/synse-server-grpc/go"
)

func init() {
	cmdReadCache.Flags().BoolVarP(&flagNoHeader, "no-header", "n", false, "do not print out column headers")
	cmdReadCache.Flags().BoolVarP(&flagJSON, "json", "", false, "print output as JSON")
	cmdReadCache.Flags().BoolVarP(&flagYaml, "yaml", "", false, "print output as YAML")
	cmdReadCache.Flags().StringVarP(&flagStart, "start", "s", "", "timestamp specifying the starting bound for windowing")
	cmdReadCache.Flags().StringVarP(&flagEnd, "end", "e", "", "timestamp specifying the ending bound for windowing")
}

var cmdReadCache = &cobra.Command{
	Use:   "read-cache",
	Short: "Get cached readings for available devices",
	Long: utils.Doc(`
		Get cached readings for available devices.

		The readings returned by this command are cached by the plugin. A start
		and end bound can be provided to window the readings within a time
		period. It is recommended to bound the request start/end times to limit
		the potentially large number of readings that would be provided otherwise.

		The start and end bounding timestamps should be specified in FRC3339
		format. An invalidly formatted timestamp may render the bound ineffective.

		The output of this command can be formatted as a table (default), as
		JSON, or as YAML. If specifying the output format, only one flag may
		be used. Using multiple output format flags will result in an error.

		The default table view only provides a summary of the data. To see
		see the data in its entirety, use the JSON or YAML output formats.
	`),
	Run: func(cmd *cobra.Command, args []string) {
		exiter := exit.FromCmd(cmd)

		// Error out if multiple output formats are specified.
		if flagJSON && flagYaml {
			exiter.Err("cannot use multiple formatting flags at once")
		}

		exiter.Err(pluginReadCache(cmd.OutOrStdout()))
	},
}

func pluginReadCache(out io.Writer) error {
	log.Debug("creating new gRPC client")
	conn, client, err := utils.NewSynseGrpcClient(flagContext, flagTLSCert)
	if err != nil {
		return err
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.WithFields(log.Fields{
		"start": flagStart,
		"end":   flagEnd,
	}).Debug("issuing gRPC read cache request")
	stream, err := client.ReadCache(ctx, &synse.V3Bounds{
		Start: flagStart,
		End:   flagEnd,
	})
	if err != nil {
		return err
	}

	var readings []*synse.V3Reading
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		readings = append(readings, resp)
	}

	if len(readings) == 0 {
		log.Debug("no cached readings reported by plugin")
		return nil
	}

	printer := utils.NewPrinter(out, flagJSON, flagYaml, flagNoHeader)
	printer.SetIntermediateYaml()
	printer.SetHeader("ID", "VALUE", "UNIT", "TYPE", "TIMESTAMP")
	printer.SetRowFunc(pluginReadingRowFunc)
	printer.SetTransformFunc(pluginReadTransformer)

	sort.Sort(Readings(readings))
	return printer.Write(readings)
}
