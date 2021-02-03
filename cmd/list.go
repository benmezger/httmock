package cmd

import (
	"fmt"
	"strings"

	"github.com/benmezger/httmock/config"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List API endpoints from http specification file",
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := rootCmd.Flags().GetString("api-file")

		file := config.OpenFile(filename)
		spec := config.ReadHTTPSpec(file)

		for path, attrs := range spec.Paths {
			for method := range attrs {
				fmt.Println(strings.ToUpper(method), path)
				params := spec.Paths[path][method].Request.Params
				if len(params) > 0 {

					fmt.Println("Params:")
					for param := range params {
						fmt.Println("\t", param, "=", params[param])
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
