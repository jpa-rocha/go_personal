package serve

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func RunServe(_ *cobra.Command, _ []string) error {
	server := NewServer()

	if server.Err != nil {
		err := fmt.Errorf("error: a problem occurred setting the file system: %w", server.Err)
		return err
	}
	log.Println("Server started at http://localhost" + server.Config.Addr)

	log.Fatal(server.Config.ListenAndServe())

	return nil
}
