package models

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type UserInfo struct {
	ID          string `json:"id"`
	Login       string `json:"login"`
	Email       string `json:"default_email"`
	DisplayName string `json:"display_name"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}

type YandexService struct {
	tokenUrl           string
	loginUrl           string
	YandexId           string
	YandexRedirectURI  string
	YandexClientSecret string
}

func NewYandexService(id, uri, secret string) *YandexService {
	return &YandexService{
		tokenUrl:           "https://oauth.yandex.ru/token",
		loginUrl:           "https://login.yandex.ru/info",
		YandexId:           id,
		YandexRedirectURI:  uri,
		YandexClientSecret: secret,
	}
}

func (s *YandexService) ExchangeCodeToken(code string) (*TokenResponse, error) {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("client_id", s.YandexId)
	data.Set("client_secret", s.YandexClientSecret)
	data.Set("redirect_uri", s.YandexRedirectURI)

	resp, err := http.PostForm(s.tokenUrl, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, err
	}

	return &tokenResp, nil
}

func (s *YandexService) GetUserInfo(token *TokenResponse) (*UserInfo, error) {
	req, err := http.NewRequest("GET", s.loginUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "OAuth "+token.AccessToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var userInfo UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}
