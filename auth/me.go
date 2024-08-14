package auth

import (
	"fmt"
	"hbd-cli/api"
	"hbd-cli/helper"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Me() *cobra.Command {
	var host string
	var port int
	var ssl bool
	var credsPath string

	var meCmd = &cobra.Command{
		Use:   "me",
		Short: "Get authenticated user's data",
		Long: `The "me" command retrieves the authenticated user's data from the HBD service.
The data includes the Telegram bot API key, user ID, reminder time, and birthdays.
A valid JWT token must be provided either through a credentials file or the environment variable HBD_TOKEN.

Environment variables:
  HBD_TOKEN - The JWT token used for authentication.

Example usage:
  hbd-cli me --host="example.com" --ssl --creds-path="~/.hbd/credentials"
		`,
		Run: func(cmd *cobra.Command, args []string) {
			// Load credentials
			creds, err := helper.LoadCredentials(credsPath)
			helper.HandleError("Error loading credentials from credentials file", err)

			// Check if the token is provided via environment variable
			token := creds.Token
			if token == "" {
				token = viper.GetString("HBD_TOKEN")
			}

			// Check if the token is still empty
			if token == "" {
				helper.HandleErrorExitStr("Error", "A valid JWT token must be provided either via the credentials file or the environment variable HBD_TOKEN")
			}

			// Determine protocol based on ssl flag
			protocol := "http"
			if ssl {
				protocol = "https"
			}

			// Create the URL
			url := fmt.Sprintf("%s://%s:%d", protocol, host, port)

			// Make the request to get user data
			userData, err := api.GetUserData(url, token)
			helper.HandleErrorExit("Error retrieving user data", err)

			// Print the retrieved user data
			fmt.Printf("User Data:\n")
			fmt.Printf("ID: %d\n", userData.ID)
			fmt.Printf("Telegram Bot API Key: %s\n", userData.TelegramBotAPIKey)
			fmt.Printf("Telegram User ID: %s\n", userData.TelegramUserID)
			fmt.Printf("Reminder Time: %s\n", userData.ReminderTime)
			fmt.Printf("Timezone: %s\n", userData.Timezone)
			fmt.Printf("To view the birthdays use the 'birthdays' verb\n")
		},
	}

	// Add flags to the me command
	meCmd.Flags().StringVar(&host, "host", "0.0.0.0", "Host for the service")
	meCmd.Flags().IntVar(&port, "port", 8417, "Port for the service")
	meCmd.Flags().BoolVar(&ssl, "ssl", false, "Use SSL (https) for the connection")
	meCmd.Flags().StringVar(&credsPath, "creds-path", helper.GetDefaultCredsPath(), "Path to the credentials file")

	// Return the me command
	return meCmd
}
