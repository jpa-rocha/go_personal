package serve

import (
	"embed"
	"io/fs"
    "net/http"
	"path/filepath"
)

type Server struct {
    Config      *http.Server
    FileSystem  embed.FS
    Err         error
}

func NewServer(content embed.FS, config *http.Server) *Server {
    return &Server{
        FileSystem: content,
        Config: config,
        Err: nil,
    }
} 


// setStaticPath determines which folders inside the public folder will be included.
func (s *Server) setStaticPaths() {
    paths := []string{
        "/",
        "/static/",
        "/templates/",
    }
    s.handlePaths(paths)
}

func (s *Server) handlePaths(paths []string) {
    for _, path := range paths {
        staticFS, err := fs.Sub(s.FileSystem, filepath.Join("public", path))
        if err != nil {
            s.Err = err
            return
        }
        http.Handle(path, http.StripPrefix(path, http.FileServer(http.FS(staticFS))))
    }
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {    
    if r.URL.Path == "/cv" {
        handleIndex(w, r, s.FileSystem)
    }
}
