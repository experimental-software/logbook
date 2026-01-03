package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/experimental-software/logbook2/core"
	"github.com/experimental-software/logbook2/logging"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [flags] title",
	Args:  cobra.ExactArgs(1),
	Short: "Add new logbook entries",

	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		result, err := core.AddLogEntry(configuration.LogDirectory, title, time.Now())
		if err != nil {
			logging.Error("Failed to create log entry", err)
			os.Exit(1)
		}
		fmt.Println(result.Directory)
	},
}
