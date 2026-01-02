package cmd

import (
	"fmt"
	"os"

	"github.com/experimental-software/logbook2/config"
	"github.com/experimental-software/logbook2/core"
	"github.com/experimental-software/logbook2/logging"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:  "add [flags] title",
	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		result, err := core.AddLogEntry(config.LogDirectory(), title)
		if err != nil {
			logging.Error("Failed to create log entry", err)
			os.Exit(1)
		}
		fmt.Println(result.Directory)
	},
}
