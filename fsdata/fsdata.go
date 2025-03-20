package fsdata

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
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

type DataProvider interface {
	AppendEntry(filename string, gists []GistEntry) error
	DeleteEntry(filename string) error
	GetEntries() ([]string, error)
	GetFileSize(absFilePath string) (int64, error)
	GetFileName(absFilePath string) string
	GetFileContent(absFilePath string) ([]byte, error)
}

type FSDataService struct{}

func NewDatatService() FSDataService {
	return FSDataService{}
}

func (d FSDataService) GetFileSize(absFilePath string) (int64, error) {
	info, err := os.Stat(absFilePath)
	if err != nil {
		return 0, err
	} else {
		return info.Size(), nil
	}
}

func (d FSDataService) GetFileName(absFilePath string) string {
	return filepath.Base(absFilePath)
}

func (d FSDataService) GetFileContent(absFilePath string) ([]byte, error) {
	return os.ReadFile(absFilePath)
}

func (d FSDataService) AppendEntry(filename string, gists []GistEntry) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

func (d FSDataService) DeleteEntry(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return err
	}
	return nil
}

func (d FSDataService) GetEntries() ([]string, error) {
	files, err := os.ReadDir(DEFAULT_DATA_FILE_FOLDER)
	if err != nil {
		return []string{}, err
	}

	if len(files) == 0 {
		return []string{}, nil
	}

	filenames := make([]string, 0, len(files))
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filenames = append(filenames, strings.ReplaceAll(file.Name(), ".g3.json", ""))
	}
	return filenames, nil
}
