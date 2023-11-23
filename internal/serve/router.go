package serve

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"

	"adamastor/public"
)

func (s *Server) setStaticPaths() {
    static := []string{"/", "/static/", "/templates/", "/components/"} //TODO: should be in config

	const base = "assets"
	staticFS, err := fs.Sub(public.Assets, base)
	if err != nil {
		s.Err = err
		return
	}
	for _, path := range static {
		s.Router.Handle(path, http.StripPrefix(path, http.FileServer(http.FS(staticFS))))
	}
}

func (s *Server) handleRoutes() {
	s.Router.HandleFunc("/cv", handleCV)
}

func handleCV(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := getTmpl("templates/cv.html")
	if err != nil {
		log.Println(err)
	}

	if err = tmpl.Execute(w, nil); err != nil {
		log.Println(err)
	}
}

func getTmpl(file string) (*template.Template, error) {
	tmpl := template.New("index")

	fileContents, err := public.Assets.ReadFile(file)
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
