package fineract

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

func (client *Client) MakeRequest(reqType, url string, payload interface{}, response interface{}) error {
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
		resp, err = client.Option.Transport.Do(req)
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

	if err = json.Unmarshal(body, &response); err != nil {
		rawMessage := json.RawMessage([]byte(err.Error()))
		return &FineractError{ErrCodeSerialization, &rawMessage}
	}

	return nil
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
	tempPath, _ := url.Parse("fineract-provider/api/v1/savingsaccounts")
	var response *FundsResponse
	if err := client.MakeRequest("GET", client.HostName.ResolveReference(tempPath).String(), nil, &response); err != nil {
		return nil, err
	}

	return response, nil
}
