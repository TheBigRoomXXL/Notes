package cli

import (
	"fmt"
	"notes/app/shared"
	"notes/app/users"

	"github.com/spf13/cobra"
)

func getUsersCmd() *cobra.Command {
	userCmd := &cobra.Command{
		Use:   "users",
		Short: "Administrate users",
	}

	userCreateCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new user",
		Args:  cobra.ExactArgs(2),
		RunE:  createUser,
	}

	userCmd.AddCommand(userCreateCmd)

	return userCmd
}

func createUser(cmd *cobra.Command, args []string) error {
	db := shared.CreateDbConnection("notes.db")
	user := users.UserSerializer{
		Username: args[0],
		Password: args[1],
	}

	_, err := users.InsertUser(db, user)
	if err != nil {
		return err
	}

	fmt.Printf("User %s successfully created\n", user.Username)
	return nil
}
