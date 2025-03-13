package handlers

import "github.com/jegj/g3/github"

type G3BaseHandler struct {
	G github.GistProvider
	// Fs
	// Data
}

func NewG3BaseHandler(token string) G3BaseHandler {
	return G3BaseHandler{
		G: github.NewGistService(token),
	}
}
