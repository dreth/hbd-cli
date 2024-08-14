package auth

import (
	"fmt"
	"hbd-cli/api"
	"hbd-cli/helper"
	"hbd-cli/structs"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ModifyUser() *cobra.Command {
	var host string
	var port int
	var newEmail string
	var newPassword string
	var newReminderTime string
	var newTimezone string
	var newTelegramBotAPIKey string
	var newTelegramUserID string
	var ssl bool
	var credsPath string

	var modifyUserCmd = &cobra.Command{
		Use:   "modify-user",
		Short: "Modify user's details",
		Long: `The modify-user command allows you to modify your HBD account details.
You can specify the new email, password, reminder time, timezone, Telegram bot API key, 
and Telegram user ID either through command-line flags or through environment variables.

Environment variables:
  HBD_NEW_EMAIL - The new email address for the user.
  HBD_NEW_PASSWORD - The new password for the user.
  HBD_NEW_REMINDER_TIME - The new reminder time (in HH:MM format).
  HBD_NEW_TIMEZONE - The new timezone for the reminder.
  HBD_NEW_TELEGRAM_BOT_API_KEY - The new Telegram bot API key.
  HBD_NEW_TELEGRAM_USER_ID - The new Telegram user ID.

Example usage:
  hbd-cli modify-user --new-email="newuser@example.com" --new-password="newpassword" --new-reminder-time="15:04" --new-timezone="America/New_York" --new-telegram-bot-api-key="your-new-bot-api-key" --new-telegram-user-id="your-new-user-id"
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

			// Check if user details are provided via environment variables
			if newEmail == "" {
				newEmail = viper.GetString("HBD_NEW_EMAIL")
			}
			if newPassword == "" {
				newPassword = viper.GetString("HBD_NEW_PASSWORD")
			}
			if newReminderTime == "" {
				newReminderTime = viper.GetString("HBD_NEW_REMINDER_TIME")
			}
			if newTimezone == "" {
				newTimezone = viper.GetString("HBD_NEW_TIMEZONE")
			}
			if newTelegramBotAPIKey == "" {
				newTelegramBotAPIKey = viper.GetString("HBD_NEW_TELEGRAM_BOT_API_KEY")
			}
			if newTelegramUserID == "" {
				newTelegramUserID = viper.GetString("HBD_NEW_TELEGRAM_USER_ID")
			}

			// Determine protocol based on ssl flag
			protocol := "http"
			if ssl {
				protocol = "https"
			}

			// Create the URL
			url := fmt.Sprintf("%s://%s:%d", protocol, host, port)

			// Get the user's existing data by calling the Me endpoint and fill up the missing fields with the existing data
			userData, err := api.GetUserData(url, token)
			helper.HandleErrorExit("Error retrieving user data", err)
			if newReminderTime == "" {
				newReminderTime = userData.ReminderTime
			}
			if newTimezone == "" {
				newTimezone = userData.Timezone
			}
			if newTelegramBotAPIKey == "" {
				newTelegramBotAPIKey = userData.TelegramBotAPIKey
			}
			if newTelegramUserID == "" {
				newTelegramUserID = userData.TelegramUserID
			}

			// Create the JSON payload for modifying user details
			modifyUserReq := structs.ModifyUserRequest{
				NewEmail:             newEmail,
				NewPassword:          newPassword,
				NewReminderTime:      newReminderTime,
				NewTimezone:          newTimezone,
				NewTelegramBotAPIKey: newTelegramBotAPIKey,
				NewTelegramUserID:    newTelegramUserID,
			}

			// Make the request to modify user details
			success, err := api.ModifyUser(url, token, modifyUserReq)
			helper.HandleErrorExit("Error modifying user details", err)

			// Print success message
			if success.Success {
				fmt.Println("User details modified successfully!")
			} else {
				fmt.Println("Failed to modify user details")
			}
		},
	}

	// Add flags to the modify-user command
	modifyUserCmd.Flags().StringVar(&host, "host", "0.0.0.0", "Host for the service")
	modifyUserCmd.Flags().IntVar(&port, "port", 8417, "Port for the service")
	modifyUserCmd.Flags().StringVar(&newEmail, "new-email", "", "New email for the user")
	modifyUserCmd.Flags().StringVar(&newPassword, "new-password", "", "New password for the user")
	modifyUserCmd.Flags().StringVar(&newReminderTime, "new-reminder-time", "", "New reminder time (HH:MM) for the user")
	modifyUserCmd.Flags().StringVar(&newTimezone, "new-timezone", "", "New timezone for the reminder")
	modifyUserCmd.Flags().StringVar(&newTelegramBotAPIKey, "new-telegram-bot-api-key", "", "New Telegram bot API key for the user")
	modifyUserCmd.Flags().StringVar(&newTelegramUserID, "new-telegram-user-id", "", "New Telegram user ID for the user")
	modifyUserCmd.Flags().BoolVar(&ssl, "ssl", false, "Use SSL (https) for the connection")
	modifyUserCmd.Flags().StringVar(&credsPath, "creds-path", helper.GetDefaultCredsPath(), "Path to the credentials file")

	// Return the modify-user command
	return modifyUserCmd
}
