package token

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type GetTokenResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
}

func GenerateToken(clientId, clientSecret, refreshToken string) string {
	url := "https://api.hubapi.com/oauth/v1/token"
	method := "POST"
	payload := strings.NewReader(fmt.Sprintf("grant_type=refresh_token&client_id=%s&client_secret=%s&refresh_token=%s", clientId, clientSecret, refreshToken))

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	res, _ := client.Do(req)

	token := &GetTokenResponse{}
	_ = json.NewDecoder(res.Body).Decode(token)

	return token.AccessToken
}
