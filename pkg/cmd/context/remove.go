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

package context

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/config"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

func init() {
	cmdRemove.Flags().BoolVarP(&flagAll, "all", "a", false, "remove all contexts")
}

var flagAll bool

var cmdRemove = &cobra.Command{
	Use:   "remove [CONTEXT_NAME]",
	Short: "Remove a context record",
	Long: utils.Doc(`
		Remove a context record from the synse configuration.

		The context record to remove should be specified by name. If the
		context being removed is the current context, no new current context
		will be set to replace it -- this must be done manually via 
		'synse context set'.

		If the --all flag is set, this will purge all contexts.
	`),
	Aliases: []string{
		"rm",
	},
	SuggestFor: []string{
		"delete",
		"del",
	},
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := ""
		if len(args) != 0 {
			name = args[0]
		}
		removeContext(name)
	},
}

func removeContext(name string) {
	if flagAll {
		log.Debug("purging all contexts")
		config.Purge()
		return
	}

	log.WithFields(log.Fields{
		"name": name,
	}).Debug("removing context")
	config.RemoveContext(name)
}