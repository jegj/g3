package handlers

import (
	"log/slog"

	"github.com/jegj/g3/config"
	"github.com/jegj/g3/fsdata"
)

// TODO: Works with relative and absolute paths
// TODO: partition file if required
// TODO: encrypt data
// TODO: sync more than one data file
// TODO: what if the file already exists but the content change
// TODO: add createdAt time for whole file
func (h *G3BaseHandler) Cp(filepath string, description string) error {
	size, err := h.D.GetFileSize(filepath)
	if err != nil {
		return err
	}
	slog.Info("File processed", "filename", filepath, "size", size)
	content, err := h.D.GetFileContent(filepath)
	if err != nil {
		return err
	}

	filename := h.D.GetFileName(filepath)
	files := map[string]map[string]string{
		filename: {
			"content": string(content),
		},
	}

	gistData, err := h.G.CreateGist(description, files, true, config.Conf.GHToken)
	if err != nil {
		return err
	}
	gistEntry := fsdata.GistEntry{
		ID:       gistData.Id,
		GistPath: gistData.Url,
	}
	g3filepath, err := fsdata.GetG3Filepath(filename)
	if err != nil {
		return nil
	}

	err = h.D.AppendEntry(g3filepath, []fsdata.GistEntry{gistEntry})
	if err != nil {
		return err
	}

	return nil
}
