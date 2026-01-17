package core

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/experimental-software/logbook2/logging"
)

// e.g. /path/to/2026/01/11/18.03_wip
var searchPathPattern = regexp.MustCompile(`.*[/\\](\d{4})[/\\](\d{2})[/\\](\d{2})[/\\](\d{2})\.(\d{2})_.*`)

// e.g. 20250111T1810
var isoDateTimeBasicFormat = regexp.MustCompile(`^\d{4}\d{2}\d{2}T\d{2}\d{2}$`)

func Search(baseDirectory, searchTerm string, from time.Time, to time.Time) []LogbookEntry {
	var result = make([]LogbookEntry, 0)

	maxDepth := strings.Count(baseDirectory, string(os.PathSeparator)) + 4

	err := filepath.WalkDir(baseDirectory,
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() && strings.Count(path, string(os.PathSeparator)) > maxDepth {
				return fs.SkipDir
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

			logEntryFileBytes, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			logFile := string(logEntryFileBytes)
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
	if !strings.HasSuffix(path, ".md") {
		return false
	}

	pathParts := strings.Split(path, string(os.PathSeparator))
	lastPathPart := pathParts[len(pathParts)-1]
	parentDirectory := pathParts[len(pathParts)-2]
	if !(regexp.MustCompile(`^\d{2}\.\d{2}_.*`).MatchString(parentDirectory)) {
		return false
	}
	parentDirectorySlug := parentDirectory[6:]
	simpleFileName := strings.Replace(lastPathPart, ".md", "", -1)
	if len(isoDateTimeBasicFormat.FindStringSubmatch(simpleFileName)) != 1 && parentDirectorySlug != simpleFileName {
		return false
	}
	pathDatetimeMatch := searchPathPattern.FindStringSubmatch(path)
	if len(pathDatetimeMatch) == 6 {
		return true
	}

	return false
}
