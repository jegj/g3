package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	API_URL                = "https://api.github.com/gists"
	DEFAULT_GITHUB_VERSION = "2022-11-28"
	DEFAULT_ACCEPT_HEADER  = "application/vnd.github+json"
)

func CreateGist(description string, files map[string]map[string]string, public bool, token string) (*GistCreateResponse, error) {
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

	req, err := http.NewRequest("POST", API_URL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", DEFAULT_ACCEPT_HEADER)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("X-GitHub-Api-Version", DEFAULT_GITHUB_VERSION)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("gist creation failed with status code %d: %s", resp.StatusCode, string(body))
	}

	var response GistCreateResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
