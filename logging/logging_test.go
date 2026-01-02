package logging

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestInfoLog(t *testing.T) {
	// Arrange
	var str bytes.Buffer
	infoLog.SetOutput(&str)

	// Act
	Info("For your information")

	// Assert
	output := strings.TrimSuffix(str.String(), "\n")
	if !(strings.HasSuffix(output, "INFO: For your information")) {
		t.Errorf("Unexpected info log output: %s", str.String())
	}
}

func TestWarningLog(t *testing.T) {
	// Arrange
	var str bytes.Buffer
	warningLog.SetOutput(&str)

	// Act
	Warn("System is in unexpected state")

	// Assert
	output := strings.TrimSuffix(str.String(), "\n")
	if !(strings.HasSuffix(output, "WARNING: System is in unexpected state")) {
		t.Errorf("Unexpected warning log output: %s", str.String())
	}
}

func TestErrorLog(t *testing.T) {
	// Arrange
	var str bytes.Buffer
	errorLog.SetOutput(&str)

	// Act
	Error("something went wrong", errors.New("this is an error"))

	// Assert
	output := strings.TrimSuffix(str.String(), "\n")
	if !(strings.Contains(output, "ERROR: something went wrong!")) || !(strings.Contains(output, "this is an error")) {
		t.Errorf("Unexpected error log output: %s", str.String())
	}
}
