package handlers

import (
	"log/slog"

	"github.com/jegj/g3/config"
	"github.com/jegj/g3/data"
)

// TODO: Works with relative and absolute paths
// TODO: partition file if required
// TODO: encrypt data
// TODO: sync more than one data file
// TODO: what if the file already exists but the content change
// TODO: add createdAt time for whole file
func (h *G3BaseHandler) Cp(filepath string, description string) error {
	size, err := h.F.GetFileSize(filepath)
	if err != nil {
		return err
	}
	slog.Info("File processed", "filename", filepath, "size", size)
	content, err := h.F.GetFileContent(filepath)
	if err != nil {
		return err
	}

	filename := h.F.GetFileName(filepath)
	files := map[string]map[string]string{
		filename: {
			"content": string(content),
		},
	}

	gistData, err := h.G.CreateGist(description, files, true, config.Conf.GHToken)
	if err != nil {
		return err
	}
	gistEntry := data.GistEntry{
		ID:       gistData.Id,
		GistPath: gistData.Url,
	}
	err = h.D.AppendEntry(filename, []data.GistEntry{gistEntry})
	if err != nil {
		return err
	}

	return nil
}
