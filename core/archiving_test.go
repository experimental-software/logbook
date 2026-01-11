package core

import (
	"os"
	"testing"
	"time"

	"github.com/experimental-software/logbook2/config"
)

func Test_Archive_happy_path(t *testing.T) {
	// Arrange
	logBaseDir := createTempDir()
	archiveBaseDir := createTempDir()
	defer func(path string) {
		_ = os.RemoveAll(logBaseDir)
		_ = os.RemoveAll(archiveBaseDir)
	}(logBaseDir)

	logEntry, err := AddLogEntry(logBaseDir, "Log entry for archive test", time.Now())
	if err != nil {
		t.Fatal(err)
	}

	c := config.Configuration{
		LogDirectory:     logBaseDir,
		ArchiveDirectory: archiveBaseDir,
	}

	// Act
	err = Archive(c, logEntry.Directory)
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	searchResultForLogBaseDir := Search(logBaseDir, "", epoc, nextCentury)
	if len(searchResultForLogBaseDir) != 0 {
		t.Fatal("Expected empty search result")
	}
	searchResultForArchiveBaseDir := Search(archiveBaseDir, "", epoc, nextCentury)
	if len(searchResultForArchiveBaseDir) != 1 {
		t.Fatal("Expected 1 search result")
	}
}

func Test_Archive_path_in_subdirectory(t *testing.T) {
	// Arrange
	logBaseDir := createTempDir()
	archiveBaseDir := createTempDir()
	defer func(path string) {
		_ = os.RemoveAll(logBaseDir)
		_ = os.RemoveAll(archiveBaseDir)
	}(logBaseDir)

	logEntry, err := AddLogEntry(logBaseDir, "Log entry for archive test", time.Now())
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
	searchResultForLogBaseDir := Search(logBaseDir, "", epoc, nextCentury)
	if len(searchResultForLogBaseDir) != 0 {
		t.Fatal("Expected empty search result")
	}
	searchResultForArchiveBaseDir := Search(archiveBaseDir, "", epoc, nextCentury)
	if len(searchResultForArchiveBaseDir) != 1 {
		t.Fatal("Expected 1 search result")
	}
}
