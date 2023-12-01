package endpoints

import "adamastor/internal/server/templates"

type Handler struct {
	t *templates.TemplateMap
}

func NewHandler(temps *templates.TemplateMap) *Handler {
	return &Handler{
		t: temps,
	}
}
