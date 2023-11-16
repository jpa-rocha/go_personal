package cmd

import (
    "adamastor/cmd/internal/serve"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command {
    Use: "serve",
    Short: "Serve your files",
    RunE: serve.RunServe,
}

func init() {
    RootCmd.AddCommand(ServeCmd)
}
