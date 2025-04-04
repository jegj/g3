package fsdata

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/jegj/g3/config"
)

const (
	DEFAULT_DATA_FILE_FOLDERNAME = "files"
)

var (
	DEFAULT_DATA_FOLDER      = filepath.Join(os.Getenv("HOME"), ".local", "share", "g3")
	DEFAULT_DATA_FILE_FOLDER = filepath.Join(DEFAULT_DATA_FOLDER, DEFAULT_DATA_FILE_FOLDERNAME)
	ErrEntryNotFound         = errors.New("entry not found")
)

type DataProvider interface {
	AppendEntry(filename string, gists []GistEntry) error
	DeleteEntry(filename string) error
	GetEntries() ([]string, error)
	GetFileSize(absFilePath string) (int64, error)
	GetFileContent(absFilePath string) ([]byte, error)
	GetEntry(filename string) (DataEntry, error)
	HasEntry(filename string) bool
}

type FSDataService struct {
	cfg config.Config
}

func NewDatatService(cfg config.Config) FSDataService {
	return FSDataService{
		cfg: cfg,
	}
}

func (d FSDataService) GetFileSize(absFilePath string) (int64, error) {
	info, err := os.Stat(absFilePath)
	if err != nil {
		return 0, err
	} else {
		return info.Size(), nil
	}
}

func (d FSDataService) GetFileContent(absFilePath string) ([]byte, error) {
	return os.ReadFile(absFilePath)
}

func (d FSDataService) AppendEntry(filename string, gists []GistEntry) (err error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

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

// FIXME: Only valid file. e.g .setting.swp or hidden files
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

func (d FSDataService) GetEntry(g3filepath string) (DataEntry, error) {
	data, err := os.ReadFile(g3filepath)
	if err != nil {
		return nil, err
	}
	dataEntry := DataEntry{}
	err = json.Unmarshal(data, &dataEntry)
	if err != nil {
		return nil, err
	}

	return dataEntry, nil
}

func (d FSDataService) HasEntry(g3filepath string) bool {
	_, err := os.Stat(g3filepath)
	return !os.IsNotExist(err)
}
