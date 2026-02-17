package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	BattleNetTokenURL = "https://oauth.battle.net/token"
)

var Token string

// BattleNetTokenResponse représente la réponse JSON de l'endpoint token Battle.net.
type BattleNetTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope,omitempty"`
}

// GetBattleNetToken demande un token OAuth à Battle.net.
func GetBattleNetToken(clientID, clientSecret string) (accessToken string, expiresIn int, err error) {
	client := &http.Client{Timeout: 15 * time.Second}

	form := url.Values{}
	form.Set("grant_type", "client_credentials")

	req, reqErr := http.NewRequest(http.MethodPost, BattleNetTokenURL, strings.NewReader(form.Encode()))
	if reqErr != nil {
		return "", 0, fmt.Errorf("création requête token: %w", reqErr)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(clientID, clientSecret)

	resp, doErr := client.Do(req)
	if doErr != nil {
		return "", 0, fmt.Errorf("appel token: %w", doErr)
	}
	defer resp.Body.Close()

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return "", 0, fmt.Errorf("lecture réponse token: %w", readErr)
	}

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("token: code %d, body: %s", resp.StatusCode, string(body))
	}

	var data BattleNetTokenResponse
	if jsonErr := json.Unmarshal(body, &data); jsonErr != nil {
		return "", 0, fmt.Errorf("décodage réponse token: %w", jsonErr)
	}

	return data.AccessToken, data.ExpiresIn, nil
}
