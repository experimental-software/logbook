package core

import (
	"os"
)

func Remove(sourcePath string) error {
	sourceDirectoryPath, err := logbookEntryRootPath(sourcePath)
	if err != nil {
		return err
	}
	err = os.RemoveAll(sourceDirectoryPath)
	return err
}
