package fineract

import (
	"time"
	"net/url"
	"path"
	"net/http"
	"fmt"
)

type ClientCharge struct {
	Id                uint64    `json:"id"`
	ChargeId          uint64    `json:"chargeId"`
	Name              string    `json:"name"`
	DueDate           []uint32  `json:"dueDate"`
	Amount            float64   `json:"amount"`
	AmountPaid        float64   `json:"amountPaid"`
	AmountOutStanding float64   `json:"amountOutstanding"`
	ChargeTime        CodeValue `json:"chargeTimeType"`
	Active            bool      `json:"isActive"`
	Paid              bool      `json:"isPaid"`
}


type chargeResponse struct {
	ResourceID int64 `json:"resourceId"`
}

func (client *Client) AddChargeToClient(clientID string, chargeID string, dueDate time.Time, amount string) (int64, error) {
	tempPath, _ := url.Parse(path.Join(clientsURL(), clientID, "charges"))
	path := client.HostName.ResolveReference(tempPath).String()
	request := struct {
		Amount     string `json:"amount"`
		ChargeID   string `json:"chargeID"`
		DateFormat string `json:"dateFormat"`
		DueDate    string `json:"dueDate"`
		Locale     string `json:"locale"`
	}{
		amount, chargeID, "dd MMMM yyyy", formatDate(dueDate), "en",
	}
	var response chargeResponse
	if err := client.MakeRequest(http.MethodPost, path, request, &response); err != nil {
		return 0, err
	}
	return response.ResourceID, nil
}

type clientChargesResponse struct {
	Items []*ClientCharge `json:"pageItems"`
}

func (client *Client) GetClientCharges(clientID string, offset int, limit int) ([]*ClientCharge, error) {
	tempPath, _ := url.Parse(path.Join(clientsURL(), clientID, "charges"))
	queryValues := tempPath.Query()
	queryValues.Add("offset", fmt.Sprintf("%d", offset))
	queryValues.Add("limit", fmt.Sprintf("%d", limit))
	tempPath.RawQuery = queryValues.Encode()

	path := client.HostName.ResolveReference(tempPath).String()
	var response clientChargesResponse
	if err := client.MakeRequest(http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}
	return response.Items, nil
}
