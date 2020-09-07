package cmd

import (
	"fmt"
	"httmock/config"
	"log"
	"net/http"

	httmock_http "httmock/http"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve API endpoints from http specification file",
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := rootCmd.Flags().GetString("api-file")

		fmt.Printf("Running from '%s' file\n", filename)
		file := config.OpenFile(filename)
		spec := config.ReadHTTPSpec(file)
		router := httmock_http.SetupRoutes(spec)
		log.Fatal(http.ListenAndServe(":8000", router))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
