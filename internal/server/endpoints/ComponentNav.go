
package endpoints

import (
	"adamastor/internal/server/utilities"
	"log"
	"net/http"
)

func HandleComponentNav(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := utilities.GetTmpl("assets/components/nav.html")
	if err != nil {
		log.Println(err)
	}

	if err = tmpl.Execute(w, nil); err != nil {
		log.Println(err)
	}
}
