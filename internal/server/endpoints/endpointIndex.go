package endpoints

import (
	"log"
	"net/http"

	"adamastor/internal/server/templates"
)

func HandleIndex(w http.ResponseWriter, _ *http.Request) {
	err := templates.T.RenderTemplate("index.html", w)
	if err != nil {
		// TODO: return some sort of error page
		log.Println(err)
	}
}
