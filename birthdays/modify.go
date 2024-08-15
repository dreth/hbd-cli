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
		Long:  `The modify-birthday command allows you to modify an existing birthday by providing its ID, new name, and/or new date.`,
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
