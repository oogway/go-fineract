package fineract

import (
	"net/http"
	"net/url"
	"path"
	"strings"
)

type UserInfoResponse struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	DisplayName string `json:"displayname"`
	MobileNo    string `json:"mobileNo"`
	CountryCode string `json:"countryCode,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
}

func (client *Client) GetClient(clientId string) (*UserInfoResponse, error) {
	tempPath, _ := url.Parse(path.Join(clientsURL(), clientId))
	path := client.HostName.ResolveReference(tempPath).String()

	var response *UserInfoResponse
	if err := client.MakeRequest(http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}

	//NOTE: For backward compatibility error is not being raised here, since many numbers are stored without country-code
	contact := strings.Split(response.MobileNo, "_")
	if len(contact) < 2 {
		response.PhoneNumber = response.MobileNo
		return response, nil
	} else {
		response.CountryCode = contact[0]
		response.PhoneNumber = contact[1]
		return response, nil
	}
	return response, nil
}
