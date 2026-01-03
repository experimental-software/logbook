package cmd

import (
	"fmt"
	"os"

	"github.com/experimental-software/logbook2/core"
	"github.com/experimental-software/logbook2/logging"
	"github.com/spf13/cobra"
)

var archiveCmd = &cobra.Command{
	Use:     "archive [flags] path [path...]",
	Short:   "Moves log entries into the archive directory",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"a"},

	Run: func(cmd *cobra.Command, args []string) {
		for _, path := range args {
			err := core.Archive(configuration, path)
			if err != nil {
				logging.Error("Archive failed", err)
				os.Exit(1)
			}
			fmt.Println(path)
		}
	},
}
