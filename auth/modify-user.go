package auth

import (
	"fmt"
	"hbd-cli/api"
	"hbd-cli/helper"
	"hbd-cli/structs"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ModifyUser() *cobra.Command {
	var host, port, newEmail, newPassword, newReminderTime, newTimezone, newTelegramBotAPIKey, newTelegramUserID, credsPath string
	var tokenDuration int
	var ssl bool

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
  HBD_HOST - The host for the service. Defaults to 0.0.0.0.
  HBD_PORT - The port for the service.
  HBD_SSL - Use SSL (https) for the connection.

Example usage:
  hbd-cli auth modify-user --new-email="newuser@hbd.lotiguere.com" --new-password="newpassword" --new-reminder-time="15:04" --new-timezone="America/New_York" --new-telegram-bot-api-key="your-new-bot-api-key" --new-telegram-user-id="your-new-user-id"
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
				helper.HandleErrorExitStr("Error", "A valid JWT token must be provided either via the credentials file")
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

			// Create the URL
			url := helper.GenUrl(host, port, ssl)

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
			// If the user provides a new email OR password, call the ModifyUserWithEmail endpoint
			if newEmail != "" || newPassword != "" {
				modifyUserReq := structs.ModifyUserRequest{
					NewEmail:             newEmail,
					NewPassword:          newPassword,
					NewReminderTime:      newReminderTime,
					NewTimezone:          newTimezone,
					NewTelegramBotAPIKey: newTelegramBotAPIKey,
					NewTelegramUserID:    newTelegramUserID,
				}

				// Make the request to modify user details
				success, err := api.ModifyUserWithEmail(url, token, modifyUserReq, tokenDuration)
				helper.HandleErrorExit("Error modifying user details", err)

				// Save the new token to the credentials file
				creds.Token = success.Token
				if err := helper.SaveCredentials(credsPath, creds); err != nil {
					helper.HandleErrorExit("Error saving credentials", err)
				}

				// Print success message
				fmt.Printf("User details modified successfully! Token saved to %s\n", credsPath)

				// If the user does not provide a new email OR password, call the ModifyUserWithoutEmail endpoint
			} else if newEmail == "" && newPassword == "" {
				modifyUserReq := structs.ModifyUserRequest{
					NewEmail:             "",
					NewPassword:          "",
					NewReminderTime:      newReminderTime,
					NewTimezone:          newTimezone,
					NewTelegramBotAPIKey: newTelegramBotAPIKey,
					NewTelegramUserID:    newTelegramUserID,
				}

				// Make the request to modify user details
				userData, err := api.ModifyUserWithoutEmail(url, token, modifyUserReq)
				helper.HandleErrorExit("Error modifying user details", err)

				// Print the modified user data
				fmt.Printf("User data modified successfully!\n\n")
				fmt.Printf("User Data:\n")
				fmt.Printf("ID: %d\n", userData.ID)
				fmt.Printf("Telegram Bot API Key: %s\n", userData.TelegramBotAPIKey)
				fmt.Printf("Telegram User ID: %s\n", userData.TelegramUserID)
				fmt.Printf("Reminder Time: %s\n", userData.ReminderTime)
				fmt.Printf("Timezone: %s\n", userData.Timezone)
				fmt.Printf("To view the birthdays use the 'birthdays' verb\n\n")

			}
		},
	}

	// Add flags to the modify-user command
	modifyUserCmd.Flags().StringVar(&host, "host", helper.DefaultHost(), "Host for the service")
	modifyUserCmd.Flags().StringVar(&port, "port", helper.DefaultPort(), "Port for the service")
	modifyUserCmd.Flags().StringVar(&newEmail, "new-email", "", "New email for the user")
	modifyUserCmd.Flags().StringVar(&newPassword, "new-password", "", "New password for the user")
	modifyUserCmd.Flags().StringVar(&newReminderTime, "new-reminder-time", "", "New reminder time (HH:MM) for the user")
	modifyUserCmd.Flags().StringVar(&newTimezone, "new-timezone", "", "New timezone for the reminder")
	modifyUserCmd.Flags().StringVar(&newTelegramBotAPIKey, "new-telegram-bot-api-key", "", "New Telegram bot API key for the user")
	modifyUserCmd.Flags().StringVar(&newTelegramUserID, "new-telegram-user-id", "", "New Telegram user ID for the user")
	modifyUserCmd.Flags().BoolVar(&ssl, "ssl", helper.DefaultSSL(), "Use SSL (https) for the connection")
	modifyUserCmd.Flags().StringVar(&credsPath, "creds-path", helper.GetDefaultCredsPath(), "Path to the credentials file")
	modifyUserCmd.Flags().IntVar(&tokenDuration, "token-duration", 720, "Duration of the JWT token in hours. Default is 720 hours (30 days), in case of reauth due to modified email or password.")

	// Return the modify-user command
	return modifyUserCmd
}
