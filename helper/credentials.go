package helper

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Credentials struct {
	Token string `json:"token"`
}

// GetDefaultCredsPath returns the default path for the credentials file
func GetDefaultCredsPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		HandleErrorExit("Error finding home directory", err)
	}
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
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(creds)
}
