package fineractor

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"

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

// TODO: The service will send username and password in headers
func getAuthenticationKey() (string, error) {
	if authenticationKey.Data != "" {
		return authenticationKey.Data, nil
	}
	tempPath, _ := url.Parse("authentication?username=mifos&password=password")
	tempPath1 := client.HostName.ResolveReference(tempPath).String()
	path := strings.Replace(tempPath1, "/savingsaccounts", "", -1)
	req, err := http.NewRequest("POST", path, bytes.NewBuffer([]byte{}))
	if err != nil {
		log.Println(err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("fineract-platform-tenantid", "default")

	var resp *http.Response
	errTry := highbrow.Try(5, func() error {
		resp, err = client.HttpClient.Do(req)
		return err
	})
	if errTry != nil {
		log.Println(errTry.Error())
		rawMessage := json.RawMessage([]byte(errTry.Error()))
		return "", &FineractError{ErrCodeSerialization, &rawMessage}
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		rawMessage := json.RawMessage(body)
		return "", &FineractError{GetFineractStatusCode(resp.StatusCode), &rawMessage}
	}

	err = json.Unmarshal(body, &authenticationKey)
	if err != nil {
		rawMessage := json.RawMessage([]byte(err.Error()))
		return "", &FineractError{ErrCodeSerialization, &rawMessage}
	}

	return authenticationKey.Data, err

}

func makeRequest(reqType, url string, payload interface{}) ([]byte, error) {
	b, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req, err := http.NewRequest(reqType, url, bytes.NewBuffer(b))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("fineract-platform-tenantid", "default")
	authKey, err := getAuthenticationKey()
	if err != nil {
		rawMessage := json.RawMessage([]byte(err.Error()))
		return nil, &FineractError{ErrAuthenticationFailure, &rawMessage}
	}
	req.Header.Set("Authorization", "Basic "+authKey)

	var resp *http.Response
	errTry := highbrow.Try(5, func() error {
		resp, err = client.HttpClient.Do(req)
		return err
	})
	if errTry != nil {
		rawMessage := json.RawMessage([]byte(errTry.Error()))
		return nil, &FineractError{ErrCodeSerialization, &rawMessage}
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		rawMessage := json.RawMessage(body)
		return nil, &FineractError{GetFineractStatusCode(resp.StatusCode), &rawMessage}
	}

	return body, err
}

func (client *Client) FundIncrement(fundId string, request FundIncrementRequest) (*FundIncrementResponse, error) {
	tempPath, _ := url.Parse(path.Join(fundId, "transactions?command=deposit"))
	path := client.HostName.ResolveReference(tempPath).String()
	body, err := makeRequest("POST", path, request)
	if err != nil {
		return nil, err
	}
	var response FundIncrementResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		rawMessage := json.RawMessage([]byte(err.Error()))
		return nil, &FineractError{ErrCodeSerialization, &rawMessage}
	}

	return &response, err
}

func (client *Client) FundDecrement(fundId string, request FundDecrementRequest) (*FundDecrementResponse, error) {
	tempPath, _ := url.Parse(path.Join(fundId, "transactions?command=withdrawal"))
	path := client.HostName.ResolveReference(tempPath).String()
	body, err := makeRequest("POST", path, request)
	if err != nil {
		return nil, err
	}
	var response FundDecrementResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		rawMessage := json.RawMessage([]byte(err.Error()))
		return nil, &FineractError{ErrCodeSerialization, &rawMessage}
	}

	return &response, err
}

func (client *Client) GetFundValue(fundId string, request FundValueRequest) (*FundValueResponse, error) {
	tempPath, _ := url.Parse(fundId)
	body, err := makeRequest("GET", client.HostName.ResolveReference(tempPath).String(), nil)
	if err != nil {
		return nil, err
	}
	var response FundValueResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		rawMessage := json.RawMessage([]byte(err.Error()))
		return nil, &FineractError{ErrCodeSerialization, &rawMessage}
	}
	return &response, err
}

func (client *Client) GetFunds(request FundsRequest) (*FundsResponse, error) {
	body, err := makeRequest("GET", client.HostName.String(), nil)
	if err != nil {
		return nil, err
	}
	var response FundsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		rawMessage := json.RawMessage([]byte(err.Error()))
		return nil, &FineractError{ErrCodeSerialization, &rawMessage}
	}

	return &response, err
}
