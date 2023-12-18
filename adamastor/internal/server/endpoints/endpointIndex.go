package endpoints

import (
	"log"
	"net/http"
)

func (h *Handler) HandleIndex(w http.ResponseWriter, _ *http.Request) {
	err := h.t.RenderTemplate(w, "index.html", nil)
	if err != nil {
		// TODO: return some sort of error page.
		log.Println(err)
	}
}
