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

	"github.com/vapor-ware/synse-cli/pkg/config"
)

func contextRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(config.ContextRecord)
	if !ok {
		return nil, fmt.Errorf("invalid row data: %T", data)
	}
	if i == (config.ContextRecord{}) {
		return nil, fmt.Errorf("got empty context record")
	}

	isCurrent := " "
	if config.IsCurrentContext(&i) {
		isCurrent = "*"
	}

	return []interface{}{
		isCurrent,
		i.Name,
		i.Type,
		i.Context.Address,
	}, nil
}
