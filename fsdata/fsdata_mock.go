package fsdata

import (
	"github.com/stretchr/testify/mock"
)

type MockDataProvider struct {
	mock.Mock
}

func (m *MockDataProvider) AppendEntry(filename string, gists []GistEntry) error {
	args := m.Called(filename, gists)
	return args.Error(0)
}

func (m *MockDataProvider) UpdateEntry(filename string, gists []GistEntry) error {
	args := m.Called(filename, gists)
	return args.Error(0)
}

func (m *MockDataProvider) DeleteEntry(filename string) error {
	args := m.Called(filename)
	return args.Error(0)
}

func (m *MockDataProvider) GetEntries() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockDataProvider) GetFileSize(absFilePath string) (int64, error) {
	args := m.Called(absFilePath)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockDataProvider) GetFileContent(absFilePath string) ([]byte, error) {
	args := m.Called(absFilePath)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockDataProvider) GetEntry(filename string) (DataEntry, error) {
	args := m.Called(filename)
	return args.Get(0).(DataEntry), args.Error(1)
}

func (m *MockDataProvider) HasEntry(filename string) bool {
	args := m.Called(filename)
	return args.Get(0).(bool)
}

func (m *MockDataProvider) GetG3Filepath(filename string) (string, error) {
	args := m.Called(filename)
	return args.Get(0).(string), args.Error(1)
}
