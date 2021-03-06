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
	"sort"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	"github.com/vapor-ware/synse-cli/pkg/utils/exit"
	"github.com/vapor-ware/synse-client-go/synse/scheme"
)

func init() {
	cmdRead.Flags().BoolVarP(&flagNoHeader, "no-header", "n", false, "do not print out column headers")
	cmdRead.Flags().BoolVarP(&flagJSON, "json", "", false, "print output as JSON")
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
		exiter := exit.FromCmd(cmd)

		// Error out if device IDs and tag selectors are both specified.
		if len(args) != 0 && len(flagTags) != 0 {
			exiter.Err("cannot specify device IDs and device tags together")
		}

		// Error out if multiple output formats are specified.
		if flagJSON && flagYaml {
			exiter.Err("cannot use multiple formatting flags at once")
		}

		exiter.Err(serverRead(cmd.OutOrStdout(), args))
	},
}

func serverRead(out io.Writer, devices []string) error {
	log.Debug("creating new HTTP client")
	client, err := utils.NewSynseHTTPClient(flagContext, flagTLSCert)
	if err != nil {
		return err
	}

	var readings []*scheme.Read
	if len(devices) != 0 {
		for _, device := range devices {
			log.WithField("device", device).Debug("issuing HTTP read device request")
			response, err := client.ReadDevice(device)
			if err != nil {
				return err
			}
			readings = append(readings, response...)
		}
	} else {
		log.WithFields(log.Fields{
			"tags": flagTags,
			"ns":   flagNS,
		}).Debug("issuing HTTP read request")
		response, err := client.Read(scheme.ReadOptions{
			Tags: utils.NormalizeTags(flagTags),
			NS:   flagNS,
		})
		if err != nil {
			return err
		}
		readings = response
	}

	if len(readings) == 0 {
		log.Debug("no readings reported from server")
		return nil
	}

	printer := utils.NewPrinter(out, flagJSON, flagYaml, flagNoHeader)
	printer.SetHeader("ID", "VALUE", "UNIT", "TYPE", "TIMESTAMP")
	printer.SetRowFunc(serverReadRowFunc)

	sort.Sort(Readings(readings))
	return printer.Write(readings)
}
