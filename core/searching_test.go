package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/experimental-software/logbook2/utils"
)

func Test_Search_without_search_term(t *testing.T) {
	result := Search("./t/2026/01", "", epoc, nextCentury)
	if len(result) != 2 {
		t.Errorf("Expected two search results but found %v", len(result))
	}
}

func Test_Search_with_search_term(t *testing.T) {
	result := Search("./t/2026/01", "ANOTHER", epoc, nextCentury)
	if len(result) != 1 {
		t.Errorf("Expected only one search result but found %v", len(result))
	}
	fmt.Println(result)
}

func Test_Search_ignore_unexpected_data(t *testing.T) {
	result := Search("./t/2023", "", epoc, nextCentury)
	if len(result) != 0 {
		t.Errorf("Expected no search result but found %v", len(result))
	}
	fmt.Println(result)
}

func Test_isInRequestedTimeRange(t *testing.T) {

	testCases := []struct {
		name     string
		dateTime string
		from     string
		to       string
		expected bool
	}{
		{
			name:     "happy path",
			dateTime: "2025-12-12T20:25",
			from:     "1970-01-01",
			to:       "2100-01-01",
			expected: true,
		},
		{
			name:     "ignore too early",
			dateTime: "2025-12-12T20:25",
			from:     "2025-12-13",
			to:       "2100-01-01",
			expected: false,
		},
		{
			name:     "ignore too late",
			dateTime: "2025-12-12T20:25",
			from:     "1970-01-01",
			to:       "2025-12-11",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			from, _ := time.Parse(utils.RFC3339date, tc.from)
			to, _ := time.Parse(utils.RFC3339date, tc.to)
			result := isInRequestedTimeRange(tc.dateTime, from, to)
			if result != tc.expected {
				t.Errorf("expected: %v, got: %v", tc.expected, result)
			}
		})
	}
}
