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

func Login() *cobra.Command {
	var host, port, email, password, credsPath string
	var tokenDuration int
	var ssl bool

	var loginCmd = &cobra.Command{
		Use:   "login",
		Short: "Log in to your HBD account",
		Long: `The login command allows you to authenticate to your HBD account.
You can specify the host, port, email, and password either through
command-line flags or through environment variables.

Environment variables:
  HBD_EMAIL - The email address used for login.
  HBD_PASSWORD - The password used for login.
  HBD_HOST - The host for the service. Defaults to 0.0.0.0.
  HBD_PORT - The port for the service. 
  HBD_SSL - Use SSL (https) for the connection.
  HBD_CREDS_PATH - Path to the credentials file.

Example usage:
  hbd-cli auth login --email="user@hbd.lotiguere.com" --password="yourpassword" --host="hbd.lotiguere.com" --ssl --creds-path="~/.hbd/credentials" --token-duration=3600
		`,
		Run: func(cmd *cobra.Command, args []string) {
			// Load env vars
			helper.LoadEnvVars()

			// Load credentials
			credsPath = filepath.Join(credsPath, host)
			creds, err := helper.LoadCredentials(credsPath)
			if err == nil {
				fmt.Println("Credentials exist, they will be overwritten.")
			}

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

			// Create the URL
			url := helper.GenUrl(host, port, ssl)

			// Create the JSON payload
			loginReq := structs.LoginRequest{
				Email:    email,
				Password: password,
			}

			// Make the request
			loginSuccess, err := api.Login(url, loginReq, tokenDuration)
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
	loginCmd.Flags().StringVar(&host, "host", helper.DefaultHost(), "Host for the service")
	loginCmd.Flags().StringVar(&port, "port", helper.DefaultPort(), "Port for the service")
	loginCmd.Flags().StringVar(&email, "email", "", "Email for login")
	loginCmd.Flags().StringVar(&password, "password", "", "Password for login")
	loginCmd.Flags().BoolVar(&ssl, "ssl", helper.DefaultSSL(), "Use SSL (https) for the connection")
	loginCmd.Flags().StringVar(&credsPath, "creds-path", helper.GetDefaultCredsPath(), "Path to the credentials file")
	loginCmd.Flags().IntVar(&tokenDuration, "token-duration", 720, "Duration of the JWT token in hours. Default is 720 hours (30 days).")

	// Return the login command
	return loginCmd
}
