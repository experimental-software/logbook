package cmd

import (
	"fmt"
	"os"

	"github.com/experimental-software/logbook2/core"
	"github.com/experimental-software/logbook2/logging"
	"github.com/spf13/cobra"
)

var archiveCmd = &cobra.Command{
	Use:  "archive [flags] path",
	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		err := core.Archive(configuration, path)
		if err != nil {
			logging.Error("Archive failed", err)
			os.Exit(1)
		}
		fmt.Println(path)
	},
}
