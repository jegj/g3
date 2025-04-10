package github

import "github.com/stretchr/testify/mock"

type MockGistProvider struct {
	mock.Mock
}

func (m *MockGistProvider) CreateGist(description string, files map[string]map[string]string, public bool) (*GistResponse, error) {
	args := m.Called(description, files, public)
	return args.Get(0).(*GistResponse), args.Error(1)
}

func (m *MockGistProvider) DeleteGist(id string) error {
	args := m.Called(id)
	return args.Error(1)
}

func (m *MockGistProvider) UpdateGist(id string, description string, files map[string]map[string]string, public bool) (*GistResponse, error) {
	args := m.Called(id, description, files, public)
	return args.Get(0).(*GistResponse), args.Error(1)
}
