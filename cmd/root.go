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
	Long:  "Watchtower for Git: automatically keep local Git repositories up to date with their remotes",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !git.BinaryExists() {
			return errors.New("cannot find git binary")
		}

		if err := config.ValidateFlags(cnFlags); err != nil {
			return err
		}

		if f, err := config.SetupLog(cnFlags); err != nil {
			return err
		} else if f != nil {
			defer f.Close()
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
	rootCmd.PersistentFlags().StringVarP(&cnFlags.ConfigPath, "config", "c", "", "where to look for your config.yaml file (. and $HOME are automatically searched)")
	rootCmd.PersistentFlags().BoolVarP(&cnFlags.Verbose, "verbose", "v", false, "write a lot more info to the log, useful for finding errors")
	rootCmd.PersistentFlags().StringVarP(&cnFlags.LogPath, "logpath", "l", "", "write the log to the given file instead of stdout. Should be a full path ending with the file name")
}

// Execute executes our application.
func Execute() {
	rootCmd.Execute()
}
