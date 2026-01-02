package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/experimental-software/logbook2/logging"
	"sigs.k8s.io/yaml"
)

var defaultConfigurationFilePath = filepath.Join(userHomeDir(), ".config", "logbook", "config.yaml")

func LogDirectory() string {
	logDirectory := loadConfiguration(defaultConfigurationFilePath).LogDirectory
	return strings.Replace(logDirectory, "~", userHomeDir(), -1)
}

type configuration struct {
	LogDirectory string `json:"logDirectory"`
}

var defaultConfiguration = configuration{
	LogDirectory: filepath.Join(userHomeDir(), "Logs"),
}

func loadConfiguration(configurationFilePath string) configuration {
	configurationBytes, err := os.ReadFile(configurationFilePath)
	if err != nil {
		logging.Warn("Failed to read configuration file: " + configurationFilePath)
		return defaultConfiguration
	}
	var result configuration
	err = yaml.Unmarshal(configurationBytes, &result)
	if err != nil {
		logging.Warn("Failed to unmarshal configuration file: " + configurationFilePath)
		return defaultConfiguration
	}
	return result
}

func userHomeDir() string {
	result, err := os.UserHomeDir()
	if err != nil {
		panic("Cannot get user home directory")
	}
	return result
}
