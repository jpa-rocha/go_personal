package serve

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// TODO: need to come from config
const timeOut = 10
const port = ":8000"

//go:embed public
var content embed.FS

func RunServe(_ *cobra.Command, _ []string) error {
	server := &http.Server{
		Addr:         port,
		ReadTimeout:  timeOut * time.Second,
		WriteTimeout: timeOut * time.Second,
	}

	setStaticPaths()
    http.HandleFunc("/", handleRoutes)
    log.Println("Server started at http://localhost" + port)
    
	log.Fatal(server.ListenAndServe())

	return nil
}

// setStaticPath determines which folders inside the public folder will be included.
func setStaticPaths() {
    paths := []string{
        "/static/",
        "/templates",
    }
    handlePaths(paths) 
}

func handlePaths(paths []string) {
	staticFS, err := fs.Sub(content, "public")
	if err != nil {
		log.Fatal(err)
	}
    
    for _, path := range paths {
        http.Handle(path, http.StripPrefix(path, http.FileServer(http.FS(staticFS))))
    }
}


