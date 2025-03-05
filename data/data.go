package data

import (
	"os"
	"path/filepath"
)

const DEFAULT_DATA_FILENAME = "data.json"

func CreateDataFileIfRequired() error {
	dataDir := filepath.Join(os.Getenv("HOME"), ".local", "share", "g3")
	if err := os.MkdirAll(dataDir, 0o700); err != nil {
		return err
	}
	dataFilePath := filepath.Join(dataDir, DEFAULT_DATA_FILENAME)
	dataFile, err := os.OpenFile(dataFilePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer dataFile.Close()
	return nil
}
