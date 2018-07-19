package fineract

import (
	"net/http"
	"time"
	"fmt"
)

const (
	registeredTableName = "m_merchant"
	officeId            = "1"
	defaultDateFormat   = "dd MMMM yyyy"
)

const (
	createClientURL = "/fineract-provider/api/v1/clients?tenantIdentifier=default"
)

type ClientInfo struct {
	FirstName      string    `json:"firstname"`
	LastName       string    `json:"lastname"`
	Active         bool      `json:"active"`
	Locale         string    `json:"locale"`
	MobileNo       string    `json:"mobileNo"`
	SubmitDate     time.Time `json:"_"`
	ActivationDate time.Time `json:"_"`
}

type createClientRequest struct {
	*ClientInfo
	OfficeID   string          `json:"officeId"`
	DateFormat string          `json:"dateFormat"`
	SubmitOn   string          `json:"submittedOnDate"`
	ActivateOn string          `json:"activationDate"`
	DataTables []dataTableInfo `json:"datatables"`
}

type dataTableInfo struct {
	TableName string            `json:"registeredTableName"`
	Data      map[string]string `json:"data"`
}

type CreateClientResponse struct {
	ID string `json:"clientId"`
}

func (client *Client) CreateClient(clientInfo *ClientInfo, merchantUserID string, merchantName string) (*CreateClientResponse, error) {
	request := &createClientRequest{
		OfficeID:   officeId,
		DateFormat: defaultDateFormat,
		SubmitOn:   formatDate(clientInfo.SubmitDate),
		ActivateOn: formatDate(clientInfo.ActivationDate),
		DataTables: []dataTableInfo{
			{
				TableName: registeredTableName,
				Data: map[string]string{
					"merchant_name":    merchantName,
					"merchant_user_id": merchantUserID,
				},
			},
		},
	}
	var response CreateClientResponse
	if err := client.MakeRequest(http.MethodPost, createClientURL, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func formatDate(input time.Time) string {
	if input.IsZero() {
		input = time.Now()
	}
	return fmt.Sprintf("%d %s %d", input.Day(), input.Month(), input.Year())
}
