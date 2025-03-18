package handlers

// TODO: FS and DATA similiar behavior
func (g *G3BaseHandler) Ls() ([]string, error) {
	files, err := g.D.GetEntries()
	if err != nil {
		return []string{}, err
	}
	return files, nil
}
