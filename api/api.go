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

// Helper function to handle JSON marshalling and HTTP requests
func makeRequest(method, url string, payload interface{}, token string) (*http.Response, error) {
	var reqBody []byte
	var err error

	if payload != nil {
		reqBody, err = json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("error encoding request body: %v", err)
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	client := &http.Client{}
	return client.Do(req)
}

// Add a new birthday
func AddBirthday(url, token string, birthday structs.BirthdayNameDateAdd) (*structs.BirthdayFull, error) {
	endpoint := fmt.Sprintf("%s/api/add-birthday", strings.TrimRight(url, "/"))
	resp, err := makeRequest("POST", endpoint, birthday, token)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiErr APIError
		json.NewDecoder(resp.Body).Decode(&apiErr)
		return nil, fmt.Errorf("error: %s", apiErr.Error)
	}

	var birthdayFull structs.BirthdayFull
	if err := json.NewDecoder(resp.Body).Decode(&birthdayFull); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &birthdayFull, nil
}

// Check user reminders
func CheckBirthdays(url, token string, user structs.LoginRequest) (*structs.Success, error) {
	endpoint := fmt.Sprintf("%s/api/check-birthdays", strings.TrimRight(url, "/"))
	resp, err := makeRequest("POST", endpoint, user, token)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiErr APIError
		json.NewDecoder(resp.Body).Decode(&apiErr)
		return nil, fmt.Errorf("error: %s", apiErr.Error)
	}

	var success structs.Success
	if err := json.NewDecoder(resp.Body).Decode(&success); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &success, nil
}

// Delete a birthday
func DeleteBirthday(url, token string, birthday structs.BirthdayNameDateModify) (*structs.Success, error) {
	endpoint := fmt.Sprintf("%s/api/delete-birthday", strings.TrimRight(url, "/"))
	resp, err := makeRequest("DELETE", endpoint, birthday, token)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiErr APIError
		json.NewDecoder(resp.Body).Decode(&apiErr)
		return nil, fmt.Errorf("error: %s", apiErr.Error)
	}

	var success structs.Success
	if err := json.NewDecoder(resp.Body).Decode(&success); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &success, nil
}

// Delete a user
func DeleteUser(url, token string) (*structs.Success, error) {
	endpoint := fmt.Sprintf("%s/api/delete-user", strings.TrimRight(url, "/"))
	resp, err := makeRequest("DELETE", endpoint, nil, token)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiErr APIError
		json.NewDecoder(resp.Body).Decode(&apiErr)
		return nil, fmt.Errorf("error: %s", apiErr.Error)
	}

	var success structs.Success
	if err := json.NewDecoder(resp.Body).Decode(&success); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &success, nil
}

// Generate a new password
func GeneratePassword(url string) (*structs.Password, error) {
	endpoint := fmt.Sprintf("%s/api/generate-password", strings.TrimRight(url, "/"))
	resp, err := makeRequest("GET", endpoint, nil, "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiErr APIError
		json.NewDecoder(resp.Body).Decode(&apiErr)
		return nil, fmt.Errorf("error: %s", apiErr.Error)
	}

	var password structs.Password
	if err := json.NewDecoder(resp.Body).Decode(&password); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &password, nil
}

// Check service readiness
func CheckHealth(url string) (*structs.Ready, error) {
	endpoint := fmt.Sprintf("%s/api/health", strings.TrimRight(url, "/"))
	resp, err := makeRequest("GET", endpoint, nil, "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiErr APIError
		json.NewDecoder(resp.Body).Decode(&apiErr)
		return nil, fmt.Errorf("error: %s", apiErr.Error)
	}

	var ready structs.Ready
	if err := json.NewDecoder(resp.Body).Decode(&ready); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &ready, nil
}

// Login a user
func Login(url string, user structs.LoginRequest) (*structs.LoginSuccess, error) {
	endpoint := fmt.Sprintf("%s/api/login", strings.TrimRight(url, "/"))
	resp, err := makeRequest("POST", endpoint, user, "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiErr APIError
		json.NewDecoder(resp.Body).Decode(&apiErr)
		return nil, fmt.Errorf("error: %s", apiErr.Error)
	}

	var loginSuccess structs.LoginSuccess
	if err := json.NewDecoder(resp.Body).Decode(&loginSuccess); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &loginSuccess, nil
}

// Get user data
func GetUserData(url, token string) (*structs.UserData, error) {
	endpoint := fmt.Sprintf("%s/api/me", strings.TrimRight(url, "/"))
	resp, err := makeRequest("GET", endpoint, nil, token)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiErr APIError
		json.NewDecoder(resp.Body).Decode(&apiErr)
		return nil, fmt.Errorf("error: %s", apiErr.Error)
	}

	var userData structs.UserData
	if err := json.NewDecoder(resp.Body).Decode(&userData); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &userData, nil
}

// Modify a birthday
func ModifyBirthday(url, token string, birthday structs.BirthdayNameDateModify) (*structs.Success, error) {
	endpoint := fmt.Sprintf("%s/api/modify-birthday", strings.TrimRight(url, "/"))
	resp, err := makeRequest("PUT", endpoint, birthday, token)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiErr APIError
		json.NewDecoder(resp.Body).Decode(&apiErr)
		return nil, fmt.Errorf("error: %s", apiErr.Error)
	}

	var success structs.Success
	if err := json.NewDecoder(resp.Body).Decode(&success); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &success, nil
}

// Modify a user's details
func ModifyUser(url, token string, user structs.ModifyUserRequest) (*structs.Success, error) {
	endpoint := fmt.Sprintf("%s/api/modify-user", strings.TrimRight(url, "/"))
	resp, err := makeRequest("PUT", endpoint, user, token)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiErr APIError
		json.NewDecoder(resp.Body).Decode(&apiErr)
		return nil, fmt.Errorf("error: %s", apiErr.Error)
	}

	var success structs.Success
	if err := json.NewDecoder(resp.Body).Decode(&success); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &success, nil
}

// Register a new user
func Register(url string, user structs.RegisterRequest) (*structs.LoginSuccess, error) {
	endpoint := fmt.Sprintf("%s/api/register", strings.TrimRight(url, "/"))
	resp, err := makeRequest("POST", endpoint, user, "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiErr APIError
		json.NewDecoder(resp.Body).Decode(&apiErr)
		return nil, fmt.Errorf("error: %s", apiErr.Error)
	}

	var loginSuccess structs.LoginSuccess
	if err := json.NewDecoder(resp.Body).Decode(&loginSuccess); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &loginSuccess, nil
}
