package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Execute the command line interface.
func Execute() {

	command := &cobra.Command{
		Use:   "certifiable",
		Short: "A command line tool for generating certificates and keys.",
	}

	command.AddCommand(Issue())
	command.AddCommand(Version())

	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
