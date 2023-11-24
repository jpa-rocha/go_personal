package serve

import (
	"log"
	"net/http"
)

func handleComponentTest(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := getTmpl("assets/components/test.html")
	if err != nil {
		log.Println(err)
	}

	if err = tmpl.Execute(w, nil); err != nil {
		log.Println(err)
	}
}

func handleComponentNav(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := getTmpl("assets/components/nav.html")
	if err != nil {
		log.Println(err)
	}

	if err = tmpl.Execute(w, nil); err != nil {
		log.Println(err)
	}
}
