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
	LogDirectory: filepath.Join(userHomeDir(), "Logs"),
}

func LoadConfiguration(configurationFilePath string) Configuration {
	var result Configuration

	configurationBytes, err := os.ReadFile(configurationFilePath)
	if err != nil {
		logging.Warn("Failed to read Configuration file: " + configurationFilePath)
		result = defaultConfiguration
	}

	err = yaml.Unmarshal(configurationBytes, &result)
	if err != nil {
		logging.Warn("Failed to unmarshal Configuration file: " + configurationFilePath)
		result = defaultConfiguration
	}

	result.LogDirectory = strings.Replace(result.LogDirectory, "~", userHomeDir(), -1)
	result.LogDirectory = strings.TrimSpace(result.LogDirectory)

	result.ArchiveDirectory = strings.Replace(result.ArchiveDirectory, "~", userHomeDir(), -1)
	result.ArchiveDirectory = strings.TrimSpace(result.ArchiveDirectory)

	return result
}

func userHomeDir() string {
	result, err := os.UserHomeDir()
	if err != nil {
		panic("Cannot get user home directory")
	}
	return result
}
