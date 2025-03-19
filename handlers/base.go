package handlers

import (
	"github.com/jegj/g3/fsdata"
	"github.com/jegj/g3/github"
)

type G3BaseHandler struct {
	G github.GistProvider
	D fsdata.DataProvider
}

func NewG3BaseHandler(token string) G3BaseHandler {
	return G3BaseHandler{
		G: github.NewGistService(token),
		D: fsdata.NewDatatService(),
	}
}
