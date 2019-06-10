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
	"github.com/vapor-ware/synse-cli/pkg/utils/exit"
	synse "github.com/vapor-ware/synse-server-grpc/go"
)

func init() {
	cmdRead.Flags().BoolVarP(&flagNoHeader, "no-header", "n", false, "do not print out column headers")
	cmdRead.Flags().BoolVarP(&flagJSON, "json", "", false, "print output as JSON")
	cmdRead.Flags().BoolVarP(&flagYaml, "yaml", "", false, "print output as YAML")
	cmdRead.Flags().StringSliceVarP(&flagTags, "tag", "t", []string{}, "specify tags to use as device selectors")
}

var cmdRead = &cobra.Command{
	Use:   "read [DEVICE...]",
	Short: "Get current readings for available devices",
	Long: utils.Doc(`
		Get current reading data for available devices.

		The output of this command can be formatted as a table (default), as
		JSON, or as YAML. If specifying the output format, only one flag may
		be used. Using multiple output format flags will result in an error.

		The default table view only provides a summary of the data. To see
		see the data in its entirety, use the JSON or YAML output formats.
	`),
	Run: func(cmd *cobra.Command, args []string) {
		exiter := exit.FromCmd(cmd)

		// Error out if device IDs and tag selectors are both specified.
		if len(args) != 0 && len(flagTags) != 0 {
			exiter.Err("cannot specify device IDs and device tags together")
		}

		// Error out if multiple output formats are specified.
		if flagJSON && flagYaml {
			exiter.Err("cannot use multiple formatting flags at once")
		}

		exiter.Err(pluginRead(cmd.OutOrStdout(), args))
	},
}

func pluginRead(out io.Writer, devices []string) error {
	conn, client, err := utils.NewSynseGrpcClient(flagContext, flagTLSCert)
	if err != nil {
		return err
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var readings []*synse.V3Reading

	if len(devices) == 0 {
		var tags []*synse.V3Tag
		for _, t := range utils.NormalizeTags(flagTags) {
			tag, err := utils.StringToTag(t)
			if err != nil {
				return err
			}
			tags = append(tags, tag)
		}

		stream, err := client.Read(ctx, &synse.V3ReadRequest{
			Selector: &synse.V3DeviceSelector{
				Tags: tags,
			},
		})
		if err != nil {
			return err
		}

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
	} else {
		for _, device := range devices {
			stream, err := client.Read(ctx, &synse.V3ReadRequest{
				Selector: &synse.V3DeviceSelector{
					Id: device,
				},
			})
			if err != nil {
				return err
			}

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
		}
	}

	if len(readings) == 0 {
		return nil
	}

	printer := utils.NewPrinter(out, flagJSON, flagYaml, flagNoHeader)
	printer.SetIntermediateYaml()
	printer.SetHeader("ID", "VALUE", "UNIT", "TYPE", "TIMESTAMP")
	printer.SetRowFunc(pluginReadingRowFunc)

	return printer.Write(readings)
}
