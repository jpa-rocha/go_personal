package endpoints

import (
	"log"
	"net/http"

	"adamastor/internal/server/templates"
)

func HandleComponentNav(w http.ResponseWriter, _ *http.Request) {
	err := templates.T.RenderTemplate("nav.html", w)
	if err != nil {
		log.Println(err)
	}
}
