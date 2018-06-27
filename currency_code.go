package fineract

import (
	"net/url"
	"path"
)

const (
	getCurrencyCodeURL = "currencies?fields=selectedCurrencyOptions"
)

// CurrencyCode presents a currency code obj returned by endpoint get currency code
type CurrencyCode struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	DecimalPlaces int64  `json:"decimalPlaces"`
	NameCode      string `json:"nameCode"`
	DisplayLabel  string `json:"displayLabel"`
	DisplaySymbol string `json:"displaySymbol"`
}

// CurrencyCodeResponse presents the response of endpoint get currency code
type CurrencyCodeResponse struct {
	SelectedCurrencyOptions []CurrencyCode `json:"selectedCurrencyOptions"`
}

func getCurrencyCodePath() string {
	return path.Join(baseURL, getCurrencyCodeURL)
}

// GetCurrencyCode gets selected currency code from fineract
func (client *Client) GetCurrencyCode() (*CurrencyCodeResponse, error) {
	tempPath, err := url.Parse(getCurrencyCodePath())
	if err != nil {
		return nil, err
	}
	response := CurrencyCodeResponse{}
	err = client.MakeRequest("GET", client.HostName.ResolveReference(tempPath).String(), nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
