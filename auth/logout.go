package auth

import (
	"bufio"
	"fmt"
	"hbd-cli/helper"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func Logout() *cobra.Command {
	var credsPath, host string
	var autoConfirm bool

	var logoutCmd = &cobra.Command{
		Use:   "logout",
		Short: "Log out from your HBD account (clean up saved token)",
		Long: `Logging out cleans up the saved token from the credentials file in the designated credentials path (default is ~/.hbd/credentials/<host>) for a given HBD host.

Environment variables:
  HBD_CREDS_PATH - Path to the credentials file.
  HBD_HOST - The host for the service. Defaults to

Example usage:
  hbd-cli auth logout --host="hbd.lotiguere.com" --creds-path="~/.hbd/credentials" -y
		`,
		Run: func(cmd *cobra.Command, args []string) {
			// Load env vars
			helper.LoadEnvVars()

			// Load credentials
			credsPath = filepath.Join(credsPath, host)
			_, err := helper.LoadCredentials(credsPath)
			helper.HandleErrorExit("Error loading credentials from credentials file, are you sure you've logged in before?", err)

			// If autoConfirm is not set, ask for confirmation
			if !autoConfirm {
				reader := bufio.NewReader(os.Stdin)
				fmt.Printf("Are you sure you want to log out? This will clear the credentials in %s. Type 'yes' to proceed: ", credsPath)
				confirmation, _ := reader.ReadString('\n')
				confirmation = strings.TrimSpace(confirmation)

				if confirmation != "yes" {
					fmt.Println("Logout cancelled.")
					return
				}
			}

			// Delete the credentials file
			if err := helper.DeleteCredentials(credsPath); err != nil {
				helper.HandleErrorExit("Error deleting credentials file", err)
			}

			// Print success message
			fmt.Println("Logged out successfully.")
		},
	}

	// Add flags to the logout command
	logoutCmd.Flags().StringVar(&host, "host", helper.DefaultHost(), "Host for the service")
	logoutCmd.Flags().StringVar(&credsPath, "creds-path", helper.GetDefaultCredsPath(), "Path to the credentials file")
	logoutCmd.Flags().BoolVarP(&autoConfirm, "yes", "y", false, "Automatic yes to confirmation prompt")

	// Return the logout command
	return logoutCmd
}
