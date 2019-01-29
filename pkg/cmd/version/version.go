package version

import (
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/version"
)

const (
	// TODO: perhaps move this and register this in the templates package?
	versionTemplate = `synse:
 version     : {{.VersionString}}
 build date  : {{.BuildDate}}
 git commit  : {{.GitCommit}}
 git tag     : {{.GitTag}}
 go version  : {{.GoVersion}}
 go compiler : {{.GoCompiler}}
 platform    : {{.OS}}/{{.Arch}}
`
)

// New returns a new instance of the 'version' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print CLI version information",
		Long: `Print the version and build-time information for the CLI binary.

Version information for other Synse components (e.g. server, plugin)
can be printed from corresponding CLI sub-command.

This command will print something like the following:

  synse:
   version     : local
   build date  : 2019-01-15T16:44:34
   git commit  : 9901437
   git tag     : 3.0.0
   go version  : go1.11.4
   go compiler : gc
   platform    : darwin/amd64

- version: The semantic version of the CLI binary.
- build date: The date and time at which the binary was built.
- git commit: The commit at which the binary was built.
- git tag: The tag at which the binary was built.
- go version: The version of Go used to build the binary.
- go compiler: The compiler used to build the binary.
- platform: The operating system and architecture of the system used to
    build the binary.

More information about a specific release version can be found on the
project's GitHub page: https://github.com/vapor-ware/synse-cli/releases`,

		Run: func(cmd *cobra.Command, args []string) {
			tmpl, err := template.New("version").Parse(versionTemplate)
			if err != nil {
				// TODO: consistent error exit
				log.Fatal(err)
			}

			if err := tmpl.ExecuteTemplate(os.Stdout, "version", version.Get()); err != nil {
				// TODO: consistent error exit
				log.Fatal(err)
			}
		},
	}

	return cmd
}
