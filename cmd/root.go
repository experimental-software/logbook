package cmd

import (
	"fmt"
	"os"

	"github.com/experimental-software/logbook2/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "logbook2",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("A markdown-based engineering logbook")
	},
}

func Execute() {
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(archiveCmd)

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var configuration = config.LoadConfiguration(config.DefaultConfigurationFilePath)
