package handlers

import (
	"errors"
	"testing"
	"time"

	"github.com/jegj/g3/config"
	"github.com/jegj/g3/fsdata"
	"github.com/jegj/g3/g3unit"
	"github.com/jegj/g3/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var cfg = config.Config{
	GHToken:    "S0m3r4nD0mK3YF0rT3st1ng",
	AESKey:     []byte("\x1F\x2A\x3C\x4D\x5E\x6F\x7A\x8B\x9C\xAD\xBE\xCF\xD1\xE2\xF3\x04\x15\x26\x37\x48\x59\x6A\x7B\x8C\x9D\xAE\xBF\xC0\xD2\xE3\xF4\x05"),
	DataFolder: "/home/testy/.local/share/g3/files",
}

func TestCp_ErrorGetFileSize(t *testing.T) {
	mockGistProvider := new(github.MockGistProvider)
	mockDataProvider := new(fsdata.MockDataProvider)
	g3handler := G3BaseHandler{
		cfg:           cfg,
		GithubService: mockGistProvider,
		DataService:   mockDataProvider,
	}
	inputFilePath := "/tmp/backup_25_03_2025.tar.gz"
	unit := g3unit.NewG3Unit(inputFilePath, cfg)
	mockDataProvider.On("GetFileSize", unit).Return(int64(0), errors.New("file not found"))

	err := g3handler.Cp(inputFilePath, "Small backup file")

	assert.Error(t, err)
	mockDataProvider.AssertExpectations(t)
}

func TestCp_ErrorGetFileContent(t *testing.T) {
	mockGistProvider := new(github.MockGistProvider)
	mockDataProvider := new(fsdata.MockDataProvider)
	g3handler := G3BaseHandler{
		cfg:           cfg,
		GithubService: mockGistProvider,
		DataService:   mockDataProvider,
	}
	inputFilePath := "/tmp/backup_25_03_2025.tar.gz"
	unit := g3unit.NewG3Unit(inputFilePath, cfg)
	mockDataProvider.On("GetFileSize", unit).Return(int64(500), nil)
	mockDataProvider.On("GetFileContent", unit).Return([]byte{}, errors.New("could not read the file"))

	err := g3handler.Cp(inputFilePath, "Small backup file")

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
		cfg:           cfg,
		GithubService: mockGistProvider,
		DataService:   mockDataProvider,
	}
	inputFilePath := "/tmp/backup_25_03_2025.tar.gz"
	unit := g3unit.NewG3Unit(inputFilePath, cfg)
	mockDataProvider.On("GetFileSize", unit).Return(int64(500), nil)
	mockDataProvider.On("GetFileContent", unit).Return(content, nil)
	mockDataProvider.On("HasEntry", unit).Return(false)
	mockGistProvider.On("CreateGist", description, mock.Anything, true).Return(&github.GistResponse{}, errors.New("could not create gist"))

	err := g3handler.Cp(inputFilePath, description)

	assert.Error(t, err)
	mockDataProvider.AssertExpectations(t)
	mockGistProvider.AssertExpectations(t)
}

func TestCp_ErrorAppendEntry(t *testing.T) {
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
		cfg:           cfg,
		GithubService: mockGistProvider,
		DataService:   mockDataProvider,
	}
	inputFilePath := "/tmp/backup_25_03_2025.tar.gz"
	unit := g3unit.NewG3Unit(inputFilePath, cfg)
	githubResponse := &github.GistResponse{
		Url:       "https://api.github.com/gists/88dfea43e4262d04557b01c0cfc3c7ba",
		Id:        "88dfea43e4262d04557b01c0cfc3c7ba",
		Public:    false,
		Truncated: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Files: map[string]github.FileGist{
			"example.go": {
				Filename:  "profile.txt",
				Type:      "text/plain",
				Size:      123,
				RawUrl:    "https://gist.githubusercontent.com/raw/example",
				Truncated: false,
				Content:   "�Ex6���=���Tg!�^�7�Tq\x16ȏ�뛫�0�u\b��䣖tN\x1f", // encrypted content
			},
		},
	}
	gistEntry := fsdata.GistEntry{
		ID:       githubResponse.Id,
		GistPath: githubResponse.Url,
	}
	mockDataProvider.On("GetFileSize", unit).Return(int64(500), nil)
	mockDataProvider.On("GetFileContent", unit).Return(content, nil)
	mockDataProvider.On("HasEntry", unit).Return(false)
	mockGistProvider.On("CreateGist", description, mock.Anything, true).Return(githubResponse, nil)
	mockDataProvider.On("AppendEntry", unit, []fsdata.GistEntry{
		gistEntry,
	}).Return(errors.New("could not append entry"))

	err := g3handler.Cp("/tmp/backup_25_03_2025.tar.gz", description)

	assert.Error(t, err)
	mockDataProvider.AssertExpectations(t)
	mockGistProvider.AssertExpectations(t)
}

func TestCp_SuccessOneFileNoOverride(t *testing.T) {
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
		cfg:           cfg,
		GithubService: mockGistProvider,
		DataService:   mockDataProvider,
	}
	inputFilePath := "/tmp/backup_25_03_2025.tar.gz"
	unit := g3unit.NewG3Unit(inputFilePath, cfg)
	githubResponse := &github.GistResponse{
		Url:       "https://api.github.com/gists/88dfea43e4262d04557b01c0cfc3c7ba",
		Id:        "88dfea43e4262d04557b01c0cfc3c7ba",
		Public:    false,
		Truncated: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Files: map[string]github.FileGist{
			"example.go": {
				Filename:  "profile.txt",
				Type:      "text/plain",
				Size:      123,
				RawUrl:    "https://gist.githubusercontent.com/raw/example",
				Truncated: false,
				Content:   "�Ex6���=���Tg!�^�7�Tq\x16ȏ�뛫�0�u\b��䣖tN\x1f", // encrypted content
			},
		},
	}
	gistEntry := fsdata.GistEntry{
		ID:       githubResponse.Id,
		GistPath: githubResponse.Url,
	}
	mockDataProvider.On("GetFileSize", unit).Return(int64(500), nil)
	mockDataProvider.On("GetFileContent", unit).Return(content, nil)
	mockDataProvider.On("HasEntry", unit).Return(false)
	mockGistProvider.On("CreateGist", description, mock.Anything, true).Return(githubResponse, nil)
	mockDataProvider.On("AppendEntry", unit, []fsdata.GistEntry{
		gistEntry,
	}).Return(nil)

	err := g3handler.Cp(inputFilePath, description)

	assert.NoError(t, err)
	mockDataProvider.AssertExpectations(t)
	mockGistProvider.AssertExpectations(t)
}

func TestCp_ErrorGettingEntryOnFileOverride(t *testing.T) {
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
		cfg:           cfg,
		GithubService: mockGistProvider,
		DataService:   mockDataProvider,
	}
	inputFilePath := "/tmp/backup_25_03_2025.tar.gz"
	unit := g3unit.NewG3Unit(inputFilePath, cfg)
	mockDataProvider.On("GetFileSize", unit).Return(int64(500), nil)
	mockDataProvider.On("GetFileContent", unit).Return(content, nil)
	mockDataProvider.On("HasEntry", unit).Return(true)
	mockDataProvider.On("GetEntry", unit).Return(fsdata.DataEntry{}, errors.New("failed to load file"))

	err := g3handler.Cp(inputFilePath, description)

	assert.Error(t, err)
	mockDataProvider.AssertExpectations(t)
}

func TestCp_ErrorUpdateGistOverride(t *testing.T) {
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
		cfg:           cfg,
		GithubService: mockGistProvider,
		DataService:   mockDataProvider,
	}
	inputFilePath := "/tmp/backup_25_03_2025.tar.gz"
	unit := g3unit.NewG3Unit(inputFilePath, cfg)
	dataEntry := fsdata.DataEntry{
		CreatedAt: time.Now(),
		Gist: []fsdata.GistEntry{
			{
				ID:       "2decf6c462d9b4418f2",
				GistPath: "https://api.github.com/gists/2decf6c462d9b4418f2",
			},
		},
	}
	mockDataProvider.On("GetFileSize", unit).Return(int64(500), nil)
	mockDataProvider.On("GetFileContent", unit).Return(content, nil)
	mockDataProvider.On("HasEntry", unit).Return(true)
	mockDataProvider.On("GetEntry", unit).Return(dataEntry, nil)
	mockGistProvider.On("UpdateGist", "2decf6c462d9b4418f2", description, mock.Anything, true).Return(&github.GistResponse{}, errors.New("failed to update gist"))

	err := g3handler.Cp(inputFilePath, description)

	assert.Error(t, err)
	mockDataProvider.AssertExpectations(t)
}

func TestCp_ErrorAppendEntryOverride(t *testing.T) {
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
		cfg:           cfg,
		GithubService: mockGistProvider,
		DataService:   mockDataProvider,
	}
	inputFilePath := "/tmp/backup_25_03_2025.tar.gz"
	unit := g3unit.NewG3Unit(inputFilePath, cfg)
	dataEntry := fsdata.DataEntry{
		CreatedAt: time.Now(),
		Gist: []fsdata.GistEntry{
			{
				ID:       "2decf6c462d9b4418f2",
				GistPath: "https://api.github.com/gists/2decf6c462d9b4418f2",
			},
		},
	}
	githubResponse := &github.GistResponse{
		Url:       "https://api.github.com/gists/99dfea43e4262c04557b12c0cfc6c7jn",
		Id:        "99dfea43e4262c04557b12c0cfc6c7jn",
		Public:    false,
		Truncated: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Files: map[string]github.FileGist{
			"backup_25_03_2025.tar.gz": {
				Filename:  "backup_25_03_2025.tar.gz",
				Type:      "text/plain",
				Size:      200,
				RawUrl:    "https://gist.githubusercontent.com/raw/backup_25_03_2025.tar.gz",
				Truncated: false,
				Content:   "�Ex6���=���Tg!�^�7�Tq\x16ȏ�뛫�0�u\b��䣖tN\x1f", // encrypted content
			},
		},
	}
	gistEntry := fsdata.GistEntry{
		ID:       githubResponse.Id,
		GistPath: githubResponse.Url,
	}
	mockDataProvider.On("GetFileSize", unit).Return(int64(500), nil)
	mockDataProvider.On("GetFileContent", unit).Return(content, nil)
	mockDataProvider.On("HasEntry", unit).Return(true)
	mockDataProvider.On("GetEntry", unit).Return(dataEntry, nil)
	mockGistProvider.On("UpdateGist", "2decf6c462d9b4418f2", description, mock.Anything, true).Return(githubResponse, nil)
	mockDataProvider.On("UpdateEntry", unit, []fsdata.GistEntry{
		gistEntry,
	}).Return(errors.New("could not append entry"))

	err := g3handler.Cp(inputFilePath, description)

	assert.Error(t, err)
	mockDataProvider.AssertExpectations(t)
}

func TestCp_SuccessOverride(t *testing.T) {
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
		cfg:           cfg,
		GithubService: mockGistProvider,
		DataService:   mockDataProvider,
	}
	inputFilePath := "/tmp/backup_25_03_2025.tar.gz"
	unit := g3unit.NewG3Unit(inputFilePath, cfg)
	dataEntry := fsdata.DataEntry{
		CreatedAt: time.Now(),
		Gist: []fsdata.GistEntry{
			{
				ID:       "2decf6c462d9b4418f2",
				GistPath: "https://api.github.com/gists/2decf6c462d9b4418f2",
			},
		},
	}
	githubResponse := &github.GistResponse{
		Url:       "https://api.github.com/gists/99dfea43e4262c04557b12c0cfc6c7jn",
		Id:        "99dfea43e4262c04557b12c0cfc6c7jn",
		Public:    false,
		Truncated: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Files: map[string]github.FileGist{
			"backup_25_03_2025.tar.gz": {
				Filename:  "backup_25_03_2025.tar.gz",
				Type:      "text/plain",
				Size:      200,
				RawUrl:    "https://gist.githubusercontent.com/raw/backup_25_03_2025.tar.gz",
				Truncated: false,
				Content:   "�Ex6���=���Tg!�^�7�Tq\x16ȏ�뛫�0�u\b��䣖tN\x1f", // encrypted content
			},
		},
	}
	gistEntry := fsdata.GistEntry{
		ID:       githubResponse.Id,
		GistPath: githubResponse.Url,
	}
	mockDataProvider.On("GetFileSize", unit).Return(int64(500), nil)
	mockDataProvider.On("GetFileContent", unit).Return(content, nil)
	mockDataProvider.On("HasEntry", unit).Return(true)
	mockDataProvider.On("GetEntry", unit).Return(dataEntry, nil)
	mockGistProvider.On("UpdateGist", "2decf6c462d9b4418f2", description, mock.Anything, true).Return(githubResponse, nil)
	mockDataProvider.On("UpdateEntry", unit, []fsdata.GistEntry{
		gistEntry,
	}).Return(nil)

	err := g3handler.Cp(inputFilePath, description)

	assert.NoError(t, err)
	mockDataProvider.AssertExpectations(t)
	mockGistProvider.AssertExpectations(t)
}
