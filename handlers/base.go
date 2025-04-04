package handlers

import (
	"github.com/jegj/g3/config"
	"github.com/jegj/g3/fsdata"
	"github.com/jegj/g3/github"
)

type G3BaseHandler struct {
	G   github.GistProvider
	D   fsdata.DataProvider
	cfg config.Config
}

func NewG3BaseHandler(cfg config.Config) G3BaseHandler {
	return G3BaseHandler{
		cfg: cfg,
		G:   github.NewGistService(cfg),
		D:   fsdata.NewDatatService(cfg),
	}
}
