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
	cmdRead.Flags().BoolVarP(&flagNoHeader, "no-header", "n", false, "do not print out column headers")
	cmdRead.Flags().BoolVarP(&flagJson, "json", "", false, "print output as JSON")
	cmdRead.Flags().BoolVarP(&flagYaml, "yaml", "", false, "print output as YAML")
	cmdRead.Flags().StringVarP(&flagNS, "ns", "", "", "default tag namespace for tags with no explicit namespace set")
	cmdRead.Flags().StringSliceVarP(&flagTags, "tag", "t", []string{}, "specify tags to use as device selectors")
}

var cmdRead = &cobra.Command{
	Use:   "read DEVICE...",
	Short: "Get current readings for available devices",
	Long: utils.Doc(`
		Get current reading data for the specified device(s).

		Devices can be specified in one of two ways: either by providing the
		ID of the device(s) to get the latest readings from, or by specifying
		a set of tags to filter devices by.

		Tags are strings with three components: a namespace (optional), an
		annotation (optional), and a label (required). They follow the format
		"namespace/annotation:label". Multiple tags can be specified either
		by calling the '--tag' flag multiple times, or by providing a comma
		separated list of tags. For example, the two lines below are equivalent:

		   --tag default/foo --tag default/type:bar
		   --tag default/foo,default/type:bar

		You cannot specify devices both by ID and tag. Doing so will result in
		an error.

		The output of this command can be formatted as a table (default), as
		JSON, or as YAML. If specifying the output format, only one flag may
		be used. Using multiple output format flags will result in an error.

		For more information, see:
		<underscore>https://vapor-ware.github.io/synse-server/#read</>
	`),
	Run: func(cmd *cobra.Command, args []string) {
		// Error out if device IDs and tag selectors are both specified.
		if len(args) != 0 && len(flagTags) != 0 {
			utils.Err("cannot specify device IDs and device tags together")
		}

		// Error out if multiple output formats are specified.
		if flagJson && flagYaml {
			utils.Err("cannot use multiple formatting flags at once")
		}

		utils.Err(serverRead(cmd.OutOrStdout(), args))
	},
}

func serverRead(out io.Writer, devices []string) error {
	client, err := utils.NewSynseHTTPClient()
	if err != nil {
		return err
	}

	var readings []scheme.Read
	if len(devices) != 0 {
		for _, device := range devices {
			response, err := client.ReadDevice(device, scheme.ReadOptions{})
			if err != nil {
				return err
			}
			readings = append(readings, *response...)
		}
	} else {
		response, err := client.Read(scheme.ReadOptions{
			Tags: utils.NormalizeTags(flagTags),
			NS:   flagNS,
		})
		if err != nil {
			return err
		}
		readings = *response
	}

	if len(readings) == 0 {
		// TODO: on no reading, should it print a message "no readings",
		//   should it print nothing, or should it just print header info
		//   with no rows?
	}

	// Format output
	// FIXME: there is probably a way to clean this up / generalize this, but
	//   that can be done later.
	if flagJson {
		o, err := json.MarshalIndent(readings, "", "  ")
		if err != nil {
			return err
		}
		_, err = out.Write(append(o, '\n'))
		return err

	} else if flagYaml {
		o, err := yaml.Marshal(readings)
		if err != nil {
			return err
		}
		_, err = out.Write(o)
		return err

	} else {
		w := utils.NewTabWriter(out)
		defer w.Flush()

		if !flagNoHeader {
			if err := printReadingHeader(w); err != nil {
				return err
			}
		}

		for _, reading := range readings {
			if err := printReadingRow(w, reading); err != nil {
				return err
			}
		}
	}
	return nil
}
