package handlers

import (
	"encoding/json"
	"log/slog"

	"github.com/jegj/g3/config"
	"github.com/jegj/g3/fsdata"
)

// TODO: Handler err when one call to github api fails and the other not
// TODO: Parallel delete calls when the file contains more than one gist
func (g *G3BaseHandler) Rm(filename string) error {
	g3filepath, err := fsdata.GetG3Filepath(filename)
	if err != nil {
		return err
	}
	content, err := g.D.GetFileContent(g3filepath)
	if err != nil {
		return err
	}

	var entry fsdata.DataEntry
	err = json.Unmarshal(content, &entry)
	if err != nil {
		return err
	}

	for filename, fileGist := range entry {
		slog.Debug("Deleting file....", "filename", filename)
		for _, gist := range fileGist.Gist {
			err := g.G.DeleteGist(gist.ID, config.Conf.GHToken)
			if err != nil {
				return err
			}
		}
	}

	err = g.D.DeleteEntry(g3filepath)
	if err != nil {
		return err
	}
	return nil
}
