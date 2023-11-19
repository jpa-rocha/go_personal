package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

var RootCmd = &cobra.Command {
    Use:    "Adamastor",
    Short:  "Go Web Server",
    Long:   "Very Large and Scary Web Server",

    Run: func (cmd *cobra.Command, args []string) {

    },
}

func Execute() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}

// func initConfig(validate bool) {
//     if cfgfile != "" {
//         viper.SetConfigFile(cfgFile)
//         parseConfig(validate)
//     }
// }
//
// func ParseConfig(val bool) {
//     
// }