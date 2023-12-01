package endpoints

import (
	"log"
	"net/http"
)

func (h *Handler) HandleCV(w http.ResponseWriter, _ *http.Request) {
	err := h.t.RenderTemplate(w, "cv.html", nil)
	if err != nil {
		log.Println(err)
	}
}
