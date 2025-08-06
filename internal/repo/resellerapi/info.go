package resellerapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reseller-chatgpt-backend/internal/env"
)

type InfoResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (r *Repo) Info(authorization string) (*InfoResponse, error) {
	url := env.GetResellerURL() + "/v1/companies/info"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+authorization)
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform login request: %w", err)
	}

	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("login failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var result InfoResponse
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse login response: %w", err)
	}

	return &result, nil
}
