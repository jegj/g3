package fsdata

import (
	"fmt"
	"path/filepath"
)

func GetG3Filepath(filename string) string {
	g3Filename := fmt.Sprintf("%s.g3.json", filename)
	return filepath.Join(DEFAULT_DATA_FILE_FOLDER, g3Filename)
}
