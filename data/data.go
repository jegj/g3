package data

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const DEFAULT_DATA_FILENAME = "data.db"

var (
	DEFAULT_DATA_FOLDER   = filepath.Join(os.Getenv("HOME"), ".local", "share", "g3")
	DEFAULT_DATA_FILEPATH = filepath.Join(DEFAULT_DATA_FOLDER, DEFAULT_DATA_FILENAME)
)

func CreateDataFileIfRequired() error {
	if err := os.MkdirAll(DEFAULT_DATA_FOLDER, 0o700); err != nil {
		return err
	}
	dataFilePath := filepath.Join(DEFAULT_DATA_FOLDER, DEFAULT_DATA_FILENAME)
	dataFile, err := os.OpenFile(dataFilePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer dataFile.Close()
	return nil
}

func AppendEntry(filename string, gists []GistEntry) error {
	file, err := os.OpenFile(DEFAULT_DATA_FILEPATH, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	fileGist := FileGist{Gist: gists}
	dataEntry := DataEntry{
		filename: fileGist,
	}

	data, err := json.Marshal(dataEntry)
	if err != nil {
		return err
	}

	_, err = file.Write(append(data, '\n'))
	return err
}
