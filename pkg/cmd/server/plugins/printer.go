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

package plugins

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/vapor-ware/synse-client-go/synse/scheme"
)

var (
	// ErrInvalidRowData is a printer function error which indicates that the
	// data type given to the printer is unexpected. This error should never be
	// induced by a user error, but may occur if there are changes in modeling.
	ErrInvalidRowData = errors.New("invalid row data")

	// ErrNilData is a printer function error which indicates that the value
	// passed to the printer is nil and can not be printed.
	ErrNilData = errors.New("row handler got nil data")
)

func serverPluginSummaryRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(*scheme.PluginMeta)
	if !ok {
		return nil, ErrInvalidRowData
	}
	if i == nil {
		return nil, ErrNilData
	}

	var isActive = "✓"
	if !i.Active {
		isActive = "✗"
	}

	return []interface{}{
		isActive,
		i.ID,
		i.Version.PluginVersion,
		i.Tag,
		i.Description,
	}, nil
}

func serverPluginRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(*scheme.Plugin)
	if !ok {
		return nil, ErrInvalidRowData
	}
	if i == nil {
		return nil, ErrNilData
	}

	var isActive = "✓"
	if !i.Active {
		isActive = "✗"
	}

	addr := fmt.Sprintf("%s://%s", i.Network.Protocol, i.Network.Address)

	return []interface{}{
		isActive,
		i.ID,
		i.Tag,
		addr,
		i.Health.Status,
		i.Health.Timestamp,
	}, nil
}

func serverPluginHealthRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(*scheme.PluginHealth)
	if !ok {
		return nil, ErrInvalidRowData
	}
	if i == nil {
		return nil, ErrNilData
	}

	return []interface{}{
		i.Status,
		len(i.Healthy),
		len(i.Unhealthy),
		i.Active,
		i.Inactive,
	}, nil
}
