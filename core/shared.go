package core

import (
	"errors"
	"regexp"
	"strings"
)

type LogbookEntry struct {
	DateTime  string `json:"dateTime"`
	Title     string `json:"title"`
	Directory string `json:"directory"`
}

func logbookEntryRootPath(path string) (string, error) {
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	re := regexp.MustCompile(`(.*[/\\]\d{4}[/\\]\d{2}[/\\]\d{2}[/\\]\d{2}\.\d{2}_.*?[/\\]).*`)
	m := re.FindStringSubmatch(path)
	if len(m) != 2 {
		return "", errors.New("invalid logbook entry path: " + path)
	}
	return m[1], nil
}
