package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// To be assigned by govvv.

// GitBranch holds the name of the current branch.
var GitBranch string

// GitSummary holds the summary of the current repo.
var GitSummary string

// BuildDate holds the timestamp when this was built.
var BuildDate string

// Version holds the version from the file VERSION.
var Version string

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information for CrowsNest",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("CrowsNest v%s\n%s on branch %s\nBuilt %s\b", Version, GitSummary, GitBranch, BuildDate)
	},
}
