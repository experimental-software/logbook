package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aquasecurity/table"
	"github.com/experimental-software/logbook2/core"
	"github.com/experimental-software/logbook2/utils"
	"github.com/spf13/cobra"
)

var FromParameter string
var ToParameter string

var searchCmd = &cobra.Command{
	Use:     "search [flags] search_term",
	Aliases: []string{"s"},
	Short:   "Search for logbook entries",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		_, err := time.Parse(utils.RFC3339date, FromParameter)
		if err != nil {
			return fmt.Errorf("unexpected format for --from flag")
		}
		_, err = time.Parse(utils.RFC3339date, ToParameter)
		if err != nil {
			return fmt.Errorf("unexpected format for --to flag")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		searchTerm := ""
		if len(args) > 0 {
			searchTerm = args[0]
		}

		from, _ := time.Parse(utils.RFC3339date, FromParameter)
		to, _ := time.Parse(utils.RFC3339date, ToParameter)

		logEntries := core.Search(
			configuration.LogDirectory, searchTerm, from, to,
		)

		outputFormat, err := cmd.Flags().GetString("output-format")
		if err != nil {
			panic(err)
		}

		switch outputFormat {
		default:
			{
				t := table.New(os.Stdout)
				t.SetHeaders("Date / Time", "Title", "Path")
				for _, entry := range logEntries {
					title := entry.Title
					if len(title) > 50 {
						title = title[:50]
						title += " (...)"
					}
					t.AddRow(strings.Replace(entry.DateTime, "T", " ", 1), title, entry.Directory)
				}
				t.Render()
			}
		case "list":
			{
				for _, entry := range logEntries {
					fmt.Println(entry.Directory)
				}
			}
		case "json":
			b, err := json.MarshalIndent(logEntries, "", "  ")
			if err != nil {
				panic(err)
			}
			fmt.Println(string(b))
		}
	},
}

func init() {
	flags := searchCmd.Flags()
	flags.StringVarP(&FromParameter, "from", "f", "1970-01-01", "RFC 3339 formatted date for the earliest considered logbook entry.")
	flags.StringVarP(&ToParameter, "to", "t", "2100-01-01", "RFC 3339 formatted date for the latest considered logbook entry.")

	flags.VarP(StringChoice([]string{
		"table", "list", "json",
	}), "output-format", "o", "The format in which the log entries are printed to the terminal.\n[table (default), list, json]")
}
