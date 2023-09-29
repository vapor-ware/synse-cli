// Synse CLI
// Copyright (c) 2023 Vapor IO
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
	"github.com/vapor-ware/synse-cli/pkg/utils/exit"
)

var cmdUnset = &cobra.Command{
	Use:   "unset CONTEXT_NAME",
	Short: "Unset the current context",
	Long: utils.Doc(`
		Unset the current active context for the CLI.
	`),
	SuggestFor: []string{
		"change",
	},
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		exit.FromCmd(cmd).Err(unsetContext(args[0]))
	},
}

func unsetContext(name string) error {
	log.WithFields(log.Fields{
		"name": name,
	}).Debug("unsetting current context")

	return config.UnsetCurrentContext(name)
}
