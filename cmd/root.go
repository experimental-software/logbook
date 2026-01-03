package cmd

import (
	"os"

	"github.com/experimental-software/logbook2/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "logbook2",
	Short: "A markdown-based engineering logbook",

	Run: func(cmd *cobra.Command, args []string) {
		// noop
	},
}

func Execute() {
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(archiveCmd)
	rootCmd.AddCommand(removeCmd)

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var configuration = config.LoadConfiguration(config.DefaultConfigurationFilePath)
