package handlers

import (
	"github.com/jegj/g3/data"
	"github.com/jegj/g3/fs"
	"github.com/jegj/g3/github"
)

type G3BaseHandler struct {
	G github.GistProvider
	D data.DataProvider
	F fs.FsProvider
}

func NewG3BaseHandler(token string) G3BaseHandler {
	return G3BaseHandler{
		G: github.NewGistService(token),
		D: data.NewDatatService(),
		F: fs.NewFsService(),
	}
}
