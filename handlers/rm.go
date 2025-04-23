package handlers

import (
	"encoding/json"

	"github.com/jegj/g3/fsdata"
	"github.com/jegj/g3/g3unit"
)

// TODO: Handler err when one call to github api fails and the other not
// TODO: Parallel delete calls when the file contains more than one gist
func (h *G3BaseHandler) Rm(filename string) error {
	g3unit := g3unit.NewG3Unit(filename, h.cfg)
	content, err := h.DataService.GetFileContent(g3unit)
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

	err = h.DataService.DeleteEntry(g3unit)
	if err != nil {
		return err
	}
	return nil
}
