package auth

import (
	"fmt"
	"hbd-cli/api"
	"hbd-cli/helper"

	"github.com/spf13/cobra"
)

func GeneratePassword() *cobra.Command {
	var host, port string
	var ssl bool

	var generatePasswordCmd = &cobra.Command{
		Use:   "generate-password",
		Short: "Generate a new password",
		Long: `The generate-password command requests the HBD server to generate a new password.

Example usage:
  hbd-cli generate-password --host="hbd.lotiguere.com" --ssl
		`,
		Run: func(cmd *cobra.Command, args []string) {
			// Create the URL
			url := helper.GenUrl(host, port, ssl)

			// Make the request
			password, err := api.GeneratePassword(url)
			helper.HandleErrorExit("Error generating password", err)

			// Print the generated password
			fmt.Println(password.Password)
		},
	}

	// Add flags to the generate password command
	generatePasswordCmd.Flags().StringVar(&host, "host", helper.DefaultHost(), "Host for the service")
	generatePasswordCmd.Flags().StringVar(&port, "port", helper.DefaultPort(), "Port for the service")
	generatePasswordCmd.Flags().BoolVar(&ssl, "ssl", helper.DefaultSSL(), "Use SSL (https) for the connection")

	// Return the generate password command
	return generatePasswordCmd
}
