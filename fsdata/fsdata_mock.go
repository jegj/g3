package fsdata

import (
	"github.com/jegj/g3/g3unit"
	"github.com/stretchr/testify/mock"
)

type MockDataProvider struct {
	mock.Mock
}

func (m *MockDataProvider) AppendEntry(unit g3unit.G3Unit, gists []GistEntry) error {
	args := m.Called(unit, gists)
	return args.Error(0)
}

func (m *MockDataProvider) UpdateEntry(unit g3unit.G3Unit, gists []GistEntry) error {
	args := m.Called(unit, gists)
	return args.Error(0)
}

func (m *MockDataProvider) DeleteEntry(unit g3unit.G3Unit) error {
	args := m.Called(unit)
	return args.Error(0)
}

func (m *MockDataProvider) GetEntries() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockDataProvider) GetFileSize(unit g3unit.G3Unit) (int64, error) {
	args := m.Called(unit)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockDataProvider) GetFileContent(unit g3unit.G3Unit) ([]byte, error) {
	args := m.Called(unit)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockDataProvider) GetEntry(unit g3unit.G3Unit) (DataEntry, error) {
	args := m.Called(unit)
	return args.Get(0).(DataEntry), args.Error(1)
}

func (m *MockDataProvider) HasEntry(unit g3unit.G3Unit) bool {
	args := m.Called(unit)
	return args.Get(0).(bool)
}
