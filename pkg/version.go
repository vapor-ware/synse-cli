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

package pkg

import (
	"runtime"
)

var (
	// BuildDate is the timestamp for when the build happened.
	BuildDate string

	// Commit is the commit hash at which the binary was built.
	Commit string

	// Tag is the tag name at which the binary was built.
	Tag string

	// GoVersion is the version of Go used to build the binary.
	GoVersion string

	// Version is the canonical version string for the binary.
	Version string
)

// BinVersion describes the version of the binary for the CLI.
//
// This should be populated via build-time args passed in for
// the corresponding variables.
type BinVersion struct {
	Arch       string
	BuildDate  string
	Commit     string
	Tag        string
	GoCompiler string
	GoVersion  string
	OS         string
	Version    string
}

// GetVersion gets the version information for the CLI. It builds
// a BinVersion using the variables that should be set as build-time
// arguments.
func GetVersion() *BinVersion {
	return &BinVersion{
		Arch:       runtime.GOARCH,
		OS:         runtime.GOOS,
		BuildDate:  BuildDate,
		Commit:     Commit,
		Tag:        Tag,
		GoCompiler: runtime.Compiler,
		GoVersion:  GoVersion,
		Version:    Version,
	}
}
