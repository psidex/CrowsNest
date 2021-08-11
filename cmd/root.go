package cmd

import (
	"errors"
	"fmt"

	"github.com/psidex/CrowsNest/internal/config"
	"github.com/spf13/cobra"
)

// Flags.
var (
	loop       bool
	method     string
	configPath string
)

var rootCmd = &cobra.Command{
	Use:   "crowsnest",
	Short: "CrowsNest is Watchtower for Git",
	Long:  `Watchtower for Git`,
	RunE: func(cmd *cobra.Command, args []string) error {
		config := config.Get(configPath)

		if method != "pull" && method != "checkpull" {
			return errors.New("method flag must be empty, pull, or checkpull")
		}

		fmt.Println(config)

		// Steps:
		//  - For each directory:
		//    - Spawn a goroutine and add to pile
		//    - Git pull (if checkpull, check remote then pull if needed)
		//    - Check if pull changed anything
		//  - Wait for goroutines to finish
		//  - Sleep for interval or exit (depending on loop)

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&loop, "loop", "l", false, "normally CrowsNest would run once then exit, set this flag to loop forever")
	rootCmd.PersistentFlags().StringVarP(&method, "method", "m", "pull", "which method to use for checking / updating, should be pull or checkpull")
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", ".", "where to look for your config.yaml file (. and $HOME are automatically searched)")
}

func Execute() {
	rootCmd.Execute()
}
