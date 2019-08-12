package cmd

import (
	"fmt"

	"github.com/deciphernow/certifiable/versions"
	"github.com/spf13/cobra"
)

// Version returns a command that prints the version.
func Version() *cobra.Command {

	command := &cobra.Command{
		Use:   "version",
		Short: "Print the version and commit.",
		Run: func(command *cobra.Command, args []string) {
			fmt.Printf("certifiable %s (%s)\n", versions.Version(), versions.Commit())
		},
	}

	return command
}
