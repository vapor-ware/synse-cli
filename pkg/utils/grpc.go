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
	log "github.com/sirupsen/logrus"
	"github.com/vapor-ware/synse-cli/pkg/config"
	synse "github.com/vapor-ware/synse-server-grpc/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Errors relating to gRPC client creation.
var (
	ErrNoCurrentPluginCtx = errors.New("failed creating plugin gRPC client: no current plugin context")
	ErrInvalidPluginCtx   = errors.New("failed creating plugin gRPC client: specified context does not exist")
	ErrNotAPluginCtx      = errors.New("failed creating plugin gRPC client: specified context is not a plugin context")
)

// NewSynseGrpcClient creates a new instance of a Synse gRPC client
// for communicating with Synse plugins.
func NewSynseGrpcClient(ctx, certFile string) (*grpc.ClientConn, synse.V3PluginClient, error) {
	var pluginContext *config.ContextRecord

	if ctx == "" {
		// If no specific context is specified, get the current context.
		currentContexts := config.GetCurrentContext()
		pluginContext = currentContexts["plugin"]
		if pluginContext == nil {
			return nil, nil, ErrNoCurrentPluginCtx
		}
	} else {
		// Get the named context.
		pluginContext = config.GetContext(ctx)
		if pluginContext == nil {
			return nil, nil, ErrInvalidPluginCtx
		}
		if pluginContext.Type != "plugin" {
			return nil, nil, ErrNotAPluginCtx
		}
	}

	// If a cert file was provided, use it to override any cert file which may
	// already be configured for the context.
	if certFile != "" {
		pluginContext.Context.ClientCert = certFile
	}

	var dialOptions []grpc.DialOption
	if pluginContext.Context.ClientCert == "" {
		log.Debug("grpc client: with insecure")
		dialOptions = append(dialOptions, grpc.WithInsecure())
	} else {
		creds, err := credentials.NewClientTLSFromFile(pluginContext.Context.ClientCert, "")
		if err != nil {
			return nil, nil, err
		}
		log.WithField("cert", pluginContext.Context.ClientCert).Debug("grpc client: with credentials")
		dialOptions = append(dialOptions, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(pluginContext.Context.Address, dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	client := synse.NewV3PluginClient(conn)

	return conn, client, nil
}
