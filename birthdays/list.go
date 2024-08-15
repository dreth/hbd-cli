package birthdays

import (
	"fmt"
	"hbd-cli/api"
	"hbd-cli/helper"
	"path/filepath"

	"github.com/spf13/cobra"
)

// ListBirthdays command
func ListBirthdays() *cobra.Command {
	var host, port, credsPath string
	var ssl bool

	var listBirthdaysCmd = &cobra.Command{
		Use:   "list",
		Short: "List all birthdays",
		Long:  `The list-birthdays command retrieves and displays all the birthdays associated with your account.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Load env vars
			helper.LoadEnvVars()

			// Load credentials
			credsPath = filepath.Join(credsPath, host)
			creds, err := helper.LoadCredentials(credsPath)
			helper.HandleError("Error loading credentials from credentials file", err)

			// Create the URL
			url := helper.GenUrl(host, port, ssl)

			// Make the request to get user data
			userData, err := api.GetUserData(url, creds.Token)
			helper.HandleErrorExit("Error retrieving user data", err)

			// Print the birthdays
			fmt.Println("Your Birthdays:")
			for _, birthday := range userData.Birthdays {
				fmt.Printf("ID: %d, Name: %s, Date: %s\n", birthday.ID, birthday.Name, birthday.Date)
			}
		},
	}

	// Add flags
	listBirthdaysCmd.Flags().StringVar(&host, "host", helper.DefaultHost(), "Host for the service")
	listBirthdaysCmd.Flags().StringVar(&port, "port", helper.DefaultPort(), "Port for the service")
	listBirthdaysCmd.Flags().BoolVar(&ssl, "ssl", helper.DefaultSSL(), "Use SSL (https) for the connection")
	listBirthdaysCmd.Flags().StringVar(&credsPath, "creds-path", helper.GetDefaultCredsPath(), "Path to the credentials file")

	return listBirthdaysCmd
}
