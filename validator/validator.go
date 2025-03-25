package validator

import (
	"os"
	"path/filepath"
)

func FileExists(path string) bool {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return false
	}

	info, err := os.Stat(absPath)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil || info != nil
}
