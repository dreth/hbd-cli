package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/fatih/color"
)

const repoURL = "https://api.github.com/repos/dreth/hbd-cli/releases/latest"

type GitHubRelease struct {
	TagName string `json:"tag_name"`
}

// GetLatestVersion fetches the latest release version from GitHub
func GetLatestVersion() (string, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(repoURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var release GitHubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", err
	}

	return release.TagName, nil
}

// CheckForNewVersion compares the current version with the latest GitHub version
func CheckForNewVersion() string {
	// Get the latest version
	latestVersion, err := GetLatestVersion()
	c := color.New(color.FgRed, color.Italic)

	if err != nil {
		return c.Sprintf(" Failed to check for the latest version: %s\n", err)
	}

	if latestVersion != "" {

		c2 := color.New(color.FgGreen, color.Bold)

		if Version != latestVersion {
			// Print the new version message in red and bold
			return c.Sprintf("\n A new version of hbd-cli is available: %s (current: %s)\n", latestVersion, Version) + c2.Sprint(" Download it: https://github.com/dreth/hbd-cli/releases/latest")
		}

		c3 := color.New(color.FgBlue, color.Bold)
		return c3.Sprintf(" You are using the latest version of hbd-cli: %s\n", Version)
	}

	return ""
}

// Print the current version of the CLI
func HBDCLIVersion() string {
	c := color.New(color.FgBlue, color.Bold)
	return c.Sprintf(` hbd-cli version: %s`+"\n", Version) + CheckForNewVersion()
}
