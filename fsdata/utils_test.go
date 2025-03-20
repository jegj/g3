package fsdata

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetG3Filepath tests GetG3Filepath using a table-driven approach
func TestGetG3Filepath(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name     string // Test case name
		input    string // Input filename
		expected string // Expected output path
	}{
		{
			name:     "Normal filename",
			input:    "example",
			expected: filepath.Join(DEFAULT_DATA_FILE_FOLDER, "example.g3.json"),
		},
		{
			name:     "Filename with extension",
			input:    "document.txt",
			expected: filepath.Join(DEFAULT_DATA_FILE_FOLDER, "document.txt.g3.json"),
		},
	}

	// Iterate through test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := GetG3Filepath(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
