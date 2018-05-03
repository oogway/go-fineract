package fineractor

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/meson10/highbrow"
)

type FundIncrementRequest struct {
	Id                string `json:"-"`
	Locale            string `json:"locale"`
	DateFormat        string `json:"dateFormat"`
	TransactionDate   string `json:"transactionDate"`
	TransactionAmount string `json:"transactionAmount"`
	PaymentTypeId     string `json:"paymentTypeId"`
}

type FundIncrementResponse struct {
	OfficeId   float64
	ClientId   float64
	ResourceId float64
}

type FundDecrementRequest struct {
	Id                string `json:"-"`
	Locale            string `json:"locale"`
	DateFormat        string `json:"dateFormat"`
	TransactionDate   string `json:"transactionDate"`
	TransactionAmount string `json:"transactionAmount"`
	PaymentTypeId     string `json:"paymentTypeId"`
}

type FundDecrementResponse struct {
	OfficeId   float64
	ClientId   float64
	ResourceId float64
}

type FundValueRequest struct {
	Id string `json:"-"`
}

type Summary struct {
	Amount           float64 `json:"accountBalance"`
	TotalDeposits    float64 `json:"totalDeposits"`
	TotalWithdrawals float64 `json:"totalWithdrawals"`
}

type FundValueResponse struct {
	AccountNo  string  `json:"accountNo"`
	ClientId   int64   `json:"clientId"`
	ClientName string  `json:"clientName"`
	Statement  Summary `json:"summary"`
}

type FundAvailablityRequest struct{}

type FundAvailablityResponse struct{}

type FundsRequest struct{}

type FundsResponse struct{}

func (client *Client) FundIncrement(request FundIncrementRequest) (*FundIncrementResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	tempPath, _ := url.Parse(path.Join(request.Id, "transactions?command=deposit"))
	path := client.HostName.ResolveReference(tempPath).String()
	req, err := http.NewRequest("POST", path, bytes.NewBuffer(b))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("fineract-platform-tenantid", "default")
	// TODO: Add a better way to generate this authorisation
	req.Header.Set("Authorization", "Basic bWlmb3M6cGFzc3dvcmQ=")

	var resp *http.Response
	errTry := highbrow.Try(5, func() error {
		resp, err = client.HttpClient.Do(req)
		return err
	})
	if errTry != nil {
		return nil, errors.New(errTry.Error())
	}
	defer resp.Body.Close()

	var response FundIncrementResponse
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return &response, err
}

func (client *Client) FundDecrement(request FundDecrementRequest) (*FundDecrementResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	tempPath, _ := url.Parse(path.Join(request.Id, "transactions?command=withdrawal"))
	path := client.HostName.ResolveReference(tempPath).String()
	log.Println(path)
	req, err := http.NewRequest("POST", path, bytes.NewBuffer(b))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("fineract-platform-tenantid", "default")
	// TODO: Add a better way to generate this authorisation
	req.Header.Set("Authorization", "Basic bWlmb3M6cGFzc3dvcmQ=")

	var resp *http.Response
	errTry := highbrow.Try(5, func() error {
		resp, err = client.HttpClient.Do(req)
		return err
	})
	if errTry != nil {
		return nil, errors.New(errTry.Error())
	}
	defer resp.Body.Close()

	var response FundDecrementResponse
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return &response, err
}

func (client *Client) GetFundValue(request FundValueRequest) (*FundValueResponse, error) {
	tempPath, _ := url.Parse(request.Id)
	req, err := http.NewRequest("GET", client.HostName.ResolveReference(tempPath).String(), nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("fineract-platform-tenantid", "default")
	req.Header.Set("Authorization", "Basic bWlmb3M6cGFzc3dvcmQ=")

	var resp *http.Response
	errTry := highbrow.Try(5, func() error {
		resp, err = client.HttpClient.Do(req)
		return err
	})
	if errTry != nil {
		return nil, errors.New(errTry.Error())
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var response FundValueResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return &response, err
}

func (client *Client) GetFundAvailablity(request FundAvailablityRequest) (*FundAvailablityResponse, error) {
	return nil, nil
}

func (client *Client) GetFunds(request FundsRequest) (*FundsResponse, error) {
	return nil, nil
}
