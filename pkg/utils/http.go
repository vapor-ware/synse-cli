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

package utils

import (
	"github.com/pkg/errors"
	"github.com/vapor-ware/synse-cli/pkg/config"
	"github.com/vapor-ware/synse-client-go/synse"
)

// Errors relating to gRPC client creation.
var (
	ErrNoCurrentServerCtx = errors.New("failed creating server HTTP client: no current server context")
	ErrInvalidServerCtx   = errors.New("failed creating server HTTP client: specified context does not exist")
	ErrNotAServerCtx      = errors.New("failed creating server HTTP client: specified context is not a server context")
)

// NewSynseHTTPClient creates a new Synse HTTP client for communicating with
// Synse Server instances.
func NewSynseHTTPClient(ctx string, certFile string) (synse.Client, error) {
	var serverContext *config.ContextRecord

	if ctx == "" {
		// If no specific context is specified, get the current context.
		currentContexts := config.GetCurrentContext()
		serverContext = currentContexts["server"]
		if serverContext == nil {
			return nil, ErrNoCurrentServerCtx
		}
	} else {
		// Get the named context.
		serverContext = config.GetContext(ctx)
		if serverContext == nil {
			return nil, ErrInvalidServerCtx
		}
		if serverContext.Type != "server" {
			return nil, ErrNotAServerCtx
		}
	}

	return synse.NewHTTPClientV3(&synse.Options{
		Address: serverContext.Context.Address,
		TLS: synse.TLSOptions{
			CertFile: certFile,
		},
	})
}
