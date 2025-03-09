package data

type DataEntry map[string]FileGist

type FileGist struct {
	Gist []GistEntry `json:"gist"`
}

type GistEntry struct {
	ID       string `json:"id"`
	GistPath string `json:"gist_path"`
}
