package helper

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadEnvVars() {
	// Load dotenv
	godotenv.Load()

	// Read env vars
	viper.AutomaticEnv()
}
