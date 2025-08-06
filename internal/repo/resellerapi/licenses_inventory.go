package resellerapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reseller-chatgpt-backend/internal/env"
)

type LicensesInventoryResponse struct {
	Items []LicensesInventoryItem `json:"items"`
	// Description string                  `json:"description"`
}

type LicensesInventoryItem struct {
	ContractType   string `json:"contractType"`
	PlanType       string `json:"planType"`
	AdditionalInfo string `json:"additionalInfo"`
	Remaining      int    `json:"remaining"`
}

// const LicensesInventoryDescription = "remaining 是天數剩餘  並且 30天為 1個月 360 天為 1 年  回答時記得轉換成年月日"

func (r *Repo) LicensesInventory(authorization string, companyID uint) (*LicensesInventoryResponse, error) {
	url := env.GetResellerURL() + fmt.Sprintf("/v1/licenses/inventory?companyId=%d", companyID)
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

	var result LicensesInventoryResponse
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse login response: %w", err)
	}

	// result.Description = LicensesInventoryDescription

	return &result, nil
}
