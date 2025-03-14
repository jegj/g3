package fs

import (
	"os"
	"path/filepath"
)

type FsProvider interface {
	GetFileSize(absFilePath string) (int64, error)
	GetFileName(absFilePath string) string
	GetFileContent(absFilePath string) ([]byte, error)
}

type FsService struct{}

func NewFsService() FsService {
	return FsService{}
}

func (f FsService) GetFileSize(absFilePath string) (int64, error) {
	info, err := os.Stat(absFilePath)
	if err != nil {
		return 0, err
	} else {
		return info.Size(), nil
	}
}

func (f FsService) GetFileName(absFilePath string) string {
	return filepath.Base(absFilePath)
}

func (f FsService) GetFileContent(absFilePath string) ([]byte, error) {
	return os.ReadFile(absFilePath)
}
