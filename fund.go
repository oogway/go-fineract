package fineractor

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	rh "github.com/tsocial/credit-line-common/retryable-http"
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

type FundDecrementRequest struct{}

type FundDecrementResponse struct{}

type FundValueRequest struct{}

type FundValueResponse struct{}

type FundAvailablityRequest struct{}

type FundAvailablityResponse struct{}

type FundsRequest struct{}

type FundsResponse struct{}

func (client *Client) FundIncrement(request FundIncrementRequest) (FundIncrementResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		log.Println(err)
		return FundIncrementResponse{}, err
	}
	req, err := http.NewRequest("POST", client.HostName+request.Id+"/transactions?command=deposit", bytes.NewBuffer(b))
	if err != nil {
		log.Println(err)
		return FundIncrementResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("fineract-platform-tenantid", "default")
	// TODO: Add a better way to generate this authorisation
	req.Header.Set("Authorization", "Basic bWlmb3M6cGFzc3dvcmQ=")

	respTry, errTry := rh.Try(3, func() (interface{}, error) {
		return client.HttpClient.Do(req)
	})
	if errTry != nil {
		return FundIncrementResponse{}, errors.New(errTry.Error())
	}
	resp := respTry.(*http.Response)
	defer resp.Body.Close()

	var response FundIncrementResponse
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, err
}

func (client *Client) FundDecrement(request FundDecrementRequest) (FundDecrementResponse, error) {
	return FundDecrementResponse{}, nil
}

func (client *Client) GetFundValue(request FundValueRequest) (FundValueResponse, error) {
	return FundValueResponse{}, nil
}

func (client *Client) GetFundAvailablity(request FundAvailablityRequest) (FundAvailablityResponse, error) {
	return FundAvailablityResponse{}, nil
}

func (client *Client) GetFunds(request FundsRequest) (FundsResponse, error) {
	return FundsResponse{}, nil
}
