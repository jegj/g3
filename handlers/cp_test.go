package handlers

import (
	"errors"
	"testing"

	"github.com/jegj/g3/fsdata"
	"github.com/jegj/g3/github"
	"github.com/stretchr/testify/assert"
)

// TODO: Add config keys
func TestCp_ErrorGetFileSize(t *testing.T) {
	mockGistProvider := new(github.MockGistProvider)
	mockDataProvider := new(fsdata.MockDataProvider)
	g3handler := G3BaseHandler{
		G: mockGistProvider,
		D: mockDataProvider,
	}
	mockDataProvider.On("GetFileSize", "/tmp/backup_25_03_2025.tar.gz").Return(int64(0), errors.New("file not found"))

	err := g3handler.Cp("/tmp/backup_25_03_2025.tar.gz", "Small backup file")

	assert.Error(t, err)
	mockDataProvider.AssertExpectations(t)
}

func TestCp_ErrorGetFileContent(t *testing.T) {
	mockGistProvider := new(github.MockGistProvider)
	mockDataProvider := new(fsdata.MockDataProvider)
	g3handler := G3BaseHandler{
		G: mockGistProvider,
		D: mockDataProvider,
	}
	mockDataProvider.On("GetFileSize", "/tmp/backup_25_03_2025.tar.gz").Return(int64(500), nil)
	mockDataProvider.On("GetFileContent", "/tmp/backup_25_03_2025.tar.gz").Return([]byte{}, errors.New("could not read the file"))

	err := g3handler.Cp("/tmp/backup_25_03_2025.tar.gz", "Small backup file")

	assert.Error(t, err)
	mockDataProvider.AssertExpectations(t)
}
