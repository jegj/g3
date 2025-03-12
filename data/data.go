package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const (
	DEFAULT_DATA_FILE_FOLDERNAME = "files"
)

var (
	DEFAULT_DATA_FOLDER      = filepath.Join(os.Getenv("HOME"), ".local", "share", "g3")
	DEFAULT_DATA_FILE_FOLDER = filepath.Join(DEFAULT_DATA_FOLDER, DEFAULT_DATA_FILE_FOLDERNAME)
	ErrEntryNotFound         = errors.New("entry not found")
)

func CreateDataFolderIfRequired() error {
	if err := os.MkdirAll(DEFAULT_DATA_FILE_FOLDER, 0o700); err != nil {
		return err
	}
	return nil
}

func AppendEntry(filename string, gists []GistEntry) error {
	g3Filename := fmt.Sprintf("%s.g3.json", filename)
	g3FilePath := filepath.Join(DEFAULT_DATA_FILE_FOLDER, g3Filename)
	file, err := os.OpenFile(g3FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
