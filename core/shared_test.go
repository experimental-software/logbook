package core

import (
	"fmt"
	"os"
	"testing"
)

func Test_logbookEntryRootPath_invalid_path(t *testing.T) {
	_, err := logbookEntryRootPath("/Users/jdoe/Notes/2025/12/21/just-a-test")
	if err == nil {
		t.Fatal("Expected error")
	}
	fmt.Println(err.Error())
}

func createTempDir() string {
	tempDir, err := os.MkdirTemp("", "test")
	if err != nil {
		panic(err)
	}
	return tempDir
}
