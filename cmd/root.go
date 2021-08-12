package cmd

import (
	"errors"
	"sync"

	"github.com/psidex/CrowsNest/internal/config"
	"github.com/psidex/CrowsNest/internal/git"
	"github.com/psidex/CrowsNest/internal/watcher"
	"github.com/spf13/cobra"
)

var cnFlags config.Flags

var rootCmd = &cobra.Command{
	Use:   "crowsnest",
	Short: "CrowsNest is Watchtower for Git",
	Long:  `Watchtower for Git`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if !git.BinaryExists() {
			return errors.New("cannot find git binary")
		}

		config, err := config.Get(cnFlags.ConfigPath)
		if err != nil {
			return err
		}

		var wg sync.WaitGroup
		i := 1

		for repoName, repoConfig := range config.Repositories {
			wg.Add(1)
			go watcher.Watch(i, &wg, cnFlags, repoName, repoConfig)
			i++
		}

		wg.Wait()
		return nil
	},
}

func init() {
	// TODO: What options does watchtower have that we might need?
	// TODO: Maybe allow notification of new pull thru messenger services or something.
	rootCmd.PersistentFlags().BoolVarP(&cnFlags.RunOnce, "run-once", "r", false, "normally CrowsNest would loop forever, set this flag to run once then exit")
	rootCmd.PersistentFlags().StringVarP(&cnFlags.ConfigPath, "config", "c", ".", "where to look for your config.yaml file (. and $HOME are automatically searched)")
	rootCmd.PersistentFlags().BoolVar(&cnFlags.Verbose, "verbose", false, "write a lot more info to the log, useful for finding errors")
}

func Execute() {
	rootCmd.Execute()
}
