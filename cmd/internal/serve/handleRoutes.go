package serve

import (
    "fmt"
    "log"
    "net/http"
    "html/template"
)

func handleRoutes(w http.ResponseWriter, r *http.Request) {    
    if r.URL.Path == "/" {
        handleIndex(w, r)
    }
}

func handleIndex(w http.ResponseWriter, r *http.Request) {    
    tmpl, err := getTmpl("public/index.html")
	if err != nil {
		log.Fatal(err)
	}

    if err := tmpl.Execute(w, nil); err != nil {
			log.Fatal(err)
	}

}

func getTmpl(file string) (*template.Template, error) {
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

