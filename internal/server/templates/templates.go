package templates

import (
	"html/template"
	"log"
	"net/http"

	"adamastor/public"
)

type TemplateMap struct {
	entries *template.Template
}

func NewTemplateMap() *TemplateMap {
	files := []string{
		"assets/templates/*html",
		"assets/components/*.html",
		"assets/routes/*/*.html",
	}
	t := template.Must(template.ParseFS(
		public.Assets,
		files...,
	))

	log.Println(t.DefinedTemplates())
	return &TemplateMap{entries: t}
}

func (t *TemplateMap) RenderTemplate(route string, w http.ResponseWriter) error {
	err := t.entries.ExecuteTemplate(w, "layout.html", nil)
	log.Println(t.entries.DefinedTemplates())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Template execution error:", err)
	}
	return err
}

var T = NewTemplateMap()
