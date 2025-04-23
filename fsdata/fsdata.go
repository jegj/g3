package fsdata

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/jegj/g3/config"
	"github.com/jegj/g3/g3unit"
)

var ErrEntryNotFound = errors.New("entry not found")

type DataProvider interface {
	AppendEntry(unit g3unit.G3Unit, gists []GistEntry) error
	UpdateEntry(unit g3unit.G3Unit, gists []GistEntry) error
	DeleteEntry(unit g3unit.G3Unit) error
	GetEntries() ([]string, error)
	GetFileSize(unit g3unit.G3Unit) (int64, error)
	GetFileContent(unit g3unit.G3Unit) ([]byte, error)
	GetEntry(unit g3unit.G3Unit) (DataEntry, error)
	HasEntry(unit g3unit.G3Unit) bool
}

type FSDataService struct {
	cfg config.Config
}

func NewDatatService(cfg config.Config) FSDataService {
	return FSDataService{
		cfg: cfg,
	}
}

func (d FSDataService) GetFileSize(unit g3unit.G3Unit) (int64, error) {
	info, err := os.Stat(unit.Filepath)
	if err != nil {
		return 0, err
	} else {
		return info.Size(), nil
	}
}

func (d FSDataService) GetFileContent(unit g3unit.G3Unit) ([]byte, error) {
	return os.ReadFile(unit.Filepath)
}

func (d FSDataService) AppendEntry(unit g3unit.G3Unit, gists []GistEntry) (err error) {
	file, err := os.OpenFile(unit.G3Filepath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	dataEntry := DataEntry{
		Gist:      gists,
		CreatedAt: time.Now(),
	}

	data, err := json.Marshal(dataEntry)
	if err != nil {
		return err
	}

	_, err = file.Write(append(data, '\n'))
	return err
}

func (d FSDataService) DeleteEntry(unit g3unit.G3Unit) error {
	err := os.Remove(unit.G3Filepath)
	if err != nil {
		return err
	}
	return nil
}

func (d FSDataService) UpdateEntry(unit g3unit.G3Unit, gists []GistEntry) (err error) {
	file, err := os.OpenFile(unit.G3Filepath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	dataEntry := DataEntry{
		Gist:      gists,
		CreatedAt: time.Now(),
	}

	data, err := json.Marshal(dataEntry)
	if err != nil {
		return err
	}

	_, err = file.Write(append(data, '\n'))
	return err
}

// FIXME: Only valid file. e.g .setting.swp or hidden files
func (d FSDataService) GetEntries() ([]string, error) {
	files, err := os.ReadDir(d.cfg.DataFolder)
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

func (d FSDataService) GetEntry(unit g3unit.G3Unit) (DataEntry, error) {
	data, err := os.ReadFile(unit.G3Filepath)
	if err != nil {
		return DataEntry{}, err
	}
	dataEntry := DataEntry{}
	err = json.Unmarshal(data, &dataEntry)
	if err != nil {
		return DataEntry{}, err
	}

	return dataEntry, nil
}

func (d FSDataService) HasEntry(unit g3unit.G3Unit) bool {
	_, err := os.Stat(unit.G3Filepath)
	return !os.IsNotExist(err)
}
