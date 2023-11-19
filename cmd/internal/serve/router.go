package serve
import (
	"embed"
	"io/fs"
    "net/http"
	"path/filepath"
)

type Router struct {
    Routes      []string
    FileSystem  embed.FS
    Err         error

}

func NewRouter(content embed.FS, routes []string) *Router {
    return &Router{
        FileSystem: content,
        Routes: routes,
    }
}

func (rt *Router) setStaticPaths() {
    for _, path := range rt.Routes {
        staticFS, err := fs.Sub(rt.FileSystem, filepath.Join("public", path))
        if err != nil {
            rt.Err = err
            return
        }
        http.Handle(path, http.StripPrefix(path, http.FileServer(http.FS(staticFS))))
    }
}

func (rt *Router) handleRoutes(w http.ResponseWriter, r *http.Request) {    
    if r.URL.Path == "/cv" {
        handleCV(w, r, rt.FileSystem)
    }
}
