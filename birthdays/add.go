package birthdays

import (
	"fmt"
	"hbd-cli/api"
	"hbd-cli/helper"
	"hbd-cli/structs"
	"path/filepath"

	"github.com/spf13/cobra"
)

// AddBirthday command
func AddBirthday() *cobra.Command {
	var name, date, host, port, credsPath string
	var ssl bool

	var addBirthdayCmd = &cobra.Command{
		Use:   "add",
		Short: "Add a new birthday",
		Long:  `The add-birthday command allows you to add a new birthday to your account.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Load env vars
			helper.LoadEnvVars()

			// Load credentials
			credsPath = filepath.Join(credsPath, host)
			creds, err := helper.LoadCredentials(credsPath)
			helper.HandleError("Error loading credentials from credentials file", err)

			// Create the URL
			url := helper.GenUrl(host, port, ssl)

			// Create the JSON payload
			birthdayReq := structs.BirthdayNameDateAdd{
				Name: name,
				Date: date,
			}

			// Make the request
			_, err = api.AddBirthday(url, creds.Token, birthdayReq)
			helper.HandleErrorExit("Error adding birthday", err)

			// Print success message
			fmt.Printf("Birthday for %s on %s added successfully!\n", name, date)
		},
	}

	// Add flags
	addBirthdayCmd.Flags().StringVar(&name, "name", "", "Name of the person (required)")
	addBirthdayCmd.Flags().StringVar(&date, "date", "", "Date of the birthday (required, format YYYY-MM-DD)")
	addBirthdayCmd.Flags().StringVar(&host, "host", helper.DefaultHost(), "Host for the service")
	addBirthdayCmd.Flags().StringVar(&port, "port", helper.DefaultPort(), "Port for the service")
	addBirthdayCmd.Flags().BoolVar(&ssl, "ssl", helper.DefaultSSL(), "Use SSL (https) for the connection")
	addBirthdayCmd.Flags().StringVar(&credsPath, "creds-path", helper.GetDefaultCredsPath(), "Path to the credentials file")

	// Mark required flags
	addBirthdayCmd.MarkFlagRequired("name")
	addBirthdayCmd.MarkFlagRequired("date")

	return addBirthdayCmd
}
