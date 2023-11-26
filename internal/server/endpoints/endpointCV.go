package endpoints

import (
	"adamastor/internal/server/templates"
	"log"
	"net/http"
)

func HandleCV(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := templates.GetTmpl("assets/routes/cv/cv.html")
	if err != nil {
		log.Println(err)
	}

	if err = tmpl.Execute(w, nil); err != nil {
		log.Println(err)
	}
}
