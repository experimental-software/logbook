package core

import "os"

func createTempDir() string {
	tempDir, err := os.MkdirTemp("", "test")
	if err != nil {
		panic(err)
	}
	return tempDir
}
