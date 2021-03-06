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

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	"github.com/vapor-ware/synse-cli/pkg/utils/exit"
	"github.com/vapor-ware/synse-client-go/synse/scheme"
)

func init() {
	cmdTags.Flags().BoolVarP(&flagNoHeader, "no-header", "n", false, "do not print out column headers")
	cmdTags.Flags().BoolVarP(&flagJSON, "json", "", false, "print output as JSON")
	cmdTags.Flags().BoolVarP(&flagYaml, "yaml", "", false, "print output as YAML")
	cmdTags.Flags().BoolVarP(&flagIds, "ids", "", false, "include id tags in the output")
	cmdTags.Flags().StringVarP(&flagNS, "ns", "", "", "default tag namespace for tags with no explicit namespace set")
}

var cmdTags = &cobra.Command{
	Use:   "tags",
	Short: "List tags associated with devices",
	Long: utils.Doc(`
		List tags currently associated with devices.

		The output of this command can be formatted as a table (default), as
		JSON, or as YAML. If specifying the output format, only one flag may
		be used. Using multiple output format flags will result in an error.

		For more information, see:
		<underscore>https://vapor-ware.github.io/synse-server/#tags</>
	`),
	Run: func(cmd *cobra.Command, args []string) {
		exiter := exit.FromCmd(cmd)

		// Error out if multiple output formats are specified.
		if flagJSON && flagYaml {
			exiter.Err("cannot use multiple formatting flags at once")
		}

		exiter.Err(serverTags(cmd.OutOrStdout()))
	},
}

func serverTags(out io.Writer) error {
	log.Debug("creating new HTTP client")
	client, err := utils.NewSynseHTTPClient(flagContext, flagTLSCert)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"ns":  flagNS,
		"ids": flagIds,
	}).Debug("issuing HTTP tags request")
	response, err := client.Tags(scheme.TagsOptions{
		NS:  []string{flagNS},
		IDs: flagIds,
	})
	if err != nil {
		return err
	}

	if len(response) == 0 {
		log.Debug("no tags reported by server")
		return nil
	}

	printer := utils.NewPrinter(out, flagJSON, flagYaml, flagNoHeader)
	printer.SetHeader("TAG")
	printer.SetRowFunc(serverTagsRowFunc)

	return printer.Write(response)
}
