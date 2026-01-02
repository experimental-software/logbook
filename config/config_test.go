package config

import "testing"

func TestLoadConfiguration(t *testing.T) {
	result := loadConfiguration("./t/config.yaml")

	if result.LogDirectory != "/path/for/tests" {
		t.Errorf("Unexpected log directory: got %s, expected /path/for/tests", result.LogDirectory)
	}
}
