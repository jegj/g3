package data

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

const DEFAULT_DATA_FILENAME = "data.jsonl"

var (
	DEFAULT_DATA_FOLDER   = filepath.Join(os.Getenv("HOME"), ".local", "share", "g3")
	DEFAULT_DATA_FILEPATH = filepath.Join(DEFAULT_DATA_FOLDER, DEFAULT_DATA_FILENAME)
	ErrEntryNotFound      = errors.New("entry not found")
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

func ReadEntries() ([]DataEntry, error) {
	file, err := os.Open(DEFAULT_DATA_FILEPATH)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var entries []DataEntry
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var entry DataEntry
		if err := json.Unmarshal([]byte(scanner.Text()), &entry); err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, scanner.Err()
}

func DeleteEntryByFilename(filename string) error {
	entries, err := ReadEntries()
	if err != nil {
		return err
	}
	file, err := os.Create(DEFAULT_DATA_FILEPATH)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	for _, entry := range entries {
		if _, exist := entry[filename]; !exist {
			data, _ := json.Marshal(entry)
			_, _ = writer.WriteString(string(data) + "\n")
		}
	}
	return nil
}
