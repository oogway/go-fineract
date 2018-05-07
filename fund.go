package fineractor

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/meson10/highbrow"
)

var (
	authenticationKey AuthenticationKey
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

type FundValueRequest struct{}

type Summary struct {
	Amount           float64 `json:"accountBalance"`
	TotalDeposits    float64 `json:"totalDeposits"`
	TotalWithdrawals float64 `json:"totalWithdrawals"`
}

type FundValueResponse struct {
	Id         int64   `json:"id"`
	AccountNo  string  `json:"accountNo"`
	ClientId   int64   `json:"clientId"`
	ClientName string  `json:"clientName"`
	Statement  Summary `json:"summary"`
}

type FundsRequest struct{}

type FundsResponse struct {
	TotalFilteredRecords int64               `json:"totalFilteredRecords"`
	FundDetail           []FundValueResponse `json:"pageItems"`
}

type AuthenticationKey struct {
	Data string `json:"base64EncodedAuthenticationKey"`
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (client *Client) makeRequest(reqType, url string, payload interface{}, response interface{}) error {
	b, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		return err
	}

	req, err := http.NewRequest(reqType, url, bytes.NewBuffer(b))
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("fineract-platform-tenantid", "default")
	req.Header.Set("Authorization", "Basic "+basicAuth(client.UserName, client.Password))

	var resp *http.Response
	errTry := highbrow.Try(5, func() error {
		resp, err = client.HttpClient.Do(req)
		return err
	})
	if errTry != nil {
		rawMessage := json.RawMessage([]byte(errTry.Error()))
		return &FineractError{ErrCodeSerialization, &rawMessage}
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		rawMessage := json.RawMessage(body)
		return &FineractError{GetFineractStatusCode(resp.StatusCode), &rawMessage}
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		rawMessage := json.RawMessage([]byte(err.Error()))
		return &FineractError{ErrCodeSerialization, &rawMessage}
	}

	return err
}

func (client *Client) FundIncrement(fundId string, request FundIncrementRequest) (*FundIncrementResponse, error) {
	tempPath, _ := url.Parse(path.Join(fundId, "transactions?command=deposit"))
	path := client.HostName.ResolveReference(tempPath).String()
	var response *FundIncrementResponse
	err := client.makeRequest("POST", path, request, &response)
	if err != nil {
		return nil, err
	}

	return response, err
}

func (client *Client) FundDecrement(fundId string, request FundDecrementRequest) (*FundDecrementResponse, error) {
	tempPath, _ := url.Parse(path.Join(fundId, "transactions?command=withdrawal"))
	path := client.HostName.ResolveReference(tempPath).String()
	var response *FundDecrementResponse
	err := client.makeRequest("POST", path, request, &response)
	if err != nil {
		return nil, err
	}

	return response, err
}

func (client *Client) GetFundValue(fundId string, request FundValueRequest) (*FundValueResponse, error) {
	tempPath, _ := url.Parse(fundId)
	var response *FundValueResponse
	err := client.makeRequest("GET", client.HostName.ResolveReference(tempPath).String(), nil, &response)
	if err != nil {
		return nil, err
	}

	return response, err
}

func (client *Client) GetFunds(request FundsRequest) (*FundsResponse, error) {
	var response *FundsResponse
	err := client.makeRequest("GET", client.HostName.String(), nil, &response)
	if err != nil {
		return nil, err
	}

	return response, err
}
