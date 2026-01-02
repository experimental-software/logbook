package core

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/experimental-software/logbook2/config"
)

func TestArchiveLogEntry(t *testing.T) {
	// Arrange
	logBaseDir := createTempDir()
	archiveBaseDir := createTempDir()
	defer func(path string) {
		_ = os.RemoveAll(logBaseDir)
		_ = os.RemoveAll(archiveBaseDir)
	}(logBaseDir)

	logEntry, err := AddLogEntry(logBaseDir, "Log entry for archive test")
	if err != nil {
		t.Fatal(err)
	}
	fileInLog := createFileInSubdirectory(logEntry)

	c := config.Configuration{
		LogDirectory:     logBaseDir,
		ArchiveDirectory: archiveBaseDir,
	}

	// Act
	err = Archive(c, fileInLog)
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	searchResultForLogBaseDir := Search(logBaseDir, "")
	if len(searchResultForLogBaseDir) != 0 {
		t.Fatal("Expected empty search result")
	}
	searchResultForArchiveBaseDir := Search(archiveBaseDir, "")
	if len(searchResultForArchiveBaseDir) != 1 {
		t.Fatal("Expected 1 search result")
	}
}

func createFileInSubdirectory(entry LogEntry) string {
	subdirectory := filepath.Join(entry.Directory, "foo")
	err := os.MkdirAll(subdirectory, 0755)
	if err != nil {
		panic(err)
	}
	filePath := filepath.Join(subdirectory, "bar.txt")
	data := []byte("Hello, World!")
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		panic(err)
	}
	return filePath
}
