package endpoints

import (
	"log"
	"net/http"

	"adamastor/internal/server/templates"
)

func HandleIndex(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := templates.GetTmpl("assets/routes/index/index.html")
	if err != nil {
		log.Println(err)
	}
	if err = tmpl.ExecuteTemplate(w, "layout.html", nil); err != nil {
		log.Println(err)
	}
}
