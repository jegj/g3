package github

import "github.com/stretchr/testify/mock"

type MockGistProvider struct {
	mock.Mock
}

func (m *MockGistProvider) CreateGist(description string, files map[string]map[string]string, public bool, token string) (*GistResponse, error) {
	args := m.Called(description, files, public, token)
	return args.Get(0).(*GistResponse), args.Error(1)
}

func (m *MockGistProvider) DeleteGist(id string, token string) error {
	args := m.Called(id, token)
	return args.Error(1)
}
