package birthdays

import (
	"fmt"
	"hbd-cli/api"
	"hbd-cli/helper"
	"hbd-cli/structs"
	"path/filepath"

	"github.com/spf13/cobra"
)

// ModifyBirthday command
func ModifyBirthday() *cobra.Command {
	var id int64
	var name, date, host, port, credsPath string
	var ssl bool

	var modifyBirthdayCmd = &cobra.Command{
		Use:   "modify",
		Short: "Modify an existing birthday",
		Long:  `The modify-birthday command allows you to modify an existing birthday by providing its ID, new name, and/or new date. The ID is required, but if the name or date is not provided, the existing value will be used.
		
Environment variables:
  HBD_CREDS_PATH - Path to the credentials file.
  HBD_HOST - The host for the service. Defaults to
  HBD_PORT - The port for the service.
  HBD_SSL - Use SSL (https) for the connection.

Example usage:
  hbd-cli birthdays modify --id=1 --name="John Doe" --date="2021-12-25" --host="hbd.lotiguere.com" --ssl --creds-path="~/.hbd/credentials"
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

			// If the ID is sent, but NOT the name or date, just look them up requesting /me
			if id != 0 && (name == "" || date == "") {
				// Get the user data
				userData, err := api.GetUserData(url, creds.Token)
				helper.HandleErrorExit("Error retrieving user data", err)

				// Find the birthday
				for _, birthday := range userData.Birthdays {
					if birthday.ID == id {

						// If the name is empty, use the existing name
						if name == "" {
							name = birthday.Name
						}

						// If the date is empty, use the existing date
						if date == "" {
							date = birthday.Date
						}

						// Exit the loop
						break
					}
				}
			}

			// Create the JSON payload
			birthdayReq := structs.BirthdayNameDateModify{
				ID:   id,
				Name: name,
				Date: date,
			}

			// Make the request
			_, err = api.ModifyBirthday(url, creds.Token, birthdayReq)
			helper.HandleErrorExit("Error modifying birthday", err)

			// Print success message
			fmt.Printf("Birthday with ID %d modified successfully to %s on %s!\n", id, name, date)
		},
	}

	// Add flags
	modifyBirthdayCmd.Flags().Int64Var(&id, "id", 0, "ID of the birthday to modify (required)")
	modifyBirthdayCmd.Flags().StringVar(&name, "name", "", "New name for the birthday")
	modifyBirthdayCmd.Flags().StringVar(&date, "date", "", "New date for the birthday (format YYYY-MM-DD)")
	modifyBirthdayCmd.Flags().StringVar(&host, "host", helper.DefaultHost(), "Host for the service")
	modifyBirthdayCmd.Flags().StringVar(&port, "port", helper.DefaultPort(), "Port for the service")
	modifyBirthdayCmd.Flags().BoolVar(&ssl, "ssl", helper.DefaultSSL(), "Use SSL (https) for the connection")
	modifyBirthdayCmd.Flags().StringVar(&credsPath, "creds-path", helper.GetDefaultCredsPath(), "Path to the credentials file")

	// Mark required flags
	modifyBirthdayCmd.MarkFlagRequired("id")

	return modifyBirthdayCmd

}
