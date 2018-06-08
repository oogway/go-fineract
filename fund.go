package fineract

import (
	"net/url"
	"path"
)

type FundIncrementRequest struct {
	Locale            string `json:"locale"`
	DateFormat        string `json:"dateFormat"`
	TransactionDate   string `json:"transactionDate"`
	TransactionAmount string `json:"transactionAmount"`
	PaymentTypeId     string `json:"paymentTypeId"`
}

type FundIncrementResponse struct {
	OfficeId   float64 `json:"officeId"`
	ClientId   float64 `json:"clientId"`
	ResourceId float64 `json:"resourceId"`
}

type FundDecrementRequest struct {
	Locale            string `json:"locale"`
	DateFormat        string `json:"dateFormat"`
	TransactionDate   string `json:"transactionDate"`
	TransactionAmount string `json:"transactionAmount"`
	PaymentTypeId     string `json:"paymentTypeId"`
}

type FundDecrementResponse struct {
	OfficeId   float64 `json:"officeId"`
	ClientId   float64 `json:"clientId"`
	ResourceId float64 `json:"resourceId"`
}

type GetPaymentTypeRequest struct{}

type GetPaymentTypeResponse struct {
	PaymentMethod []PaymentType
}

type PaymentType struct {
	Id            uint32 `json:"id"`
	Name          string `json:"name"`
	IsCashPayment bool   `json:"isCashPayment"`
}

type FundValueRequest struct{}

type Summary struct {
	Amount           float64 `json:"accountBalance"`
	TotalDeposits    float64 `json:"totalDeposits"`
	TotalWithdrawals float64 `json:"totalWithdrawals"`
}

type Currency struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	DisplaySymbol string `json:"displaySymbol"`
}

type FundValueResponse struct {
	Id         int64    `json:"id"`
	AccountNo  string   `json:"accountNo"`
	ClientId   int64    `json:"clientId"`
	ClientName string   `json:"clientName"`
	Statement  Summary  `json:"summary"`
	Currency   Currency `json:"currency"`
}

type FundsRequest struct{}

type FundsResponse struct {
	TotalFilteredRecords int64               `json:"totalFilteredRecords"`
	FundDetail           []FundValueResponse `json:"pageItems"`
}

func (client *Client) FundIncrement(fundId string, request *FundIncrementRequest) (*FundIncrementResponse, error) {
	tempPath, _ := url.Parse(path.Join("fineract-provider/api/v1/savingsaccounts", path.Join(fundId, "transactions?command=deposit")))
	path := client.HostName.ResolveReference(tempPath).String()
	var response *FundIncrementResponse
	if err := client.MakeRequest("POST", path, request, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (client *Client) FundDecrement(fundId string, request *FundDecrementRequest) (*FundDecrementResponse, error) {
	tempPath, _ := url.Parse(path.Join("fineract-provider/api/v1/savingsaccounts", path.Join(fundId, "transactions?command=withdrawal")))
	path := client.HostName.ResolveReference(tempPath).String()
	var response *FundDecrementResponse
	if err := client.MakeRequest("POST", path, request, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (client *Client) GetFundValue(fundId string, request *FundValueRequest) (*FundValueResponse, error) {
	tempPath, _ := url.Parse(path.Join("fineract-provider/api/v1/savingsaccounts", fundId))
	var response *FundValueResponse
	if err := client.MakeRequest("GET", client.HostName.ResolveReference(tempPath).String(), nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (client *Client) GetFunds(request *FundsRequest) (*FundsResponse, error) {
	tempPath, _ := url.Parse("fineract-provider/api/v1/savingsaccounts?fields=id,accountNo,clientId,clientName,summary&offset=0&limit=10")
	var response *FundsResponse
	if err := client.MakeRequest("GET", client.HostName.ResolveReference(tempPath).String(), nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (client *Client) GetPaymentType(request *GetPaymentTypeRequest) (*GetPaymentTypeResponse, error) {
	tempPath, _ := url.Parse("fineract-provider/api/v1/paymenttypes")
	var response []PaymentType
	if err := client.MakeRequest("GET", client.HostName.ResolveReference(tempPath).String(), nil, &response); err != nil {
		return nil, err
	}
	return &GetPaymentTypeResponse{response}, nil
}
