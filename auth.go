package fineract

import (
	"net/http"
	"path"
	"fmt"
	"net/url"
)

type AuthResponse struct {
	Username    string   `json:"username"`
	UserId      int64    `json:"userId"`
	Token       string   `json:"base64EncodedAuthenticationKey"`
	OfficeId    int64    `json:"officeId"`
	OfficeName  string   `json:"officeName"`
	Permissions []string `json:"permissions"`
}

type AuthRequest struct {
	Username string
	Password string
}

func (client *Client) Auth(r *AuthRequest) (*AuthResponse, error) {
	if r.Username == "" || r.Password == "" {
		return nil, fmt.Errorf("username and password can not be empty")
	}

	tempPath, err := url.Parse(authURLWithCredential(r))
	if err != nil {
		return nil, err
	}

	response := AuthResponse{}
	err = client.MakeRequest(http.MethodPost, client.HostName.ResolveReference(tempPath).String(), nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func authURLWithCredential(r *AuthRequest) string {
	return path.Join(baseURL, fmt.Sprintf("authentication?username=%s&password=%s", r.Username, r.Password))
}
