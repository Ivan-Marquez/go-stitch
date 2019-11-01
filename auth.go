package gostitch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// AuthResponse with auth token
type AuthResponse struct {
	AccessToken string `json:"access_token"`
}

func createAuthRequest(user string, apiKey string) (*http.Request, error) {
	const url = BaseURL + "/auth/providers/mongodb-cloud/login"

	// set body
	json := fmt.Sprintf(`{"username": "%v", "apiKey": "%v"}`, user, apiKey)
	jsonStr := []byte(json)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	if err != nil {
		return nil, err
	}

	// set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return req, nil
}

// Authenticate to get an authorization token from MongoDB Cloud API
func Authenticate(user string, apiKey string) (response *AuthResponse, err error) {
	client := http.Client{}

	req, err := createAuthRequest(user, apiKey)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}
