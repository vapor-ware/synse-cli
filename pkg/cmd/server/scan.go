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
	cmdScan.Flags().BoolVarP(&flagNoHeader, "no-header", "n", false, "do not print out column headers")
	cmdScan.Flags().BoolVarP(&flagJson, "json", "", false, "print output as JSON")
	cmdScan.Flags().BoolVarP(&flagYaml, "yaml", "", false, "print output as YAML")
	cmdScan.Flags().BoolVarP(&flagFull, "full", "f", false, "display the full scan record")
	cmdScan.Flags().BoolVarP(&flagForce, "force", "", false, "force a cache rebuild on the server")
	cmdScan.Flags().StringVarP(&flagNS, "ns", "", "", "default tag namespace for tags with no explicit namespace set")
	cmdScan.Flags().StringSliceVarP(&flagTags, "tag", "t", []string{}, "specify tags to use as device selectors")
}

var cmdScan = &cobra.Command{
	Use:   "scan",
	Short: "List devices available to the server",
	Long: utils.Doc(`
		Enumerate devices available to Synse Server.

		If no tags are specified, this command will enumerate all devices.
		Tags can be specified to filter the results to only include the
		devices which match the tag set.

		Tags are strings with three components: a namespace (optional), an
		annotation (optional), and a label (required). They follow the format
		"namespace/annotation:label". Multiple tags can be specified either
		by calling the '--tag' flag multiple times, or by providing a comma
		separated list of tags. For example, the two lines below are equivalent:

		   --tag default/foo --tag default/type:bar
		   --tag default/foo,default/type:bar

		The output of this command can be formatted as a table (default), as
		JSON, or as YAML. If specifying the output format, only one flag may
		be used. Using multiple output format flags will result in an error.

		For more information, see:
		<underscore>https://vapor-ware.github.io/synse-server/#scan</>
	`),

	Run: func(cmd *cobra.Command, args []string) {
		// Error out if multiple output formats are specified.
		if flagJson && flagYaml {
			utils.Err("cannot use multiple formatting flags at once")
		}

		utils.Err(serverScan(cmd.OutOrStdout()))
	},
}

func serverScan(out io.Writer) error {
	client, err := utils.NewSynseHTTPClient()
	if err != nil {
		return err
	}

	response, err := client.Scan(scheme.ScanOptions{
		Tags:  utils.NormalizeTags(flagTags),
		Force: flagForce,
		NS:    flagNS,
	})
	if err != nil {
		return err
	}

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
			if err := printScanHeader(w, flagFull); err != nil {
				return err
			}
		}

		for _, dev := range *response {
			if err := printScanRow(w, &dev, flagFull); err != nil {
				return err
			}
		}
	}
	return nil
}
