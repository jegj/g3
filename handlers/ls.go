package handlers

func (h *G3BaseHandler) Ls() ([]string, error) {
	files, err := h.DataService.GetEntries()
	if err != nil {
		return []string{}, err
	}
	return files, nil
}
