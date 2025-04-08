package handlers

import (
	"log/slog"

	"github.com/jegj/g3/crypto"
	"github.com/jegj/g3/fsdata"
)

// TODO: partition file if required
// TODO: sync more than one data file
// TODO: what if the file already exists but the content change
// TODO: add checksum for each file/part to avoid upload the same chunk from the file
func (h *G3BaseHandler) Cp(filepath string, description string) error {
	filename := fsdata.GetFileName(filepath)

	size, err := h.DataService.GetFileSize(filepath)
	if err != nil {
		return err
	}
	// TODO: REMOVE THIS LATER
	slog.Info("File processed", "filename", filepath, "size", size)

	content, err := h.DataService.GetFileContent(filepath)
	if err != nil {
		return err
	}

	g3filepath, err := h.DataService.GetG3Filepath(filename)
	if err != nil {
		return err
	}

	/*
		if h.DataService.HasEntry(g3filepath) {
			dataentry, err := h.DataService.GetEntry(g3filepath)
			if err != nil {
				return err
			}
		} else {
		}
	*/

	encryptedContent, err := crypto.EncryptAESGCM(content, h.cfg.AESKey)
	if err != nil {
		return err
	}

	files := map[string]map[string]string{
		// TODO: USER GENERIC NAME FOR FILES
		filename: {
			"content": string(encryptedContent),
		},
	}

	gistData, err := h.GithubService.CreateGist(description, files, true)
	if err != nil {
		return err
	}
	gistEntry := fsdata.GistEntry{
		ID:       gistData.Id,
		GistPath: gistData.Url,
	}

	err = h.DataService.AppendEntry(g3filepath, []fsdata.GistEntry{gistEntry})
	if err != nil {
		return err
	}

	return nil
}
