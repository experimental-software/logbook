package core

import (
	"errors"
	"os"
	"regexp"
	"strings"

	"github.com/experimental-software/logbook2/config"
	"github.com/plus3it/gorecurcopy"
)

var archivePathPattern = regexp.MustCompile(`(.*[/\\]\d{4}[/\\]\d{2}[/\\]\d{2}[/\\]\d{2}\.\d{2}_.*?[/\\]).*`)

func Archive(configuration config.Configuration, sourcePath string) error {
	if !strings.HasSuffix(sourcePath, "/") {
		sourcePath += "/"
	}
	m := archivePathPattern.FindStringSubmatch(sourcePath)
	if len(m) != 2 {
		return errors.New("invalid source path: " + sourcePath)
	}
	sourceDirectoryPath := strings.TrimSpace(m[1])
	targetDirectoryPath := strings.Replace(sourceDirectoryPath, configuration.LogDirectory, configuration.ArchiveDirectory, 1)

	err := os.MkdirAll(targetDirectoryPath, 0777)
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
