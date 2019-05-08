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
	"fmt"

	"github.com/vapor-ware/synse-cli/pkg/config"
	"github.com/vapor-ware/synse-client-go/synse"
)

func NewSynseHTTPClient() (synse.Client, error) {
	currentContexts := config.GetCurrentContext()
	serverCtx := currentContexts["server"]
	if serverCtx == nil {
		return nil, fmt.Errorf("cannot create HTTP client for server: no current server context")
	}

	return synse.NewHTTPClientV3(&synse.Options{
		Address: serverCtx.Context.Address,
	})
}
