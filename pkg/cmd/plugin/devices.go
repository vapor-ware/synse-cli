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
	cmdDevices.Flags().BoolVarP(&flagNoHeader, "no-header", "n", false, "do not print out column headers")
	cmdDevices.Flags().BoolVarP(&flagJSON, "json", "", false, "print output as JSON")
	cmdDevices.Flags().BoolVarP(&flagYaml, "yaml", "", false, "print output as YAML")
	cmdDevices.Flags().StringSliceVarP(&flagTags, "tag", "t", []string{}, "specify tags to use as device selectors")
}

var cmdDevices = &cobra.Command{
	Use:   "devices",
	Short: "Display the devices managed by the plugin",
	Long: utils.Doc(`
		Display the devices exposed by the plugin.

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

		exiter.Err(pluginDevices(cmd.OutOrStdout()))
	},
}

func pluginDevices(out io.Writer) error {
	log.Debug("creating new gRPC client")
	conn, client, err := utils.NewSynseGrpcClient(flagContext, flagTLSCert)
	if err != nil {
		return err
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var tags []*synse.V3Tag
	for _, t := range utils.NormalizeTags(flagTags) {
		tag, err := utils.StringToTag(t)
		if err != nil {
			return err
		}
		tags = append(tags, tag)
	}

	log.WithField("tags", tags).Debug("issuing gRPC devices request")
	stream, err := client.Devices(ctx, &synse.V3DeviceSelector{
		Tags: tags,
	})
	if err != nil {
		return err
	}

	var devices []*synse.V3Device
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		devices = append(devices, resp)
	}

	if len(devices) == 0 {
		log.Debug("no devices reported from plugin")
		return nil
	}
	log.WithField("total", len(devices)).Debug("got devices from plugin")

	printer := utils.NewPrinter(out, flagJSON, flagYaml, flagNoHeader)
	printer.SetIntermediateYaml()
	printer.SetHeader("ID", "ALIAS", "TYPE", "INFO", "PLUGIN")
	printer.SetRowFunc(pluginDeviceRowFunc)

	sort.Sort(Devices(devices))
	return printer.Write(devices)
}
