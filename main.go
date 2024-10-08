package main

import (
	"hbd-cli/auth"
	"hbd-cli/birthdays"
	"hbd-cli/general"

	"github.com/spf13/cobra"
)

var Version = "dev"

func main() {
	var rootCmd = &cobra.Command{
		Use:     "hbd",
		Short:   general.SplashScreen(true) + "\n" + CheckForNewVersion(),
		Long:    general.SplashScreen(true) + "\n" + CheckForNewVersion(),
		Version: general.SplashScreen(false) + "\n" + HBDCLIVersion() + "\n",
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
	authCmd.AddCommand(auth.Logout())

	// Add birthday subcommands under the 'birthdays' parent command
	birthdaysCmd.AddCommand(birthdays.AddBirthday())
	birthdaysCmd.AddCommand(birthdays.ListBirthdays())
	birthdaysCmd.AddCommand(birthdays.DeleteBirthday())
	birthdaysCmd.AddCommand(birthdays.ModifyBirthday())
	birthdaysCmd.AddCommand(birthdays.CheckBirthdays())

	// Add the internal verbs to the root command
	rootCmd.AddCommand(authCmd)
	rootCmd.AddCommand(birthdaysCmd)

	// Healthcheck command
	rootCmd.AddCommand(general.HealthCheck())

	// Execute the root command
	rootCmd.Execute()
}
