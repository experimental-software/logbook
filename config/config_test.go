package config

import (
	"strings"
	"testing"
)

func Test_LoadConfiguration_happy_path(t *testing.T) {
	result := LoadConfiguration("./t/config.yaml")

	if result.LogDirectory != "/path/for/tests" {
		t.Errorf("Unexpected log directory: got %s, expected /path/for/tests", result.LogDirectory)
	}
	if result.ArchiveDirectory != "/archive/path/for/tests" {
		t.Errorf("Unexpected archive directory: got %s, expected /archive/path/for/tests", result.ArchiveDirectory)
	}
}

func Test_LoadConfiguration_config_file_not_existent(t *testing.T) {
	result := LoadConfiguration("./t/67b2070e-8b21-44d1-8a02-21909d8037f9.yaml")

	if !strings.HasSuffix(result.LogDirectory, "Logs") {
		t.Errorf("Unexpected log directory: got '%s', expected default value", result.LogDirectory)
	}
}

func Test_LoadConfiguration_config_file_invalid(t *testing.T) {
	result := LoadConfiguration("./t/invalid_syntax_config.yaml")

	if !strings.HasSuffix(result.LogDirectory, "Logs") {
		t.Errorf("Unexpected log directory: got '%s', expected default value", result.LogDirectory)
	}
}
