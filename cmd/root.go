package cmd

import (
	"github.com/deciphernow/certifiable/cmd/issue"
	"github.com/spf13/cobra"
)

// Issue returns a command that issues new certificates.
func Issue() *cobra.Command {

	command := &cobra.Command{
		Use:   "issue",
		Short: "Issue one or more certificates.",
	}

	command.AddCommand(issue.Batch())
	command.AddCommand(issue.Single())

	return command
}

/*
[
	{
		"common": "localhost",
		"alternative": ["localhost"],
		"expire": 365,
		"client": true,
		"server": true
	}
]
*/
