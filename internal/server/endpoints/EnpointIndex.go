package endpoints

import (
	"adamastor/internal/server/utilities"
	"log"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := utilities.GetTmpl("assets/routes/index/index.html")
	if err != nil {
		log.Println(err)
	}
	if err = tmpl.Execute(w, nil); err != nil {
		log.Println(err)
	}
}
