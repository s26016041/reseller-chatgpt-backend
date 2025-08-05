package resellerapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reseller-chatgpt-backend/internal/env"
)

type loginBody struct {
	AuthFlow       string         `json:"AuthFlow"`
	AuthParameters authParameters `json:"AuthParameters"`
	ClientID       string         `json:"ClientId"`
}

type authParameters struct {
	Username string `json:"USERNAME"`
	Password string `json:"PASSWORD"`
}

const (
	authFlowPassword = "USER_PASSWORD_AUTH"
	cognitoTarget    = "AWSCognitoIdentityProviderService.InitiateAuth"
)

type loginResponse struct {
	AuthenticationResult struct {
		IDToken      string `json:"IdToken"`
		AccessToken  string `json:"AccessToken"`
		RefreshToken string `json:"RefreshToken"`
	} `json:"AuthenticationResult"`
}

func (r *Repo) Login(username string, password string) (string, error) {
	url := env.GetCognitoLoginURL() + "/"

	body := loginBody{
		AuthFlow: authFlowPassword,
		AuthParameters: authParameters{
			Username: username,
			Password: password,
		},
		ClientID: env.GetCognitoClientID(),
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("failed to marshal login body: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to create login request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-amz-json-1.1")
	req.Header.Set("X-Amz-Target", cognitoTarget)

	resp, err := r.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to perform login request: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("login failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var result loginResponse
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return "", fmt.Errorf("failed to parse login response: %w", err)
	}

	return result.AuthenticationResult.IDToken, nil
}
