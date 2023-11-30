package endpoints

import (
	"log"
	"net/http"

	"adamastor/internal/server/templates"
)

func HandleCV(w http.ResponseWriter, _ *http.Request) {
	err := templates.T.RenderTemplate("cv.html", w)
	if err != nil {
		log.Println(err)
	}
}
