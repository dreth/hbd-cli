package auth

import (
	"fmt"
	"hbd-cli/api"
	"hbd-cli/helper"
	"hbd-cli/structs"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Login() *cobra.Command {
	var host string
	var port int
	var email string
	var password string
	var ssl bool
	var credsPath string

	var loginCmd = &cobra.Command{
		Use:   "login",
		Short: "Log in to your HBD account",
		Long: `The login command allows you to authenticate to your HBD account.
You can specify the host, port, email, and password either through
command-line flags or through environment variables.

Environment variables:
  HBD_EMAIL - The email address used for login.
  HBD_PASSWORD - The password used for login.

Example usage:
  hbd-cli login --email="user@example.com" --password="yourpassword" --host="example.com" --ssl --creds-path="~/.hbd/credentials"
		`,
		Run: func(cmd *cobra.Command, args []string) {
			// Read env vars
			viper.AutomaticEnv()

			// Load credentials
			creds, err := helper.LoadCredentials(credsPath)
			helper.HandleError("Error loading credentials from credentials file", err)

			// Check if email and password are an environment variable
			if email == "" {
				email = viper.GetString("HBD_EMAIL")
			}

			if password == "" {
				password = viper.GetString("HBD_PASSWORD")
			}

			// Check if email and password are empty
			if email == "" || password == "" {
				helper.HandleErrorExitStr("Error authenticating", "email and password must be provided either via flags or environment variables")
			}

			// Determine protocol based on ssl flag
			protocol := "http"
			if ssl {
				protocol = "https"
			}

			// Create the URL
			url := fmt.Sprintf("%s://%s:%d", protocol, host, port)

			// Create the JSON payload
			loginReq := structs.LoginRequest{
				Email:    email,
				Password: password,
			}

			// Make the request
			loginSuccess, err := api.Login(url, loginReq)
			helper.HandleErrorExit("Error logging in, wrong email or password", err)

			// Save the token to the credentials file
			creds.Token = loginSuccess.Token
			if err := helper.SaveCredentials(credsPath, creds); err != nil {
				helper.HandleErrorExit("Error saving credentials", err)
			}

			// Print success message
			fmt.Printf("Login successful! Token saved to %s\n", credsPath)
		},
	}

	// Add flags to the login command
	loginCmd.Flags().StringVar(&host, "host", "0.0.0.0", "Host for the service")
	loginCmd.Flags().IntVar(&port, "port", 8417, "Port for the service")
	loginCmd.Flags().StringVar(&email, "email", "", "Email for login")
	loginCmd.Flags().StringVar(&password, "password", "", "Password for login")
	loginCmd.Flags().BoolVar(&ssl, "ssl", false, "Use SSL (https) for the connection")
	loginCmd.Flags().StringVar(&credsPath, "creds-path", helper.GetDefaultCredsPath(), "Path to the credentials file")

	// Return the login command
	return loginCmd
}
