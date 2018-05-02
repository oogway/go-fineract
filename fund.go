package fineractor

import (
	"log"
)

type FundIncrementRequest struct {
	Locale            string `json:"locale"`
	DateFormat        string `json:"dateFormat"`
	TransactionDate   string `json:"transactionDate"`
	TransactionAmount string `json:"transactionAmount"`
	PaymentTypeId     string `json:"paymentTypeId"`
}

type FundIncrementResponse struct{}

type FundDecrementRequest struct{}

type FundDecrementResponse struct{}

type FundValueRequest struct{}

type FundValueResponse struct{}

type FundAvailablityRequest struct{}

type FundAvailablityResponse struct{}

type FundsRequest struct{}

type FundsResponse struct{}

func (client Client) FundIncrement(request FundIncrementRequest) (FundIncrementResponse, error) {
	log.Println("Fineractor increment fund called for Client")
	return FundIncrementResponse{}, nil
}

func (client Client) FundDecrement(request FundDecrementRequest) (FundDecrementResponse, error) {
	return FundDecrementResponse{}, nil
}

func (client Client) GetFundValue(request FundValueRequest) (FundValueResponse, error) {
	return FundValueResponse{}, nil
}

func (client Client) GetFundAvailablity(request FundAvailablityRequest) (FundAvailablityResponse, error) {
	return FundAvailablityResponse{}, nil
}

func (client Client) GetFunds(request FundsRequest) (FundsResponse, error) {
	return FundsResponse{}, nil
}
