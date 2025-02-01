package cmd

import (
	"github.com/krau/shisoimg/api"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Run:   api.Serve,
}

func init() {
	serveCmd.Flags().StringP("host", "a", ":34180", "host to listen on")
	rootCmd.AddCommand(serveCmd)
}
