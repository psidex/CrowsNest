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

// Version/build info, NNL = No New Lines.
var cnMetaString string = fmt.Sprintf("CrowsNest v%s\n%s on branch %s\nBuilt %s", Version, GitSummary, GitBranch, BuildDate)
var cnMetaStringNNL string = fmt.Sprintf("CrowsNest v%s / %s on branch %s / Built %s", Version, GitSummary, GitBranch, BuildDate)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information for CrowsNest",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cnMetaString)
	},
}
