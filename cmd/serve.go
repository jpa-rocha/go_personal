package cmd

import (
	"github.com/spf13/cobra"

	"adamastor/internal/serve"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve your files",
	RunE:  serve.RunServe,
}

func init() {
	RootCmd.AddCommand(ServeCmd)
}
