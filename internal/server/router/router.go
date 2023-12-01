package router

import (
	"io/fs"
	"net/http"

	"adamastor/internal/server/endpoints"
	"adamastor/internal/server/templates"
	"adamastor/public"
)

type Router struct {
	Mux     *http.ServeMux
	Handler *endpoints.Handler
	Err     error
}

func NewRouter() *Router {
	router := Router{
		Mux:     http.NewServeMux(),
		Handler: endpoints.NewHandler(templates.NewTemplateMap()),
		Err:     nil,
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

	r.Mux.HandleFunc("/cv", r.Handler.HandleCV)

	r.Mux.HandleFunc("/components/nav", r.Handler.HandleComponentNav)

	r.Mux.HandleFunc("/", r.Handler.HandleIndex)
}
