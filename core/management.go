package core

import (
	"errors"
	"os"
	"regexp"
	"strings"

	"github.com/experimental-software/logbook2/config"
	"github.com/plus3it/gorecurcopy"
)

func Archive(configuration config.Configuration, sourcePath string) error {
	sourceDirectoryPath, err := logbookEntryRootPath(sourcePath)
	if err != nil {
		return err
	}
	targetDirectoryPath := strings.Replace(
		sourceDirectoryPath, configuration.LogDirectory, configuration.ArchiveDirectory, 1,
	)

	err = os.MkdirAll(targetDirectoryPath, 0777)
	if err != nil {
		return err
	}

	err = gorecurcopy.CopyDirectory(sourceDirectoryPath, targetDirectoryPath)
	if err != nil {
		return err
	}
	err = os.RemoveAll(sourceDirectoryPath)
	return err
}

func Remove(sourcePath string) error {
	sourceDirectoryPath, err := logbookEntryRootPath(sourcePath)
	if err != nil {
		return err
	}
	err = os.RemoveAll(sourceDirectoryPath)
	return err
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
