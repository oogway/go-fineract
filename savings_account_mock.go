package fineractor

import (
	"log"
)

func (mockClient NewMockClient) FundIncrement(fv FundIncrementRequest) (FundIncrementResponse, error) {
	log.Println("Fineractor increment fund called for NewMockClient")
	return FundIncrementResponse{}, nil
}

func (client NewMockClient) FundDecrement(request FundDecrementRequest) (FundDecrementResponse, error) {
	return FundDecrementResponse{}, nil
}

func (client NewMockClient) GetFundValue(request FundValueRequest) (FundValueResponse, error) {
	return FundValueResponse{}, nil
}

func (client NewMockClient) GetFundAvailablity(request FundAvailablityRequest) (FundAvailablityResponse, error) {
	return FundAvailablityResponse{}, nil
}

func (client NewMockClient) GetFunds(request FundsRequest) (FundsResponse, error) {
	return FundsResponse{}, nil
}
