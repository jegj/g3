package handlers

import (
	"errors"
	"testing"

	"github.com/jegj/g3/fsdata"
	"github.com/jegj/g3/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	aesKeyString = "\x1F\x2A\x3C\x4D\x5E\x6F\x7A\x8B\x9C\xAD\xBE\xCF\xD1\xE2\xF3\x04\x15\x26\x37\x48\x59\x6A\x7B\x8C\x9D\xAE\xBF\xC0\xD2\xE3\xF4\x05"
	GH_TOKEN     = "S0m3r4nD0mK3YF0rT3st1ng"
)

var AES_KEY = []byte(aesKeyString)

func TestCp_ErrorGetFileSize(t *testing.T) {
	mockGistProvider := new(github.MockGistProvider)
	mockDataProvider := new(fsdata.MockDataProvider)
	g3handler := G3BaseHandler{
		aeskey:  AES_KEY,
		ghtoken: GH_TOKEN,
		G:       mockGistProvider,
		D:       mockDataProvider,
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
		aeskey:  AES_KEY,
		ghtoken: GH_TOKEN,
		G:       mockGistProvider,
		D:       mockDataProvider,
	}
	mockDataProvider.On("GetFileSize", "/tmp/backup_25_03_2025.tar.gz").Return(int64(500), nil)
	mockDataProvider.On("GetFileContent", "/tmp/backup_25_03_2025.tar.gz").Return([]byte{}, errors.New("could not read the file"))

	err := g3handler.Cp("/tmp/backup_25_03_2025.tar.gz", "Small backup file")

	assert.Error(t, err)
	mockDataProvider.AssertExpectations(t)
}

func TestCp_ErrorCreateGist(t *testing.T) {
	content := []byte{
		0x1F, 0x2A, 0x3C, 0x4D, 0x5E, 0x6F, 0x7A, 0x8B,
		0x9C, 0xAD, 0xBE, 0xCF, 0xD1, 0xE2, 0xF3, 0x04,
		0x15, 0x26, 0x37, 0x48, 0x59, 0x6A, 0x7B, 0x8C,
		0x9D, 0xAE, 0xBF, 0xC0, 0xD2, 0xE3, 0xF4, 0x05,
	}
	description := "Small backup file"
	mockGistProvider := new(github.MockGistProvider)
	mockDataProvider := new(fsdata.MockDataProvider)
	g3handler := G3BaseHandler{
		aeskey:  AES_KEY,
		ghtoken: GH_TOKEN,
		G:       mockGistProvider,
		D:       mockDataProvider,
	}
	/*
		files := map[string]map[string]string{
			"backup_25_03_2025.tar.gz": {
				"content": mock.Anything,
			},
		}
	*/
	mockDataProvider.On("GetFileSize", "/tmp/backup_25_03_2025.tar.gz").Return(int64(500), nil)
	mockDataProvider.On("GetFileContent", "/tmp/backup_25_03_2025.tar.gz").Return(content, nil)
	mockGistProvider.On("CreateGist", description, mock.Anything, true, GH_TOKEN).Return(&github.GistResponse{}, errors.New("could not create gist"))

	err := g3handler.Cp("/tmp/backup_25_03_2025.tar.gz", description)

	assert.Error(t, err)
	mockDataProvider.AssertExpectations(t)
}
