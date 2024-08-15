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
	home, err := os.UserHomeDir()
	HandleErrorExit("Error finding home directory", err)

	return filepath.Join(home, ".hbd", "credentials")
}

// LoadCredentials loads the credentials from the specified path
func LoadCredentials(path string) (*Credentials, error) {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &Credentials{}, nil // File doesn't exist yet, return empty credentials
		}
		return nil, err
	}
	defer file.Close()

	var creds Credentials
	if err := json.NewDecoder(file).Decode(&creds); err != nil {
		return nil, err
	}

	return &creds, nil
}

// SaveCredentials saves the credentials to the specified path
func SaveCredentials(path string, creds *Credentials) error {
	err := os.MkdirAll(filepath.Dir(path), 0700)
	HandleErrorExit("Error creating credentials directory", err)

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	HandleErrorExit("Error opening credentials file", err)
	defer file.Close()

	return json.NewEncoder(file).Encode(creds)
}

// DeleteCredentials deletes the credentials file at the specified path.
func DeleteCredentials(credsPath string) error {
	// Attempt to remove the file
	err := os.Remove(credsPath)
	if err != nil {
		return fmt.Errorf("failed to delete credentials file at %s: %v", credsPath, err)
	}
	return nil
}
