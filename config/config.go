package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/experimental-software/logbook2/logging"
	"sigs.k8s.io/yaml"
)

func LogDirectory() string {
	logDirectory := loadConfiguration().LogDirectory
	return strings.Replace(logDirectory, "~", userHomeDir(), -1)
}

type configuration struct {
	LogDirectory string `json:"logDirectory"`
}

var defaultConfiguration = configuration{
	LogDirectory: filepath.Join(userHomeDir(), "Logs"),
}

func loadConfiguration() configuration {
	configurationPath := filepath.Join(userHomeDir(), ".config", "logbook", "config.yaml")
	configurationBytes, err := os.ReadFile(configurationPath)
	if err != nil {
		logging.Warn("Failed to read configuration file: " + configurationPath)
		return defaultConfiguration
	}
	var result configuration
	err = yaml.Unmarshal(configurationBytes, &result)
	if err != nil {
		logging.Warn("Failed to unmarshal configuration file: " + configurationPath)
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
