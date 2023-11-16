package serve

import (
    "net/http"
    "time"
    "log"
    "github.com/spf13/cobra"
)

const timeOut = 10
const port = ":8000"
const dirname = "./"

func RunServe(_ *cobra.Command, _ []string) error {

	// setServePath()
    http.Handle("/", http.FileServer(http.Dir(dirname)))

	log.Println("Server started at http://localhost" + port)

	s := &http.Server{
		Addr:         port,
		ReadTimeout:  timeOut * time.Second,
		WriteTimeout: timeOut * time.Second,
	}

	log.Fatal(s.ListenAndServe())

	return nil
}


