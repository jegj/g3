package handlers

import (
	"fmt"
	"log/slog"

	"github.com/jegj/g3/config"
	"github.com/jegj/g3/data"
	"github.com/jegj/g3/fs"
	"github.com/jegj/g3/github"
)

func Cp(filepath string) error {
	size, err := fs.GetFileSize(filepath)
	if err != nil {
		return err
	}
	slog.Info("File processed", "filename", filepath, "size", size)
	// TODO: partition file if required
	content, err := fs.GetFileContent(filepath)
	if err != nil {
		return err
	}
	fmt.Println("First 10 bytes:", string(content[:10]))

	// TODO: encrypt data
	filename := fs.GetFileName(filepath)
	files := map[string]map[string]string{
		filename: {
			"content": string(content),
		},
	}

	gistData, err := github.CreateGist("", files, true, config.Conf.GHToken)
	if err != nil {
		return err
	}
	// TODO: sync more than one data file
	gistEntry := data.GistEntry{
		ID:       gistData.Id,
		GistPath: gistData.Url,
	}
	err = data.AppendEntry(filename, []data.GistEntry{gistEntry})
	if err != nil {
		return err
	}

	return nil
}
