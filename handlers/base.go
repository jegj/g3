package handlers

import (
	"github.com/jegj/g3/fsdata"
	"github.com/jegj/g3/github"
)

type G3BaseHandler struct {
	ghtoken string
	aeskey  []byte
	G       github.GistProvider
	D       fsdata.DataProvider
}

func NewG3BaseHandler(token string, aeskey []byte) G3BaseHandler {
	return G3BaseHandler{
		ghtoken: token,
		aeskey:  aeskey,
		G:       github.NewGistService(token),
		D:       fsdata.NewDatatService(),
	}
}
