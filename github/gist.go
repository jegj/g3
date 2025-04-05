package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jegj/g3/config"
)

const (
	API_URL                = "https://api.github.com/gists"
	DEFAULT_GITHUB_VERSION = "2022-11-28"
	DEFAULT_ACCEPT_HEADER  = "application/vnd.github+json"
)

type GistProvider interface {
	CreateGist(description string, files map[string]map[string]string, public bool) (*GistResponse, error)
	DeleteGist(id string) error
}

type GistService struct {
	cfg config.Config
}

func NewGistService(cfg config.Config) GistService {
	return GistService{
		cfg: cfg,
	}
}

func (g GistService) CreateGist(description string, files map[string]map[string]string, public bool) (*GistResponse, error) {
	client := &http.Client{}
	requestData := GistCreateRequest{
		Description: description,
		Public:      public,
		Files:       files,
	}
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, API_URL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", DEFAULT_ACCEPT_HEADER)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", g.cfg.GHToken))
	req.Header.Set("X-GitHub-Api-Version", DEFAULT_GITHUB_VERSION)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if cerr := resp.Body.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("gist creation failed with status code %d: %s", resp.StatusCode, string(body))
	}

	var response GistResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (g GistService) DeleteGist(id string) (err error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s/%s", API_URL, id)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", DEFAULT_ACCEPT_HEADER)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", g.cfg.GHToken))
	req.Header.Set("X-GitHub-Api-Version", DEFAULT_GITHUB_VERSION)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		if cerr := resp.Body.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("gist deletion failed with status code %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

func (g GistService) GetGist(id string) (err error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s/%s", API_URL, id)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", DEFAULT_ACCEPT_HEADER)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", g.cfg.GHToken))
	req.Header.Set("X-GitHub-Api-Version", DEFAULT_GITHUB_VERSION)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		if cerr := resp.Body.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("gist deletion failed with status code %d: %s", resp.StatusCode, string(body))
	}

	return nil
}
