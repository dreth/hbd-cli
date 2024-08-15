package main

import (
	"fmt"
	"hbd-cli/auth"
	"hbd-cli/birthdays"
	"hbd-cli/general"

	"github.com/spf13/cobra"
)

var Version = "dev"

func main() {
	var rootCmd = &cobra.Command{
		Use:     "hbd",
		Short:   general.SplashScreen(true),
		Long:    general.SplashScreen(true),
		Version: general.SplashScreen(false) + "\n" + fmt.Sprintf("hbd-cli version: %s", Version) + "\033[32m",
	}

	// Create an 'auth' parent command
	var authCmd = &cobra.Command{
		Use:   "auth",
		Short: "Authentication related commands (login, register, etc.)",
	}

	// Create a 'birthdays' parent command
	var birthdaysCmd = &cobra.Command{
		Use:   "birthdays",
		Short: "Birthday related commands (add, list, delete, modify)",
	}

	// Add authentication subcommands under the 'auth' parent command
	authCmd.AddCommand(auth.Login())
	authCmd.AddCommand(auth.Register())
	authCmd.AddCommand(auth.Me())
	authCmd.AddCommand(auth.ModifyUser())
	authCmd.AddCommand(auth.DeleteUser())
	authCmd.AddCommand(auth.GeneratePassword())

	// Add birthday subcommands under the 'birthdays' parent command
	birthdaysCmd.AddCommand(birthdays.AddBirthday())
	birthdaysCmd.AddCommand(birthdays.ListBirthdays())
	birthdaysCmd.AddCommand(birthdays.DeleteBirthday())
	birthdaysCmd.AddCommand(birthdays.ModifyBirthday())

	// Add the internal verbs to the root command
	rootCmd.AddCommand(authCmd)
	rootCmd.AddCommand(birthdaysCmd)

	// Healthcheck command
	rootCmd.AddCommand(general.HealthCheck())

	// Execute the root command
	rootCmd.Execute()
}
