package helper

import (
	"os"
	"path/filepath"
	"strings"
)

func InterpretTildeAsHomeDir(path string) string {
	// Expand `~` to home directory if present
	if strings.HasPrefix(path, "~") {
		homeDir, err := os.UserHomeDir()
		HandleErrorExit("Error fetching the home directory", err)

		// Replace `~` with the home directory
		path = filepath.Join(homeDir, path[1:])
	}

	return path
}
