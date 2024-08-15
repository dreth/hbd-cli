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
func makeRequest(method, url string, payload interface{}, token string, result interface{}, tokenDuration int) error {
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

	// Set the token duration if provided
	if tokenDuration > 0 {
		req.Header.Set("X-Jwt-Token-Duration", fmt.Sprintf("%d", tokenDuration))
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
	err := makeRequest("POST", endpoint, birthday, token, &birthdayFull, 0)
	return &birthdayFull, err
}

// Check user reminders
func CheckBirthdays(url, token string, user structs.LoginRequest) (*structs.Success, error) {
	var success structs.Success
	endpoint := fmt.Sprintf("%s/api/check-birthdays", strings.TrimRight(url, "/"))
	err := makeRequest("POST", endpoint, user, token, &success, 0)
	return &success, err
}

// Delete a birthday
func DeleteBirthday(url, token string, birthday structs.BirthdayNameDateModify) (*structs.Success, error) {
	var success structs.Success
	endpoint := fmt.Sprintf("%s/api/delete-birthday", strings.TrimRight(url, "/"))
	err := makeRequest("DELETE", endpoint, birthday, token, &success, 0)
	return &success, err
}

// Delete a user
func DeleteUser(url, token string) (*structs.Success, error) {
	var success structs.Success
	endpoint := fmt.Sprintf("%s/api/delete-user", strings.TrimRight(url, "/"))
	err := makeRequest("DELETE", endpoint, nil, token, &success, 0)
	return &success, err
}

// Generate a new password
func GeneratePassword(url string) (*structs.Password, error) {
	var password structs.Password
	endpoint := fmt.Sprintf("%s/api/generate-password", strings.TrimRight(url, "/"))
	err := makeRequest("GET", endpoint, nil, "", &password, 0)
	return &password, err
}

// Check service readiness
func CheckHealth(url string) (*structs.Ready, error) {
	var ready structs.Ready
	endpoint := fmt.Sprintf("%s/api/health", strings.TrimRight(url, "/"))
	err := makeRequest("GET", endpoint, nil, "", &ready, 0)
	return &ready, err
}

// Login a user
func Login(url string, user structs.LoginRequest, tokenDuration int) (*structs.LoginSuccess, error) {
	var loginSuccess structs.LoginSuccess
	endpoint := fmt.Sprintf("%s/api/login", strings.TrimRight(url, "/"))
	err := makeRequest("POST", endpoint, user, "", &loginSuccess, tokenDuration)
	return &loginSuccess, err
}

// Get user data
func GetUserData(url, token string) (*structs.UserData, error) {
	var userData structs.UserData
	endpoint := fmt.Sprintf("%s/api/me", strings.TrimRight(url, "/"))
	err := makeRequest("GET", endpoint, nil, token, &userData, 0)
	return &userData, err
}

// Modify a birthday
func ModifyBirthday(url, token string, birthday structs.BirthdayNameDateModify) (*structs.Success, error) {
	var success structs.Success
	endpoint := fmt.Sprintf("%s/api/modify-birthday", strings.TrimRight(url, "/"))
	err := makeRequest("PUT", endpoint, birthday, token, &success, 0)
	return &success, err
}

// Modify a user's details (including email or password)
func ModifyUserWithEmail(url, token string, user structs.ModifyUserRequest, tokenDuration int) (*structs.LoginSuccess, error) {
	var loginSuccess structs.LoginSuccess
	endpoint := fmt.Sprintf("%s/api/modify-user", strings.TrimRight(url, "/"))
	err := makeRequest("PUT", endpoint, user, token, &loginSuccess, tokenDuration)
	return &loginSuccess, err
}

// Modify a user's details (excluding email or password)
func ModifyUserWithoutEmail(url, token string, user structs.ModifyUserRequest) (*structs.UserData, error) {
	var userData structs.UserData
	endpoint := fmt.Sprintf("%s/api/modify-user", strings.TrimRight(url, "/"))
	err := makeRequest("PUT", endpoint, user, token, &userData, 0)
	return &userData, err
}

// Register a new user
func Register(url string, user structs.RegisterRequest, tokenDuration int) (*structs.LoginSuccess, error) {
	var loginSuccess structs.LoginSuccess
	endpoint := fmt.Sprintf("%s/api/register", strings.TrimRight(url, "/"))
	err := makeRequest("POST", endpoint, user, "", &loginSuccess, tokenDuration)
	return &loginSuccess, err
}
