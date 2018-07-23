package fineract

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	officeId            = "1"
	registeredTableName = "m_merchant"
	defaultDateFormat   = "dd MMMM yyyy"
)

type ClientInfo struct {
	FirstName      string    `json:"firstname"`
	LastName       string    `json:"lastname"`
	Active         bool      `json:"active"`
	Locale         string    `json:"locale"`
	MobileNo       string    `json:"mobileNo"`
	CountryCode    string    `json:"countryCode,omitempty"`
	PhoneNumber    string    `json:"phoneNumber,omitempty"`
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
	ID int64 `json:"clientId"`
}

func (client *Client) CreateClient(clientInfo *ClientInfo, merchantUserID string, merchantName string) (*CreateClientResponse, error) {
	// Store phone number in "<country-code>_<phone_number>"
	clientInfo.MobileNo = fmt.Sprintf("%s_%s", clientInfo.CountryCode, clientInfo.PhoneNumber)
	// Set these to empty strings so on marshalling these will be ignored, as this endpoint doesn't accept these parameters
	clientInfo.PhoneNumber = ""
	clientInfo.CountryCode = ""

	request := &createClientRequest{
		ClientInfo: clientInfo,
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

	tempPath, _ := url.Parse(clientsURL())
	path := client.HostName.ResolveReference(tempPath).String()
	if err := client.MakeRequest(http.MethodPost, path, request, &response); err != nil {
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
