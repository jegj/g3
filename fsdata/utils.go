package fsdata

import (
	"path/filepath"
)

func GetFileName(absFilePath string) string {
	return filepath.Base(absFilePath)
}
