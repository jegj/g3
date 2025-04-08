package handlers

import (
	"encoding/json"

	"github.com/jegj/g3/fsdata"
)

// TODO: Handler err when one call to github api fails and the other not
// TODO: Parallel delete calls when the file contains more than one gist
func (h *G3BaseHandler) Rm(filename string) error {
	g3filepath, err := h.DataService.GetG3Filepath(filename)
	if err != nil {
		return err
	}
	content, err := h.DataService.GetFileContent(g3filepath)
	if err != nil {
		return err
	}

	var entry fsdata.DataEntry
	err = json.Unmarshal(content, &entry)
	if err != nil {
		return err
	}

	for _, gist := range entry.Gist {
		err := h.GithubService.DeleteGist(gist.ID)
		if err != nil {
			return err
		}
	}

	err = h.DataService.DeleteEntry(g3filepath)
	if err != nil {
		return err
	}
	return nil
}
