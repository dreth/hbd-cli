package auth

import (
	"fmt"
	"hbd-cli/api"
	"hbd-cli/helper"
	"hbd-cli/structs"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Register() *cobra.Command {
	var host string
	var port int
	var email string
	var password string
	var reminderTime string
	var timezone string
	var telegramBotAPIKey string
	var telegramUserID string
	var ssl bool
	var credsPath string

	var registerCmd = &cobra.Command{
		Use:   "register",
		Short: "Register a new HBD account",
		Long: `The register command allows you to create a new HBD account.
You can specify the host, port, and all required user details either through
command-line flags or through environment variables.

Environment variables:
  HBD_EMAIL - The email address used for registration.
  HBD_PASSWORD - The password used for registration.
  HBD_REMINDER_TIME - The reminder time (in HH:MM format).
  HBD_TIMEZONE - The timezone for the reminder.
  HBD_TELEGRAM_BOT_API_KEY - The Telegram bot API key.
  HBD_TELEGRAM_USER_ID - The Telegram user ID.

Example usage:
  hbd-cli register --email="user@example.com" --password="yourpassword" --reminder-time="15:04" --timezone="America/New_York" --telegram-bot-api-key="your-bot-api-key" --telegram-user-id="your-user-id"
		`,
		Run: func(cmd *cobra.Command, args []string) {
			// Read env vars
			viper.AutomaticEnv()

			// Load credentials
			creds, err := helper.LoadCredentials(credsPath)
			helper.HandleError("Error loading credentials from credentials file", err)

			// Check if user details are provided via environment variables
			if email == "" {
				email = viper.GetString("HBD_EMAIL")
			}
			if password == "" {
				password = viper.GetString("HBD_PASSWORD")
			}
			if reminderTime == "" {
				reminderTime = viper.GetString("HBD_REMINDER_TIME")
			}
			if timezone == "" {
				timezone = viper.GetString("HBD_TIMEZONE")
			}
			if telegramBotAPIKey == "" {
				telegramBotAPIKey = viper.GetString("HBD_TELEGRAM_BOT_API_KEY")
			}
			if telegramUserID == "" {
				telegramUserID = viper.GetString("HBD_TELEGRAM_USER_ID")
			}

			// Check if any required details are empty
			if email == "" || password == "" || reminderTime == "" || timezone == "" || telegramBotAPIKey == "" || telegramUserID == "" {
				helper.HandleErrorExitStr("Error registering", "All registration details must be provided either via flags or environment variables")
			}

			// Determine protocol based on ssl flag
			protocol := "http"
			if ssl {
				protocol = "https"
			}

			// Create the URL
			url := fmt.Sprintf("%s://%s:%d", protocol, host, port)

			// Create the JSON payload for registration
			registerReq := structs.RegisterRequest{
				Email:             email,
				Password:          password,
				ReminderTime:      reminderTime,
				Timezone:          timezone,
				TelegramBotAPIKey: telegramBotAPIKey,
				TelegramUserID:    telegramUserID,
			}

			// Make the registration request
			loginSuccess, err := api.Register(url, registerReq)
			helper.HandleErrorExit("Error registering user", err)

			// Save the token to the credentials file
			creds.Token = loginSuccess.Token
			if err := helper.SaveCredentials(credsPath, creds); err != nil {
				helper.HandleErrorExit("Error saving credentials", err)
			}

			// Print success message
			fmt.Printf("Registration successful! Token saved to %s\n", credsPath)
		},
	}

	// Add flags to the register command
	registerCmd.Flags().StringVar(&host, "host", "0.0.0.0", "Host for the service")
	registerCmd.Flags().IntVar(&port, "port", 8417, "Port for the service")
	registerCmd.Flags().StringVar(&email, "email", "", "Email for registration")
	registerCmd.Flags().StringVar(&password, "password", "", "Password for registration")
	registerCmd.Flags().StringVar(&reminderTime, "reminder-time", "", "Reminder time (HH:MM) for registration")
	registerCmd.Flags().StringVar(&timezone, "timezone", "", "Timezone for the reminder")
	registerCmd.Flags().StringVar(&telegramBotAPIKey, "telegram-bot-api-key", "", "Telegram bot API key for registration")
	registerCmd.Flags().StringVar(&telegramUserID, "telegram-user-id", "", "Telegram user ID for registration")
	registerCmd.Flags().BoolVar(&ssl, "ssl", false, "Use SSL (https) for the connection")
	registerCmd.Flags().StringVar(&credsPath, "creds-path", helper.GetDefaultCredsPath(), "Path to the credentials file")

	// Return the register command
	return registerCmd
}
