package cmd

import (
	"errors"
	"sync"

	"github.com/psidex/CrowsNest/internal/config"
	"github.com/psidex/CrowsNest/internal/git"
	"github.com/psidex/CrowsNest/internal/watcher"
	"github.com/spf13/cobra"
)

// Flags.
var (
	runOnce    bool
	configPath string
)

var rootCmd = &cobra.Command{
	Use:   "crowsnest",
	Short: "CrowsNest is Watchtower for Git",
	Long:  `Watchtower for Git`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if !git.BinaryExists() {
			return errors.New("cannot find git binary")
		}

		config, err := config.Get(configPath)
		if err != nil {
			return err
		}

		var wg sync.WaitGroup
		i := 1

		for repoName, repoConfig := range config.Respositories {
			wg.Add(1)
			go watcher.Watch(i, &wg, runOnce, repoName, repoConfig)
			i++
		}

		wg.Wait()
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&runOnce, "run-once", "r", false, "normally CrowsNest would loop forever, set this flag to run once then exit")
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", ".", "where to look for your config.yaml file (. and $HOME are automatically searched)")
}

func Execute() {
	rootCmd.Execute()
}
