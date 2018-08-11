package fineract

import (
	"net/http"
	"net/url"
	"path"
)

type AddressType string

type Address struct {
	TypeDesc     string `json:"addressType"`
	Type         string `json:"addressTypeId"`
	IsActive     string `json:"isActive,omitempty"`
	Street       string `json:"street,omitempty"`
	AddressLine1 string `json:"addressLine1,omitempty"`
	AddressLine2 string `json:"addressLine2,omitempty"`
	AddressLine3 string `json:"addressLine3,omitempty"`
	TownVillage  string `json:"townVillage,omitempty"`
	City         string `json:"city,omitempty"`
	District     string `json:"countyDistrict,omitempty"`
	State        string `json:"stateProvinceId,omitempty"`
	Country      string `json:"countryId,omitempty"`
	PostalCode   string `json:"postalCode,omitempty"`
}

type CreateAddressRequest struct {
	ClientId        string
	AddressTypeCode string
	Address         Address
}

type CreateAddressResponse struct{}

func (client *Client) CreateAddress(r *CreateAddressRequest) (*CreateAddressResponse, error) {
	tempPath, err := url.Parse(clientAddressURL(r.ClientId) + "?type=" + r.AddressTypeCode)
	if err != nil {
		return nil, err
	}
	response := CreateAddressResponse{}

	err = client.MakeRequest(http.MethodPost, client.HostName.ResolveReference(tempPath).String(), r.Address, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func clientAddressURL(clientID string) string {
	return path.Join(path.Join(clientURL(), clientID), "addresses")
}
