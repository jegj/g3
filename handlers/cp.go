package handlers

import (
	"log/slog"

	"github.com/jegj/g3/crypto"
	"github.com/jegj/g3/fsdata"
	"github.com/jegj/g3/g3unit"
)

func (h *G3BaseHandler) IsOverrindingFile(filepath string) bool {
	g3unit := g3unit.NewG3Unit(filepath, h.cfg)
	return h.DataService.HasEntry(g3unit.G3Filepath)
}

// TODO: partition file if required
// TODO: sync more than one data file
// TODO: add checksum for each file/part to avoid upload the same chunk from the file
func (h *G3BaseHandler) Cp(filepath string, description string) error {
	g3unit := g3unit.NewG3Unit(filepath, h.cfg)
	filename := fsdata.GetFileName(filepath)

	size, err := h.DataService.GetFileSize(filepath)
	if err != nil {
		return err
	}
	// TODO: REMOVE THIS LATER
	slog.Info("File processed", "filename", filepath, "size", size)

	content, err := h.DataService.GetFileContent(g3unit.Filepath)
	if err != nil {
		return err
	}

	/*
		*  REMOVE THIS
			g3filepath, err := h.DataService.GetG3Filepath(filename)
			if err != nil {
				return err
			}
	*/

	if h.DataService.HasEntry(g3unit.G3Filepath) {
		dataEntry, err := h.DataService.GetEntry(g3unit.G3Filepath)
		if err != nil {
			return err
		}

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

		// TODO: GETTING FIRST ONE
		gistData, err := h.GithubService.UpdateGist(dataEntry.Gist[0].ID, description, files, true)
		if err != nil {
			return err
		}

		gistEntry := fsdata.GistEntry{
			ID:       gistData.Id,
			GistPath: gistData.Url,
		}

		err = h.DataService.UpdateEntry(g3unit.G3Filepath, []fsdata.GistEntry{gistEntry})
		if err != nil {
			return err
		}
		return nil

	} else {
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

		err = h.DataService.AppendEntry(g3unit.G3Filepath, []fsdata.GistEntry{gistEntry})
		if err != nil {
			return err
		}
		return nil
	}
}
