package router

import (
	"io/fs"
	"net/http"

	"github.com/a-h/templ"

	"adamastor/internal/server/templates"
	"adamastor/public"
)

type Router struct {
	Mux *http.ServeMux
	Err error
}

func NewRouter() *Router {
	router := Router{
		Mux: http.NewServeMux(),
		Err: nil,
	}
	return &router
}

func (r *Router) SetStaticPaths() {
	static := []string{
		"/static/",
		"/components/",
		"/routes/",
		"/templates",
	} // TODO: should be in config.

	const base = "assets"
	staticFS, err := fs.Sub(public.Assets, base)
	if err != nil {
		r.Err = err
		return
	}
	for _, path := range static {
		r.Mux.Handle(path, http.FileServer(http.FS(staticFS)))
	}
}

func (r *Router) HandleRoutes() {
	r.SetStaticPaths()

	r.Mux.Handle("/cv", templ.Handler(templates.Layout(templates.CV())))
	r.Mux.Handle("/projects", templ.Handler(templates.Layout(templates.Project())))
	r.Mux.Handle("/littleprofessor", templ.Handler(templates.Professor()))
    
	r.Mux.Handle("/", templ.Handler(templates.Layout(templates.Index())))
}
