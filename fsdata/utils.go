package fsdata

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

func GetG3Filepath(filename string) (string, error) {
	if strings.TrimSpace(filename) == "" {
		return "", errors.New("filename cannot be empty or only whitespace")
	}

	invalidChars := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
	for _, char := range invalidChars {
		if strings.Contains(filename, char) {
			return "", fmt.Errorf("filename contains an invalid character: %q", char)
		}
	}

	g3Filename := fmt.Sprintf("%s.g3.json", filename)
	return filepath.Join(DEFAULT_DATA_FILE_FOLDER, g3Filename), nil
}

func GetFileName(absFilePath string) string {
	return filepath.Base(absFilePath)
}
