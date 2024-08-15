package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hbd-cli/structs"
	"net/http"
	"strings"
)

// Define a common error type for API responses
type APIError struct {
	Error string `json:"error"`
}

// Helper function to handle JSON marshalling, HTTP requests, and response decoding
func makeRequest(method, url string, payload interface{}, token string, result interface{}) error {
	var reqBody []byte
	var err error

	if payload != nil {
		reqBody, err = json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("error encoding request body: %v", err)
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiErr APIError
		json.NewDecoder(resp.Body).Decode(&apiErr)
		return fmt.Errorf("error: %s", apiErr.Error)
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return fmt.Errorf("error decoding response: %v", err)
		}
	}

	return nil
}

// Add a new birthday
func AddBirthday(url, token string, birthday structs.BirthdayNameDateAdd) (*structs.BirthdayFull, error) {
	var birthdayFull structs.BirthdayFull
	endpoint := fmt.Sprintf("%s/api/add-birthday", strings.TrimRight(url, "/"))
	err := makeRequest("POST", endpoint, birthday, token, &birthdayFull)
	return &birthdayFull, err
}

// Check user reminders
func CheckBirthdays(url, token string, user structs.LoginRequest) (*structs.Success, error) {
	var success structs.Success
	endpoint := fmt.Sprintf("%s/api/check-birthdays", strings.TrimRight(url, "/"))
	err := makeRequest("POST", endpoint, user, token, &success)
	return &success, err
}

// Delete a birthday
func DeleteBirthday(url, token string, birthday structs.BirthdayNameDateModify) (*structs.Success, error) {
	var success structs.Success
	endpoint := fmt.Sprintf("%s/api/delete-birthday", strings.TrimRight(url, "/"))
	err := makeRequest("DELETE", endpoint, birthday, token, &success)
	return &success, err
}

// Delete a user
func DeleteUser(url, token string) (*structs.Success, error) {
	var success structs.Success
	endpoint := fmt.Sprintf("%s/api/delete-user", strings.TrimRight(url, "/"))
	err := makeRequest("DELETE", endpoint, nil, token, &success)
	return &success, err
}

// Generate a new password
func GeneratePassword(url string) (*structs.Password, error) {
	var password structs.Password
	endpoint := fmt.Sprintf("%s/api/generate-password", strings.TrimRight(url, "/"))
	err := makeRequest("GET", endpoint, nil, "", &password)
	return &password, err
}

// Check service readiness
func CheckHealth(url string) (*structs.Ready, error) {
	var ready structs.Ready
	endpoint := fmt.Sprintf("%s/api/health", strings.TrimRight(url, "/"))
	err := makeRequest("GET", endpoint, nil, "", &ready)
	return &ready, err
}

// Login a user
func Login(url string, user structs.LoginRequest) (*structs.LoginSuccess, error) {
	var loginSuccess structs.LoginSuccess
	endpoint := fmt.Sprintf("%s/api/login", strings.TrimRight(url, "/"))
	err := makeRequest("POST", endpoint, user, "", &loginSuccess)
	return &loginSuccess, err
}

// Get user data
func GetUserData(url, token string) (*structs.UserData, error) {
	var userData structs.UserData
	endpoint := fmt.Sprintf("%s/api/me", strings.TrimRight(url, "/"))
	err := makeRequest("GET", endpoint, nil, token, &userData)
	return &userData, err
}

// Modify a birthday
func ModifyBirthday(url, token string, birthday structs.BirthdayNameDateModify) (*structs.Success, error) {
	var success structs.Success
	endpoint := fmt.Sprintf("%s/api/modify-birthday", strings.TrimRight(url, "/"))
	err := makeRequest("PUT", endpoint, birthday, token, &success)
	return &success, err
}

// Modify a user's details
func ModifyUser(url, token string, user structs.ModifyUserRequest) (*structs.Success, error) {
	var success structs.Success
	endpoint := fmt.Sprintf("%s/api/modify-user", strings.TrimRight(url, "/"))
	err := makeRequest("PUT", endpoint, user, token, &success)
	return &success, err
}

// Register a new user
func Register(url string, user structs.RegisterRequest) (*structs.LoginSuccess, error) {
	var loginSuccess structs.LoginSuccess
	endpoint := fmt.Sprintf("%s/api/register", strings.TrimRight(url, "/"))
	err := makeRequest("POST", endpoint, user, "", &loginSuccess)
	return &loginSuccess, err
}
