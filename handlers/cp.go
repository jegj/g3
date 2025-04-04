package handlers

import (
	"log/slog"

	"github.com/jegj/g3/crypto"
	"github.com/jegj/g3/fsdata"
)

// TODO: partition file if required
// TODO: sync more than one data file
// TODO: what if the file already exists but the content change
// TODO: add createdAt time for whole file
func (h *G3BaseHandler) Cp(filepath string, description string) error {
	filename := fsdata.GetFileName(filepath)

	size, err := h.D.GetFileSize(filepath)
	if err != nil {
		return err
	}
	// TODO: REMOVE THIS LATER
	slog.Info("File processed", "filename", filepath, "size", size)

	content, err := h.D.GetFileContent(filepath)
	if err != nil {
		return err
	}

	g3filepath, err := h.D.GetG3Filepath(filename)
	if err != nil {
		return err
	}

	//TODO: Clean up this
	/*
		if h.D.HasEntry(g3filepath) {
			dataentry, err := h.D.GetEntry(g3filepath)
			if err != nil {
				return err
			}
		} else {
		}*/

	encryptedContent, err := crypto.EncryptAESGCM(content, h.cfg.AESKey)
	if err != nil {
		return err
	}

	files := map[string]map[string]string{
		filename: {
			"content": string(encryptedContent),
		},
	}

	gistData, err := h.G.CreateGist(description, files, true, h.cfg.GHToken)
	if err != nil {
		return err
	}
	gistEntry := fsdata.GistEntry{
		ID:       gistData.Id,
		GistPath: gistData.Url,
	}

	err = h.D.AppendEntry(g3filepath, []fsdata.GistEntry{gistEntry})
	if err != nil {
		return err
	}

	return nil
}
