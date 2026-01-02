package core

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/experimental-software/logbook2/logging"
)

var logFilePathPattern = regexp.MustCompile(`.*[/\\](\d{4})[/\\](\d{2})[/\\](\d{2})[/\\](\d{2})\.(\d{2})_.*`)
var logfileParentDirectoryPattern = regexp.MustCompile(`^\d{2}\.\d{2}_.*`)

func Search(logDirectory string, searchTerm string) []LogEntry {
	var result = make([]LogEntry, 0)
	err := filepath.Walk(logDirectory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			pathParts := strings.Split(path, string(os.PathSeparator))
			lastPathPart := pathParts[len(pathParts)-1]
			if !strings.HasSuffix(lastPathPart, ".md") {
				return nil
			}
			parentDirectory := pathParts[len(pathParts)-2]
			if !(logfileParentDirectoryPattern.MatchString(parentDirectory)) {
				return nil
			}
			parentDirectorySlug := parentDirectory[6:]
			fileSlug := strings.Replace(lastPathPart, ".md", "", -1)
			if parentDirectorySlug != fileSlug {
				return nil
			}
			pathDatetimeMatch := logFilePathPattern.FindStringSubmatch(path)
			if len(pathDatetimeMatch) != 6 {
				return nil
			}
			logDatetime := fmt.Sprintf("%s-%s-%sT%s:%s",
				pathDatetimeMatch[1],
				pathDatetimeMatch[2],
				pathDatetimeMatch[3],
				pathDatetimeMatch[4],
				pathDatetimeMatch[5],
			)

			logFileBytes, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			logFile := string(logFileBytes)
			title := strings.Replace(strings.Split(logFile, "\n")[0], "# ", "", 1)
			if searchTerm != "" && !strings.Contains(strings.ToLower(title), strings.ToLower(searchTerm)) {
				return nil
			}

			logDirectory, _ := filepath.Abs(filepath.Dir(path))
			result = append(result, LogEntry{
				DateTime:  logDatetime,
				Directory: logDirectory,
				Title:     title,
			})

			return nil
		})
	if err != nil {
		logging.Error("Could not traverse log directory.", err)
	}
	return result
}
