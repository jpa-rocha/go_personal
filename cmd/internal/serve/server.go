package serve

import (
    "net/http"
)

type Server struct {
    Config      *http.Server
    Router      *Router
    Err         error
}

func NewServer(config *http.Server, router *Router) *Server {
    return &Server{
        Config: config,
        Router: router,
        Err: nil,
    }
} 

func (s *Server) ServeStaticPaths() {
    s.Router.setStaticPaths()
    s.Err = s.Router.Err
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {    
    s.Router.handleRoutes(w, r)
}
