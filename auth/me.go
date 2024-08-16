package auth

import (
	"fmt"
	"hbd-cli/api"
	"hbd-cli/helper"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Me() *cobra.Command {
	var host, port, credsPath string
	var ssl, dotEnvFormat bool

	var meCmd = &cobra.Command{
		Use:   "me",
		Short: "Get authenticated user's data",
		Long: `The "me" command retrieves the authenticated user's data from the HBD service.
The data includes the Telegram bot API key, user ID, reminder time, and birthdays.

Environment variables:
  HBD_HOST - The host for the service. Defaults to 0.0.0.0.
  HBD_PORT - The port for the service.
  HBD_SSL - Use SSL (https) for the connection.

Example usage:
  hbd-cli auth me --host="hbd.lotiguere.com" --ssl --creds-path="~/.hbd/credentials" --dotenv
		`,
		Run: func(cmd *cobra.Command, args []string) {
			// Load env vars
			helper.LoadEnvVars()

			// Load credentials
			credsPath = filepath.Join(credsPath, host)
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

			// Create the URL
			url := helper.GenUrl(host, port, ssl)

			// Make the request to get user data
			userData, err := api.GetUserData(url, token)
			helper.HandleErrorExit("Error retrieving user data", err)

			// Print the retrieved user data
			// In dotenv format
			if dotEnvFormat {
				fmt.Printf("HBD_USER_ID=%d\n", userData.ID)
				fmt.Printf("HBD_TELEGRAM_BOT_API_KEY=%s\n", userData.TelegramBotAPIKey)
				fmt.Printf("HBD_TELEGRAM_USER_ID=%s\n", userData.TelegramUserID)
				fmt.Printf("HBD_REMINDER_TIME=%s\n", userData.ReminderTime)
				fmt.Printf("HBD_TIMEZONE=%s\n\n", userData.Timezone)
				fmt.Printf("# To view the birthdays use the 'birthdays' verb\n")
				return
			}
			
			// In regular format
			fmt.Printf("User Data:\n")
			fmt.Printf("ID: %d\n", userData.ID)
			fmt.Printf("Telegram Bot API Key: %s\n", userData.TelegramBotAPIKey)
			fmt.Printf("Telegram User ID: %s\n", userData.TelegramUserID)
			fmt.Printf("Reminder Time: %s\n", userData.ReminderTime)
			fmt.Printf("Timezone: %s\n", userData.Timezone)
			fmt.Printf("To view the birthdays use the 'birthdays' verb\n\n")
		},
	}

	// Add flags to the me command
	meCmd.Flags().StringVar(&host, "host", helper.DefaultHost(), "Host for the service")
	meCmd.Flags().StringVar(&port, "port", helper.DefaultPort(), "Port for the service")
	meCmd.Flags().BoolVar(&ssl, "ssl", helper.DefaultSSL(), "Use SSL (https) for the connection")
	meCmd.Flags().BoolVar(&dotEnvFormat, "dotenv", false, "Print the output in dotenv format (KEY=VALUE)")
	meCmd.Flags().StringVar(&credsPath, "creds-path", helper.GetDefaultCredsPath(), "Path to the credentials file")

	// Return the me command
	return meCmd
}
