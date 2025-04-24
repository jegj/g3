package handlers

import (
	"github.com/jegj/g3/g3unit"
)

// TODO: Handler err when one call to github api fails and the other not
// TODO: Parallel delete calls when the file contains more than one gist
func (h *G3BaseHandler) Rm(filename string) error {
	g3unit := g3unit.NewG3Unit(filename, h.cfg)
	entry, err := h.DataService.GetEntry(g3unit)
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
