package handlers

import (
	"log/slog"

	"github.com/jegj/g3/crypto"
	"github.com/jegj/g3/fsdata"
	"github.com/jegj/g3/g3unit"
)

func (h *G3BaseHandler) IsOverrindingFile(inputFilePath string) bool {
	g3unit := g3unit.NewG3Unit(inputFilePath, h.cfg)
	return h.DataService.HasEntry(g3unit)
}

// TODO: partition file if required
// TODO: sync more than one data file
func (h *G3BaseHandler) Cp(inputFilePath string, description string) error {
	g3unit := g3unit.NewG3Unit(inputFilePath, h.cfg)
	filename := g3unit.Filename

	size, err := h.DataService.GetFileSize(g3unit)
	if err != nil {
		return err
	}
	// TODO: REMOVE THIS LATER
	slog.Info("File processed", "filename", inputFilePath, "size", size)

	content, err := h.DataService.GetFileContent(g3unit)
	if err != nil {
		return err
	}

	if h.DataService.HasEntry(g3unit) {
		dataEntry, err := h.DataService.GetEntry(g3unit)
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

		err = h.DataService.UpdateEntry(g3unit, []fsdata.GistEntry{gistEntry})
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

		err = h.DataService.AppendEntry(g3unit, []fsdata.GistEntry{gistEntry})
		if err != nil {
			return err
		}
		return nil
	}
}
