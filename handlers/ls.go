package handlers

import "fmt"

func (g *G3BaseHandler) Ls() error {
	files, err := g.D.GetEntries()
	if err != nil {
		return err
	}
	for _, file := range files {
		fmt.Printf("%s\n", file)
	}
	return nil
}
