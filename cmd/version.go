package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// To be assigned by govvv.
var GitBranch, GitSummary, BuildDate, Version string

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
