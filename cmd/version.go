package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// To be assigned by govvv.

// GitBranch holds the name of the current branch.
var GitBranch string

// GitCommit holds the short commit hash of the most recent commit.
var GitCommit string

// BuildDate holds the timestamp when this was built.
var BuildDate string

// Version holds the version from the file VERSION.
var Version string

// Version/build info, NNL = No New Lines.
var cnMetaString string = fmt.Sprintf("CrowsNest v%s\n%s on %s\nBuilt %s", Version, GitCommit, GitBranch, BuildDate)
var cnMetaStringNNL string = fmt.Sprintf("CrowsNest v%s / %s on %s / Built %s", Version, GitCommit, GitBranch, BuildDate)

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
