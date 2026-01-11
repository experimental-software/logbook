package core

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func Test_Delete_happy_path(t *testing.T) {
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
	searchResultForArchiveBaseDir := Search(logBaseDir, "", epoc, nextCentury)
	if len(searchResultForArchiveBaseDir) != 1 {
		t.Fatal("Expected 1 search result")
	}

	// Act
	err = Remove(logEntry.Directory)
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	searchResultForLogBaseDir := Search(logBaseDir, "", epoc, nextCentury)
	if len(searchResultForLogBaseDir) != 0 {
		t.Fatal("Expected empty search result")
	}
}

func createFileInSubdirectory(entry LogbookEntry) string {
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
