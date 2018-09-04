package fineract

import (
	"net/url"
	"net/http"
)

type ChargeTemplate struct {
	Id              uint64    `json:"id"`
	Name            string    `json:"name"`
	Amount          float64   `json:"amount"`
	ChargeTime      CodeValue `json:"chargeTimeType"`
	ChargeAppliesTo CodeValue `json:"chargeAppliesTo"`
	Currency        Currency  `json:"currency"`
	Active          bool      `json:"active"`
	Penalty         bool      `json:"penalty"`
}

func (client *Client) GetAllCharges() ([]*ChargeTemplate, error) {
	tempPath, _ := url.Parse(chargesURL())
	path := client.HostName.ResolveReference(tempPath).String()
	var response []*ChargeTemplate
	if err := client.MakeRequest(http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}
