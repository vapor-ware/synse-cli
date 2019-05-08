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

var cmdTags = &cobra.Command{
	Use:   "tags",
	Short: "",
	Long:  heredoc.Doc(``),

	Run: func(cmd *cobra.Command, args []string) {
		utils.Err(serverTags(cmd.OutOrStdout()))
	},
}

func serverTags(out io.Writer) error {
	client, err := utils.NewSynseHTTPClient()
	if err != nil {
		return err
	}

	response, err := client.Tags(scheme.TagsOptions{})
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
