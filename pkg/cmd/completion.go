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

package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	"github.com/vapor-ware/synse-cli/pkg/utils/exit"
)

var cmdCompletion = &cobra.Command{
	Use:   "completion",
	Short: "Generate bash completion scripts",
	Long: utils.Doc(`
		Generate bash completion scripts.

		To load bash completion for the current session, run:

		  <bold>. <(synse completion)</>

		To configure your bash shell to load synse completion for all
		new sessions, add the above to your bashrc, e.g.

		  <bold>echo ". <(synse completion)" >> ~/.bashrc</>
	`),

	Run: func(cmd *cobra.Command, args []string) {
		exit.FromCmd(cmd).Err(
			rootCmd.GenBashCompletion(os.Stdout),
		)
	},
}

const bashCompletionFunc = `
__synse_list_ctxs() {
	local synse_output
	if synse_output=$(synse context list --no-header 2>/dev/null); then
		out=($(echo "$synse_output" | awk '{print $2}'))
		COMPREPLY=( $( compgen -W "${out[*]}" -- "$cur" ) )
	fi
}

__synse_server_list_devices() {
	local synse_output
	if synse_output=$(synse server scan --no-header 2>/dev/null); then
		out=($(echo "$synse_output" | awk '{print $1}'))
		COMPREPLY=( $( compgen -W "${out[*]}" -- "$cur" ) )
	fi
}

__synse_server_list_txns() {
	local synse_output
	if synse_output=$(synse server transaction --no-header 2>/dev/null); then
		out=($(echo "$synse_output" | awk '{print $1}'))
		COMPREPLY=( $( compgen -W "${out[*]}" -- "$cur" ) )
	fi
}

__synse_server_list_plugins() {
	local synse_output
	if synse_output=$(synse server plugins list --no-header 2>/dev/null); then
		out=($(echo "$synse_output" | awk '{print $2}'))
		COMPREPLY=( $( compgen -W "${out[*]}" -- "$cur" ) )
	fi
}

__synse_plugin_list_devices() {
	local synse_output
	if synse_output=$(synse plugin devices --no-header 2>/dev/null); then
		out=($(echo "$synse_output" | awk '{print $1}'))
		COMPREPLY=( $( compgen -W "${out[*]}" -- "$cur" ) )
	fi
}

__synse_plugin_list_txns() {
	local synse_output
	if synse_output=$(synse plugin transaction --no-header 2>/dev/null); then
		out=($(echo "$synse_output" | awk '{print $1}'))
		COMPREPLY=( $( compgen -W "${out[*]}" -- "$cur" ) )
	fi
}

__custom_func() {
	case ${last_command} in 
		synse_server_read | \
		synse_server_info | \
		synse_server_write)
			__synse_server_list_devices
			return
			;;
		synse_server_transaction)
			__synse_server_list_txns
			return
			;;
		synse_server_plugins_info)
			__synse_server_list_plugins
			return
			;;
		synse_context_remove | \
		synse_context_set)
			__synse_list_ctxs
			return
			;;
		synse_plugin_read | \
		synse_plugin_write)
			__synse_plugin_list_devices
			return
			;;
		synse_plugin_transaction)
			__synse_plugin_list_txns
			return
			;;
		*)
			;;
	esac
}
`
