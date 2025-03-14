package handlers

import (
	"log/slog"

	"github.com/jegj/g3/config"
	"github.com/jegj/g3/data"
)

func (h *G3BaseHandler) Cp(filepath string) error {
	// TODO: Works with relative and absolute paths
	size, err := h.F.GetFileSize(filepath)
	if err != nil {
		return err
	}
	slog.Info("File processed", "filename", filepath, "size", size)
	// TODO: partition file if required
	content, err := h.F.GetFileContent(filepath)
	if err != nil {
		return err
	}

	// TODO: encrypt data
	filename := h.F.GetFileName(filepath)
	files := map[string]map[string]string{
		filename: {
			"content": string(content),
		},
	}

	gistData, err := h.G.CreateGist("", files, true, config.Conf.GHToken)
	if err != nil {
		return err
	}
	// TODO: sync more than one data file
	gistEntry := data.GistEntry{
		ID:       gistData.Id,
		GistPath: gistData.Url,
	}
	// TODO: what if the file already exists but the content change
	err = h.D.AppendEntry(filename, []data.GistEntry{gistEntry})
	if err != nil {
		return err
	}

	return nil
}
