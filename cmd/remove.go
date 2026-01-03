package cmd

import (
	"fmt"
	"os"

	"github.com/experimental-software/logbook2/core"
	"github.com/experimental-software/logbook2/logging"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove [flags] path [path...]",
	Short:   "Deletes log entries",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"rm"},

	Run: func(cmd *cobra.Command, args []string) {
		for _, path := range args {
			err := core.Remove(path)
			if err != nil {
				logging.Error("Removing logbook entry filed failed for path: "+path, err)
				os.Exit(1)
			}
			fmt.Println(path)
		}
	},
}
