package cli

import (
	"notes/app/server"

	"github.com/spf13/cobra"
)

func getServeCmd() *cobra.Command {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Start the server",
		RunE:  serve,
	}
	return serveCmd
}

func serve(cmd *cobra.Command, args []string) error {
	server.Run()
	return nil
}
