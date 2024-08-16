package birthdays

import (
	"fmt"
	"hbd-cli/api"
	"hbd-cli/helper"
	"path/filepath"

	"github.com/spf13/cobra"
)

// CheckBirthdays command
func CheckBirthdays() *cobra.Command {
	var host, port, credsPath string
	var ssl bool

	var checkBirthdaysCmd = &cobra.Command{
		Use:   "check",
		Short: "Check birthdays for reminders",
		Long: `This command force checks if there's any birthday today.

Environment variables:
  HBD_CREDS_PATH - Path to the credentials file.
  HBD_HOST - The host for the service. Defaults to 0.0.0.0.
  HBD_PORT - The port for the service.
  HBD_SSL - Use SSL (https) for the connection.

Example usage:
  hbd-cli birthdays check --host="hbd.lotiguere.com" --ssl --creds-path="~/.hbd/credentials"
`,
		Run: func(cmd *cobra.Command, args []string) {
			// Load env vars
			helper.LoadEnvVars()

			// Load credentials
			credsPath = filepath.Join(credsPath, host)
			creds, err := helper.LoadCredentials(credsPath)
			helper.HandleError("Error loading credentials from credentials file", err)

			// Create the URL
			url := helper.GenUrl(host, port, ssl)

			// Make the request to check birthdays
			_, err = api.CheckBirthdays(url, creds.Token)
			helper.HandleErrorExit("Error checking birthdays", err)

			// Print the result
			fmt.Printf("Check performed, if there's a birthday today, you should receive a message.")
		},
	}

	// Add flags
	checkBirthdaysCmd.Flags().StringVar(&host, "host", helper.DefaultHost(), "Host for the service")
	checkBirthdaysCmd.Flags().StringVar(&port, "port", helper.DefaultPort(), "Port for the service")
	checkBirthdaysCmd.Flags().BoolVar(&ssl, "ssl", helper.DefaultSSL(), "Use SSL (https) for the connection")
	checkBirthdaysCmd.Flags().StringVar(&credsPath, "creds-path", helper.GetDefaultCredsPath(), "Path to the credentials file")

	return checkBirthdaysCmd
}
