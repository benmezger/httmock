package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var apiFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "httmock",
	Short: "Mock a HTTP server from a http specification",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&apiFile, "api-file", ".httmock.yaml",
		"api file (default is .httmock.yaml)")
}
