package fineract

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"path"
	"strconv"
)

const (
	fineractHost     = "https://13.209.34.65:8443" //"https://demo.openmf.org"
	fineractUser     = "mifos"
	fineractPassword = "password"
	baseURL          = "fineract-provider/api/v1"
	Locale           = "en"
	DateFormat       = "dd MMMM yyyy"

	active   = "Active"
	pgOffset = "0"
	pgLimit  = "100"

	Principal AccountType = "Fund Principal"
	Interest  AccountType = "Fund Interest"
	Deposit   AccountTx   = "deposit"
	Withdraw  AccountTx   = "withdrawal"
)

type AccountType string

func toString(a interface{}) string {
	return fmt.Sprintf("%v", a)
}

type AccountTx string

type FundAccountId struct {
	PrincipalAccountId string
	InterestAccountId  string
}

type Office struct {
	Id   uint32 `json:"id"`
	Name string `json:"externalId"`
}

type PaymentType struct {
	Id            uint32 `json:"id"`
	Name          string `json:"name"`
	IsCashPayment bool   `json:"isCashPayment"`
}

type StatusT struct {
	Code  string `json:"code"`
	Value string `json:"value"`
}

type TxRequest struct {
	Locale            string `json:"locale"`
	DateFormat        string `json:"dateFormat"`
	TransactionDate   string `json:"transactionDate"`
	TransactionAmount string `json:"transactionAmount"`
	PaymentTypeId     string `json:"paymentTypeId"`
}

type TxResponse struct {
	OfficeId   float64 `json:"officeId"`
	ClientId   float64 `json:"clientId"`
	ResourceId float64 `json:"resourceId"`
}

type Summary struct {
	Balance          float64 `json:"accountBalance"`
	Limit            float64 `json:"availableBalance"`
	TotalDeposits    float64 `json:"totalDeposits"`
	TotalWithdrawals float64 `json:"totalWithdrawals"`
}

type AccountDetails struct {
	Id          uint64  `json:"id"`
	AccountNo   string  `json:"accountNo"`
	ProductId   uint64  `json:"savingsProductId"`
	ProductName string  `json:"savingsProductName"`
	Status      StatusT `json:"status"`
	Statement   Summary `json:"summary"`
}

type FundAccount struct {
	Id             uint64  `json:"id"`
	AccountNo      string  `json:"accountNo"`
	ProductId      uint64  `json:"productId"`
	ProductName    string  `json:"productName"`
	Status         StatusT `json:"status"`
	AccountBalance float64 `json:"accountBalance"`
}

type FundAccountResponse struct {
	FundAccount []FundAccount `json:"savingsAccounts"`
}

type GetPaymentTypeRequest struct {
}

type GetPaymentTypeResponse struct {
	PaymentMethod []PaymentType
}

type GetFundTypeRequest struct {
	Name string
}

type Fund struct {
	Id      uint64  `json:"id"`
	Name    string  `json:"fullname"`
	Status  StatusT `json:"status"`
	Balance float64
}

type FundsRequest struct {
	Type string
}

type FundsResponse struct {
	TotalFilteredRecords uint64 `json:"totalFilteredRecords"`
	Fund                 []Fund `json:"pageItems"`
}

func (client *Client) AccountDeposit(accountId string, request *TxRequest) (*TxResponse, error) {
	return client.TransactSavingsAccount(accountId, Deposit, request)
}

func (client *Client) AccountWithdraw(accountId string, request *TxRequest) (*TxResponse, error) {
	return client.TransactSavingsAccount(accountId, Withdraw, request)
}

func (client *Client) TransactSavingsAccount(accountId string, txType AccountTx, request *TxRequest) (*TxResponse, error) {
	if amt, _ := strconv.ParseUint(request.TransactionAmount, 10, 64); amt == 0 {
		return &TxResponse{}, nil
	}

	command := "transactions?command=" + toString(txType)
	tempPath, _ := url.Parse(path.Join(savingsAccountsURL(), path.Join(accountId, command)))
	path := client.HostName.ResolveReference(tempPath).String()
	var response *TxResponse
	if err := client.MakeRequest("POST", path, request, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (client *Client) GetAccount(accountId string) (*AccountDetails, error) {
	tempPath, _ := url.Parse(path.Join(savingsAccountsURL(), accountId))
	path := client.HostName.ResolveReference(tempPath).String()
	var response *AccountDetails
	if err := client.MakeRequest("GET", path, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (client *Client) GetFund(fundId string) (*Fund, error) {
	tempPath, _ := url.Parse(path.Join(clientsURL(), fundId))
	path := client.HostName.ResolveReference(tempPath).String()
	var response *Fund
	if err := client.MakeRequest("GET", path, nil, &response); err != nil {
		return nil, err
	}
	log.Printf("%v", response)
	return response, nil
}

func (client *Client) GetPaymentType(request *GetPaymentTypeRequest) (*GetPaymentTypeResponse, error) {
	tempPath, _ := url.Parse(paymentTypesURL())
	var response []PaymentType
	if err := client.MakeRequest("GET", client.HostName.ResolveReference(tempPath).String(), nil, &response); err != nil {
		return nil, err
	}
	return &GetPaymentTypeResponse{response}, nil
}

func (client *Client) GetFundType(request *GetFundTypeRequest) (uint32, error) {
	tempPath, _ := url.Parse(headOfficeURL())
	var office []Office
	if err := client.MakeRequest("GET", client.HostName.ResolveReference(tempPath).String(), nil, &office); err != nil {
		return 0, err
	}

	for _, cursor := range office {
		if cursor.Name == request.Name {
			return cursor.Id, nil
		}
	}
	return 0, errors.New(fmt.Sprintf("No FundType with name %v found", request.Name))
}

func (client *Client) GetFundValue(fundId string) (float64, error) {
	response, err := client.GetFundAccounts(fundId)
	if err != nil {
		return 0, err
	}

	for _, cursor := range response.FundAccount {
		if cursor.ProductName == toString(Principal) && cursor.Status.Value == active {
			return cursor.AccountBalance, nil
		}
	}
	return 0, errors.New("No active account of type " + toString(Principal))
}

func (client *Client) GetFundAccounts(fundId string) (*FundAccountResponse, error) {
	tempPath, _ := url.Parse(path.Join(clientsURL(), path.Join(fundId, "accounts?fields=savingsAccounts")))
	var response *FundAccountResponse
	if err := client.MakeRequest("GET", client.HostName.ResolveReference(tempPath).String(), nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (client *Client) GetFunds(request *FundsRequest) (*FundsResponse, error) {
	tempPath, _ := url.Parse(clientsURL() + fundsURLParams(request.Type))
	var fundsResponse *FundsResponse
	if err := client.MakeRequest("GET", client.HostName.ResolveReference(tempPath).String(), nil, &fundsResponse); err != nil {
		return nil, err
	}
	for _, cursor := range fundsResponse.Fund {
		if cursor.Status.Value == active {
			balance, err := client.GetFundValue(toString(cursor.Id))
			if err == nil {
				cursor.Balance = balance
			}
		}
	}
	log.Println(fundsResponse)
	return fundsResponse, nil
}

func (client *Client) GetFundAccountId(fundId string) (*FundAccountId, error) {
	response, err := client.GetFundAccounts(fundId)
	if err != nil {
		return nil, err
	}

	fundAccountId := &FundAccountId{}

	for _, cursor := range response.FundAccount {
		if cursor.Status.Value == active && cursor.ProductName == toString(Principal) {
			fundAccountId.PrincipalAccountId = toString(cursor.Id)
		}
		if cursor.Status.Value == active && cursor.ProductName == toString(Interest) {
			fundAccountId.InterestAccountId = toString(cursor.Id)
		}
		if fundAccountId.PrincipalAccountId != "" && fundAccountId.InterestAccountId != "" {
			break
		}
	}
	log.Println(fundAccountId)
	return fundAccountId, nil
}

func fundsURLParams(officeId string) string {
	fieldFilter := "fields=id,status,fullname"
	offset := "offset=" + pgOffset
	limit := "limit=" + pgLimit
	office := "officeId=" + officeId
	return "?" + office + "&" + fieldFilter + "&" + offset + "&" + limit
}

func clientsURL() string {
	return path.Join(baseURL, "clients")
}

func paymentTypesURL() string {
	return path.Join(baseURL, "paymenttypes")
}

func savingsAccountsURL() string {
	return path.Join(baseURL, "savingsaccounts")
}

func headOfficeURL() string {
	return path.Join(baseURL, "offices")
}

func NewMockClient() (*Client, error) {
	return NewClient(fineractHost, fineractUser, fineractPassword, FineractOption{
		Transport: &MockTransport{DirectoryPath: "../testdata"},
	})
}
