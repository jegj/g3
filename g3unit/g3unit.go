package g3unit

import (
	"fmt"
	"path/filepath"

	"github.com/jegj/g3/config"
)

type G3Unit struct {
	G3Filename string
	G3Filepath string
	Filename   string
	Filepath   string
}

func NewG3Unit(fpath string, cfg config.Config) G3Unit {
	filename := filepath.Base(fpath)
	g3filename := fmt.Sprintf("%s.g3.json", filename)
	return G3Unit{
		G3Filename: g3filename,
		G3Filepath: filepath.Join(cfg.DataFolder, g3filename),
		Filename:   filename,
		Filepath:   fpath,
	}
}
