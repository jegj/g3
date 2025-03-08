package fs

import (
	"os"
	"path/filepath"
)

func GetFileSize(absFilePath string) (int64, error) {
	info, err := os.Stat(absFilePath)
	if err != nil {
		return 0, err
	} else {
		return info.Size(), nil
	}
}

func GetFileName(absFilePath string) string {
	return filepath.Base(absFilePath)
}

func GetFileContent(absFilePath string) ([]byte, error) {
	return os.ReadFile(absFilePath)
}
