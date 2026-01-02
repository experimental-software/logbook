package config

import "testing"

func TestLoadConfiguration(t *testing.T) {
	result := LoadConfiguration("./t/config.yaml")

	if result.LogDirectory != "/path/for/tests" {
		t.Errorf("Unexpected log directory: got %s, expected /path/for/tests", result.LogDirectory)
	}
	if result.ArchiveDirectory != "/archive/path/for/tests" {
		t.Errorf("Unexpected archive directory: got %s, expected /archive/path/for/tests", result.ArchiveDirectory)
	}
}
