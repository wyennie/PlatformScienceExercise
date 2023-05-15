package io

import (
	"os"
	"testing"
)

func TestReadLines(t *testing.T) {
	filename := "test.txt"
	expectedLines := []string{"line 1", "line 2", "line 3"}

	// Create a temporary file with some sample data to read
	file, err := os.Create(filename)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(filename)

	for _, line := range expectedLines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			t.Fatalf("Failed to write to test file: %v", err)
		}
	}
	file.Close()

	// Call readLines and verify that the output matches the expected lines
	actualLines, err := readLines(filename)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	if len(actualLines) != len(expectedLines) {
		t.Errorf("Unexpected number of lines. Expected %d, got %d", len(expectedLines), len(actualLines))
	}

	for i, expectedLine := range expectedLines {
		if actualLines[i] != expectedLine {
			t.Errorf("Unexpected line at index %d. Expected '%s', got '%s'", i, expectedLine, actualLines[i])
		}
	}
}
