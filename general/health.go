package general

import (
	"fmt"
	"hbd-cli/api"
	"hbd-cli/helper"

	"github.com/spf13/cobra"
)

func HealthCheck() *cobra.Command {
	var host string
	var port string
	var ssl bool

	var healthCheckCmd = &cobra.Command{
		Use:   "health",
		Short: "Health check the HBD service",
		Long: `The health command checks the readiness of the HBD service.

Environment variables:
  HBD_HOST - The host for the service. Defaults to 0.0.0.0.
  HBD_PORT - The port for the service. 
  HBD_SSL - Use SSL (https) for the connection.

Example usage:
  hbd-cli health --host="hbd.lotiguere.com" --ssl --port="8080"
		`,
		Run: func(cmd *cobra.Command, args []string) {
			// Create the URL
			url := helper.GenUrl(host, port, ssl)
			fmt.Printf("Performing a health check on HBD host: %s\n", url)

			// Make the request
			health, err := api.CheckHealth(url)
			helper.HandleErrorExit("Error checking health", err)

			// Print the health status
			fmt.Printf("Service health status: %s\n", health.Status)
		},
	}

	// Add flags to the health check command
	healthCheckCmd.Flags().StringVar(&host, "host", helper.DefaultHost(), "Host for the service")
	healthCheckCmd.Flags().StringVar(&port, "port", helper.DefaultPort(), "Port for the service")
	healthCheckCmd.Flags().BoolVar(&ssl, "ssl", helper.DefaultSSL(), "Use SSL (https) for the connection")

	// Return the health check command
	return healthCheckCmd
}
