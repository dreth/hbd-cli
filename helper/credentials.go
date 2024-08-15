package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Credentials struct {
	Token string `json:"token"`
}

// GetDefaultCredsPath returns the default path for the credentials file
func GetDefaultCredsPath() string {
	// If there env variable HBD_CREDS_PATH is set, use this as default
	if path := os.Getenv("HBD_CREDS_PATH"); path != "" {
		return path
	}

	// Get the home directory
	home, err := os.UserHomeDir()
	HandleErrorExit("Error finding home directory", err)

	return filepath.Join(home, ".hbd", "credentials")
}

// LoadCredentials loads the credentials from the specified path
func LoadCredentials(path string) (*Credentials, error) {
	// Open the file
	file, err := os.Open(path)
	if err != nil {
		// If the file does not exist, return an error
		if os.IsNotExist(err) {
			return &Credentials{}, fmt.Errorf("Credentials file does not exist at %s", path)
		}
		return nil, err
	}
	defer file.Close()

	// Decode the credentials
	var creds Credentials
	if err := json.NewDecoder(file).Decode(&creds); err != nil {
		return nil, err
	}

	return &creds, nil
}

// SaveCredentials saves the credentials to the specified path
func SaveCredentials(path string, creds *Credentials) error {
	// Create the directory if it does not exist
	err := os.MkdirAll(filepath.Dir(path), 0700)
	HandleErrorExit("Error creating credentials directory", err)

	// Open the file
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	HandleErrorExit("Error opening credentials file", err)
	defer file.Close()

	// Encode the credentials
	return json.NewEncoder(file).Encode(creds)
}

// DeleteCredentials deletes the credentials file at the specified path.
func DeleteCredentials(credsPath string) error {
	// Attempt to remove the file
	err := os.Remove(credsPath)
	if err != nil {
		return fmt.Errorf("Failed to delete credentials file at %s: %v", credsPath, err)
	}
	return nil
}
