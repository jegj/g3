package fsdata

import "time"

type DataEntry struct {
	Gist      []GistEntry `json:"gist"`
	CreatedAt time.Time   `json:"created_at"`
}

type GistEntry struct {
	ID       string `json:"id"`
	GistPath string `json:"gist_path"`
}
