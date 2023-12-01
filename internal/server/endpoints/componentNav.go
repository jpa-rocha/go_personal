package endpoints

import (
	"log"
	"net/http"
)

func (h *Handler) HandleComponentNav(w http.ResponseWriter, _ *http.Request) {
	err := h.t.RenderTemplate(w, "nav.html", nil)
	if err != nil {
		log.Println(err)
	}
}
