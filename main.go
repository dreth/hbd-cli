package main

import (
	"hbd-cli/auth"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "hbd",
		Short: "hbd-cli is a CLI tool to manage birthday reminders and notifications using the HBD backend",
	}

	// Add the login command to the root command
	rootCmd.AddCommand(auth.Login())
	rootCmd.AddCommand(auth.Register())
	rootCmd.AddCommand(auth.Me())
	rootCmd.AddCommand(auth.ModifyUser())

	// Execute the root command
	rootCmd.Execute()
}
