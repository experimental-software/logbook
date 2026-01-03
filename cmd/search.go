package cmd

import (
	"os"
	"strings"

	"github.com/aquasecurity/table"
	"github.com/experimental-software/logbook2/core"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:     "search [flags] search_term",
	Aliases: []string{"s"},

	Run: func(cmd *cobra.Command, args []string) {
		searchTerm := ""
		if len(args) > 0 {
			searchTerm = args[0]
		}
		logEntries := core.Search(configuration.LogDirectory, searchTerm)

		t := table.New(os.Stdout)
		t.SetHeaders("Date / Time", "Title", "Path")
		for _, entry := range logEntries {
			title := entry.Title
			if len(title) > 45 {
				title = title[:45]
				title += " (...)"
			}
			t.AddRow(strings.Replace(entry.DateTime, "T", " ", 1), title, entry.Directory)
		}
		t.Render()
	},
}
