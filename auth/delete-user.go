package auth

import (
	"bufio"
	"fmt"
	"hbd-cli/api"
	"hbd-cli/helper"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func DeleteUser() *cobra.Command {
	var host, port, credsPath string
	var ssl bool

	var deleteUserCmd = &cobra.Command{
		Use:   "delete-user",
		Short: "Delete your HBD account",
		Long: `The delete-user command will permanently delete your HBD account.
This action is irreversible, and all your data will be lost.

Environment variables:
  HBD_CREDS_PATH - Path to the credentials file.
  HBD_HOST - The host for the service. Defaults to
  HBD_PORT - The port for the service.
  HBD_SSL - Use SSL (https) for the connection.

Example usage:
  hbd-cli auth delete-user --host="hbd.lotiguere.com" --ssl --creds-path="~/.hbd/credentials"
		`,
		Run: func(cmd *cobra.Command, args []string) {
			// Load env vars
			helper.LoadEnvVars()

			// Load credentials
			credsPath = filepath.Join(credsPath, host)
			creds, err := helper.LoadCredentials(credsPath)
			helper.HandleError("Error loading credentials from credentials file", err)

			// Ask for confirmation
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Are you sure you want to delete your account? This action cannot be undone. Type 'yes' to proceed: ")
			confirmation, _ := reader.ReadString('\n')
			confirmation = strings.TrimSpace(confirmation)

			if confirmation != "yes" {
				fmt.Println("Account deletion aborted.")
				return
			}

			// Create the URL
			url := helper.GenUrl(host, port, ssl)

			// Make the request
			_, err = api.DeleteUser(url, creds.Token)
			helper.HandleErrorExit("Error deleting user", err)

			// Delete the credentials file
			if err := helper.DeleteCredentials(credsPath); err != nil {
				helper.HandleErrorExit("Error deleting credentials file", err)
			}

			// Print success message
			fmt.Println("Account deleted successfully!")
		},
	}

	// Add flags to the delete user command
	deleteUserCmd.Flags().StringVar(&host, "host", helper.DefaultHost(), "Host for the service")
	deleteUserCmd.Flags().StringVar(&port, "port", helper.DefaultPort(), "Port for the service")
	deleteUserCmd.Flags().BoolVar(&ssl, "ssl", helper.DefaultSSL(), "Use SSL (https) for the connection")
	deleteUserCmd.Flags().StringVar(&credsPath, "creds-path", helper.GetDefaultCredsPath(), "Path to the credentials file")

	// Return the delete user command
	return deleteUserCmd
}
