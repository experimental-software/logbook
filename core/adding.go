package core

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var epoc, _ = time.Parse("2006-01-02", "1970-01-01")
var nextCentury, _ = time.Parse("2006-01-02", "2100-01-01")

func AddLogEntry(baseDirectory, title string, dateTime time.Time) (LogbookEntry, error) {
	slug := slugify(title)
	formattedDateTime := fmt.Sprintf("%d-0%d-0%dT0%d:0%d",
		dateTime.Year(),
		dateTime.Month(),
		dateTime.Day(),
		dateTime.Hour(),
		dateTime.Minute(),
	)

	logDirectoryPath := filepath.Join(baseDirectory,
		fmt.Sprintf("%d", dateTime.Year()),
		fmt.Sprintf("%02d", dateTime.Month()),
		fmt.Sprintf("%02d", dateTime.Day()),
		fmt.Sprintf("%02d.%02d_%s", dateTime.Hour(), dateTime.Minute(), slug),
	)
	err := os.MkdirAll(logDirectoryPath, 0777)
	if err != nil {
		return LogbookEntry{}, err
	}

	logFilePath := filepath.Join(logDirectoryPath, slug+".md")
	err = os.WriteFile(logFilePath, []byte(fmt.Sprintf("# %s\n\n", title)), 0777)
	if err != nil {
		return LogbookEntry{}, err
	}

	return LogbookEntry{DateTime: formattedDateTime, Title: title, Directory: logDirectoryPath}, nil
}

func slugify(s string) string {
	result := strings.ToLower(s)

	result = regexp.MustCompile(`\|`).ReplaceAllString(result, "_")
	result = regexp.MustCompile(`[^A-Za-z0-9_]`).ReplaceAllString(result, "-")
	result = regexp.MustCompile(`-_-`).ReplaceAllString(result, "_")
	result = regexp.MustCompile(`-+`).ReplaceAllString(result, "-")
	result = regexp.MustCompile(`^-`).ReplaceAllString(result, "")
	result = regexp.MustCompile(`-$`).ReplaceAllString(result, "")

	if len(result) > 35 {
		result = result[:35]
	}

	return result
}
