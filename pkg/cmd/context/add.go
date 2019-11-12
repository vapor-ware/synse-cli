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
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/config"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	"github.com/vapor-ware/synse-cli/pkg/utils/exit"
)

func init() {
	cmdAdd.Flags().BoolVarP(&flagSet, "set", "", false, "set as the current context after adding")
	cmdAdd.Flags().StringVarP(&flagClientCert, "tlscert", "", "", "path to TLS certificate file (e.g. ./synse.pem)")
}

var cmdAdd = &cobra.Command{
	Use:   "add TYPE NAME ADDRESS",
	Short: "Add a new context",
	Long: utils.Doc(`
		Add a new context to the synse configuration.

		Each context serves as a reference to an instance of a Synse component
		which you can interact with. The contexts are persisted, making it
		easy to interact with different components.

		In order to add a context, you must specify the following args:

		<bold>TYPE</>    : The type of the component.
		<bold>NAME</>    : The name that will be used to reference the context.
		<bold>ADDRESS</> : The address of the component.

		Currently, the supported types are:
		- plugin
		- server
	`),
	SuggestFor: []string{
		"new",
	},
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		exit.FromCmd(cmd).Err(
			addContext(args[0], args[1], args[2]),
		)
	},
}

func addContext(ctxType, ctxName, ctxAddress string) error {
	log.WithFields(log.Fields{
		"type":    ctxType,
		"name":    ctxName,
		"address": ctxAddress,
		"tlscert": flagClientCert,
	}).Debug("adding new context")

	// Verify that the provided context type is supported.
	if ctxType != "plugin" && ctxType != "server" {
		return fmt.Errorf("unsupported context type: %s", ctxType)
	}

	err := config.AddContext(&config.ContextRecord{
		Name: ctxName,
		Type: ctxType,
		Context: config.Context{
			Address:    ctxAddress,
			ClientCert: flagClientCert,
		},
	})
	if err != nil {
		return err
	}

	if flagSet {
		log.Debug("setting new context as current context")
		return config.SetCurrentContext(ctxName)
	}
	return nil
}
