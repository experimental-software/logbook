package core

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/experimental-software/logbook2/logging"
)

var searchPathPattern = regexp.MustCompile(`.*[/\\](\d{4})[/\\](\d{2})[/\\](\d{2})[/\\](\d{2})\.(\d{2})_.*`)

func Search(baseDirectory, searchTerm string, from time.Time, to time.Time) []LogbookEntry {
	var result = make([]LogbookEntry, 0)
	err := filepath.Walk(baseDirectory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !isLogEntryFile(path) {
				return nil
			}
			pathDatetimeMatch := searchPathPattern.FindStringSubmatch(path)
			logDatetime := fmt.Sprintf("%s-%s-%sT%s:%s",
				pathDatetimeMatch[1],
				pathDatetimeMatch[2],
				pathDatetimeMatch[3],
				pathDatetimeMatch[4],
				pathDatetimeMatch[5],
			)
			if !isInRequestedTimeRange(logDatetime, from, to) {
				return nil
			}

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
			result = append(result, LogbookEntry{
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

func isInRequestedTimeRange(datetime string, from time.Time, to time.Time) bool {
	return datetime >= from.Format(time.RFC3339) && datetime <= to.Format(time.RFC3339)
}

func isLogEntryFile(path string) bool {
	pathParts := strings.Split(path, string(os.PathSeparator))
	lastPathPart := pathParts[len(pathParts)-1]
	if !strings.HasSuffix(lastPathPart, ".md") {
		return false
	}
	parentDirectory := pathParts[len(pathParts)-2]
	if !(regexp.MustCompile(`^\d{2}\.\d{2}_.*`).MatchString(parentDirectory)) {
		return false
	}
	parentDirectorySlug := parentDirectory[6:]
	fileSlug := strings.Replace(lastPathPart, ".md", "", -1)
	if parentDirectorySlug != fileSlug {
		return false
	}
	pathDatetimeMatch := searchPathPattern.FindStringSubmatch(path)
	if len(pathDatetimeMatch) != 6 {
		return false
	}
	return true
}
