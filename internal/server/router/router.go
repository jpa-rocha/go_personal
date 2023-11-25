package router

import (
	"io/fs"
	"net/http"

	"adamastor/internal/server/endpoints"
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
	static := []string{"/static/", "/components/", "/routes/"} // TODO: should be in config.

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
	r.Mux.HandleFunc("/cv", endpoints.HandleCV)

	r.Mux.HandleFunc("/components/nav", endpoints.HandleComponentNav)

	r.Mux.HandleFunc("/", endpoints.HandleIndex)
}
