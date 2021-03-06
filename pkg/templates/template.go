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

package templates

import (
	"github.com/MakeNowJust/heredoc"
)

var (
	// CmdVersionTemplate is the template string for the version command response.
	CmdVersionTemplate = heredoc.Doc(`
	synse:
	 version     : {{.Version}}
	 build date  : {{.BuildDate}}
	 git commit  : {{.Commit}}
	 git tag     : {{.Tag}}
	 go version  : {{.GoVersion}}
	 go compiler : {{.GoCompiler}}
	 platform    : {{.OS}}/{{.Arch}}
	`)
)
