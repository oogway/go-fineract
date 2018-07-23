package fineract

import (
	"net/http"
	"net/url"
	"path"
)

type UserInfoResponse struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	DisplayName string `json:"displayname"`
	MobileNo    string `json:"mobileNo"`
}

func (client *Client) GetClientInfo(clientId string) (*UserInfoResponse, error) {
	tempPath, _ := url.Parse(path.Join(clientsURL(), clientId))
	path := client.HostName.ResolveReference(tempPath).String()

	var response *UserInfoResponse
	if err := client.MakeRequest(http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}
