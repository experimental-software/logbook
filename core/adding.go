package core

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func AddLogEntry(baseDirectory, title string) (LogEntry, error) {
	currentTime := time.Now()
	slug := slugify(title)
	dateTime := fmt.Sprintf("%d-0%d-0%dT0%d:0%d",
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Minute(),
	)

	logDirectoryPath := filepath.Join(baseDirectory,
		fmt.Sprintf("%d", currentTime.Year()),
		fmt.Sprintf("%02d", currentTime.Month()),
		fmt.Sprintf("%02d", currentTime.Day()),
		fmt.Sprintf("%02d.%02d_%s", currentTime.Hour(), currentTime.Minute(), slug),
	)
	err := os.MkdirAll(logDirectoryPath, 0777)
	if err != nil {
		return LogEntry{}, err
	}

	logFilePath := filepath.Join(logDirectoryPath, slug+".md")
	err = os.WriteFile(logFilePath, []byte(fmt.Sprintf("# %s\n\n", title)), 0777)
	if err != nil {
		return LogEntry{}, err
	}

	return LogEntry{DateTime: dateTime, Title: title, Directory: logDirectoryPath}, nil
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
