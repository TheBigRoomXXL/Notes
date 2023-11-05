package cli

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "notes-server",
	Args: cobra.ArbitraryArgs,
	RunE: serve,
}

func init() {
	RootCmd.AddCommand(getServeCmd())
	RootCmd.AddCommand(getUsersCmd())
}

func FuckThat() {
	RootCmd.Execute()
}
