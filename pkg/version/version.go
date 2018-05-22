package version

import (
	"runtime"
)

var (
	// BuildDate is the timestamp for when the build happened.
	BuildDate string

	// GitCommit is the commit hash at which the binary was built.
	GitCommit string

	// GitTag is the git tag at which the binary was built.
	GitTag string

	// GoVersion is is the version of Go used to build the binary.
	GoVersion string

	// VersionString is the canonical version string for the binary.
	VersionString string
)

// BinVersion describes the version of the binary for the CLI.
//
// This should be populated via build-time args passed in for
// the corresponding variables.
type BinVersion struct {
	Arch          string
	BuildDate     string
	GitCommit     string
	GitTag        string
	GoVersion     string
	OS            string
	VersionString string
}

// Get gets the version information for the CLI. It builds
// a BinVersion using the variables that should be set as build-time
// arguments.
func Get() *BinVersion {
	return &BinVersion{
		Arch:          runtime.GOARCH,
		OS:            runtime.GOOS,
		BuildDate:     BuildDate,
		GitCommit:     GitCommit,
		GitTag:        GitTag,
		GoVersion:     GoVersion,
		VersionString: VersionString,
	}
}
