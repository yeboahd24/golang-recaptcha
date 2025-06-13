package recaptcha

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const verifyURL = "https://www.google.com/recaptcha/api/siteverify"

type RecaptchaResponse struct {
	Success     bool     `json:"success"`
	Score       float64  `json:"score"`
	Action      string   `json:"action"`
	ChallengeTS string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
	ErrorCodes  []string `json:"error-codes"`
}

func Verify(token, secret string) (*RecaptchaResponse, error) {
	// Prepare POST request
	resp, err := http.PostForm(verifyURL, url.Values{
		"secret":   {secret},
		"response": {token},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to send verification request: %w", err)
	}
	defer resp.Body.Close()

	// Decode response
	var result RecaptchaResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if !result.Success {
		return &result, fmt.Errorf("verification failed: %v", result.ErrorCodes)
	}

	return &result, nil
}
