package templates

import (
	"html/template"
	"log"
	"net/http"

	"adamastor/public"
)

func RenderTemplate(route string, w http.ResponseWriter) error {
	t := AssembleTemplate(route)
	err := t.ExecuteTemplate(w, "layout.html", "content")
	if err != nil {
		log.Println(err)
	}
	return err
}

func AssembleTemplate(route string) *template.Template {
	t := template.Must(template.ParseFS(
		public.Assets,
		"assets/templates/layout.html",
		route))
	log.Println(t.DefinedTemplates())
	return t
}
