package birthdays

import (
	"fmt"
	"hbd-cli/api"
	"hbd-cli/helper"
	"hbd-cli/structs"
	"path/filepath"

	"github.com/spf13/cobra"
)

// DeleteBirthday command
func DeleteBirthday() *cobra.Command {
	var id int64
	var host, port, credsPath string
	var ssl bool

	var deleteBirthdayCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a birthday",
		Long:  `The delete-birthday command allows you to delete an existing birthday by providing its ID.`,
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
				ID: id,
			}

			// Make the request
			_, err = api.DeleteBirthday(url, creds.Token, birthdayReq)
			helper.HandleErrorExit("Error deleting birthday", err)

			// Print success message
			fmt.Printf("Birthday with ID %d deleted successfully!\n", id)
		},
	}

	// Add flags
	deleteBirthdayCmd.Flags().Int64Var(&id, "id", 0, "ID of the birthday to delete (required)")
	deleteBirthdayCmd.Flags().StringVar(&host, "host", helper.DefaultHost(), "Host for the service")
	deleteBirthdayCmd.Flags().StringVar(&port, "port", helper.DefaultPort(), "Port for the service")
	deleteBirthdayCmd.Flags().BoolVar(&ssl, "ssl", helper.DefaultSSL(), "Use SSL (https) for the connection")
	deleteBirthdayCmd.Flags().StringVar(&credsPath, "creds-path", helper.GetDefaultCredsPath(), "Path to the credentials file")

	// Mark required flags
	deleteBirthdayCmd.MarkFlagRequired("id")

	return deleteBirthdayCmd
}
