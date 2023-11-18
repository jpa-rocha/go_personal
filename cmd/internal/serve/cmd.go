package serve

import (
	"embed"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// TODO: need to come from config
const timeOut = 10
const port = ":8000"

// TODO: is embeded needed is there a non global solution?
//go:embed public
var content embed.FS

func RunServe(_ *cobra.Command, _ []string) error {
	server := NewServer(
        content, 
        &http.Server{
		    Addr:         port,
		    ReadTimeout:  timeOut * time.Second,
	    	WriteTimeout: timeOut * time.Second,
	    },
    )

    server.setStaticPaths()
    if server.Err != nil {
        log.Printf("error: a problem occurred setting the file system: %s", server.Err.Error())
        return server.Err 
    }
    http.HandleFunc("/", server.ServeHTTP)
    log.Println("Server started at http://localhost" + port)
    
	log.Fatal(server.Config.ListenAndServe())

	return nil
}


