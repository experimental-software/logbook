package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/experimental-software/logbook2/logging"
	"sigs.k8s.io/yaml"
)

var DefaultConfigurationFilePath = filepath.Join(userHomeDir(), ".config", "logbook", "config.yaml")

type Configuration struct {
	LogDirectory     string `json:"logDirectory"`
	ArchiveDirectory string `json:"archiveDirectory"`
}

var defaultConfiguration = Configuration{
	LogDirectory:     filepath.Join(userHomeDir(), "Logs"),
	ArchiveDirectory: filepath.Join(userHomeDir(), "Archive"),
}

func LoadConfiguration(configurationFilePath string) Configuration {
	result := defaultConfiguration

	configurationBytes, err := os.ReadFile(configurationFilePath)
	if err != nil {
		logging.Warn("Failed to read Configuration file: " + configurationFilePath)
		return result
	}

	err = yaml.Unmarshal(configurationBytes, &result)
	if err != nil {
		logging.Warn("Failed to unmarshal Configuration file: " + configurationFilePath)
		return result
	}

	result.LogDirectory = strings.ReplaceAll(result.LogDirectory, "~", userHomeDir())
	result.ArchiveDirectory = strings.ReplaceAll(result.ArchiveDirectory, "~", userHomeDir())

	return result
}

func userHomeDir() string {
	result, err := os.UserHomeDir()
	if err != nil {
		panic("Cannot get user home directory")
	}
	return result
}
