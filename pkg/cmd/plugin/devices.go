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
	cmdDevices.Flags().BoolVarP(&flagNoHeader, "no-header", "n", false, "do not print out column headers")
	cmdDevices.Flags().BoolVarP(&flagJson, "json", "", false, "print output as JSON")
	cmdDevices.Flags().BoolVarP(&flagYaml, "yaml", "", false, "print output as YAML")
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
		// Error out if multiple output formats are specified.
		if flagJson && flagYaml {
			utils.Err("cannot use multiple formatting flags at once")
		}

		utils.Err(pluginDevices(cmd.OutOrStdout()))
	},
}

func pluginDevices(out io.Writer) error {
	conn, client, err := utils.NewSynseGrpcClient()
	if err != nil {
		return err
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := client.Devices(ctx, &synse.V3DeviceSelector{})
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
		// TODO: on no reading, should it print a message "no readings",
		//   should it print nothing, or should it just print header info
		//   with no rows?
	}

	// Format output
	// FIXME: there is probably a way to clean this up / generalize this, but
	//   that can be done later.
	if flagJson {
		o, err := json.MarshalIndent(devices, "", "  ")
		if err != nil {
			return err
		}
		_, err = out.Write(append(o, '\n'))
		return err

	} else if flagYaml {
		o, err := yaml.Marshal(devices)
		if err != nil {
			return err
		}
		_, err = out.Write(o)
		return err

	} else {
		w := utils.NewTabWriter(out)
		defer w.Flush()

		if !flagNoHeader {
			if err := printDeviceHeader(w); err != nil {
				return err
			}
		}

		for _, device := range devices {
			if err := printDeviceRow(w, device); err != nil {
				return err
			}
		}
	}
	return nil
}
