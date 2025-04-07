package handlers

import (
	"github.com/jegj/g3/config"
	"github.com/jegj/g3/fsdata"
	"github.com/jegj/g3/github"
)

type G3BaseHandler struct {
	GithubService github.GistProvider
	DataService   fsdata.DataProvider
	cfg           config.Config
}

func NewG3BaseHandler(cfg config.Config) G3BaseHandler {
	return G3BaseHandler{
		cfg:           cfg,
		GithubService: github.NewGistService(cfg),
		DataService:   fsdata.NewDatatService(cfg),
	}
}
