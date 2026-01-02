package core

import (
	"fmt"
	"testing"
)

func TestSearchWithoutSearchTerm(t *testing.T) {
	result := Search("./t/2026/01", "")
	if len(result) != 2 {
		t.Errorf("actual: %v, expected: %v", len(result), 2)
	}
}

func TestSearchWithSearchTerm(t *testing.T) {
	result := Search("./t/2026/01", "ANOTHER")
	if len(result) != 1 {
		t.Errorf("actual: %v, expected: %v", len(result), 1)
	}
	fmt.Println(result)
}

func TestIgnoreUnexpectedData(t *testing.T) {
	result := Search("./t/2023", "")
	if len(result) != 0 {
		t.Errorf("actual: %v, expected: %v", len(result), 0)
	}
	fmt.Println(result)
}
