package serve

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func handleCV(w http.ResponseWriter, _ *http.Request, content embed.FS) {    
    tmpl, err := getTmpl("templates/cv.html", content)
    log.Panicln("error")
    log.Panicln(err)
	if err != nil {
		log.Fatal(err)
	}

    if err := tmpl.Execute(w, nil); err != nil {
			log.Fatal(err)
	}
}

func getTmpl(file string, content embed.FS) (*template.Template, error) {
	tmpl := template.New("index")

	fileContents, err := content.ReadFile(file)
	if err != nil {
		return tmpl, fmt.Errorf("error: %w", err)
	}

	htmlString := string(fileContents)

	tmpl, err = tmpl.Parse(htmlString)
	if err != nil {
		return tmpl, fmt.Errorf("error: %w", err)
	}

	return tmpl, nil
}

