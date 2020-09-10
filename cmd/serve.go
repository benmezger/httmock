package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/benmezger/httmock/config"

	httmock_http "github.com/benmezger/httmock/http"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve API endpoints from http specification file",
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := rootCmd.Flags().GetString("api-file")
		host, _ := cmd.Flags().GetString("host")

		file := config.OpenFile(filename)
		spec := config.ReadHTTPSpec(file)
		router := httmock_http.SetupRoutes(spec)

		fmt.Printf("Running from '%s' file\n", filename)
		fmt.Printf("At host '%s'\n", host)

		log.Fatal(http.ListenAndServe(host, router))
	},
}

func init() {
	var host string
	serveCmd.PersistentFlags().StringVar(&host, "host", "localhost:8000",
		"ip address")
	rootCmd.AddCommand(serveCmd)
}
