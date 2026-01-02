package core

import (
	"fmt"
	"testing"
)

func TestSearchWithoutSearchTerm(t *testing.T) {
	result := Search("./t/2026/01", "")
	if len(result) != 2 {
		t.Errorf("Expected two search results but found %v", len(result))
	}
}

func TestSearchWithSearchTerm(t *testing.T) {
	result := Search("./t/2026/01", "ANOTHER")
	if len(result) != 1 {
		t.Errorf("Expected only one search result but found %v", len(result))
	}
	fmt.Println(result)
}

func TestIgnoreUnexpectedData(t *testing.T) {
	result := Search("./t/2023", "")
	if len(result) != 0 {
		t.Errorf("Expected no search result but found %v", len(result))
	}
	fmt.Println(result)
}
