package core

import (
	"os"
	"testing"
	"time"
)

func TestAddLogEntry(t *testing.T) {
	// Arrange
	tempDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer os.RemoveAll(tempDir)

	// Act
	entry, err := AddLogEntry(tempDir, "This is a new log entry", time.Now())

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	if entry.Title != "This is a new log entry" {
		t.Errorf("AddLogEntry returned wrong title")
	}
	allLogEntries := Search(tempDir, "", epoc, nextCentury)
	if len(allLogEntries) != 1 {
		t.Errorf("AddLogEntry returned wrong number of entries")
	}
}
