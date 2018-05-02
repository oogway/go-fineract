package fineractor

import (
	"log"
)

func (mockClient MockClient) FundIncrement(fv FundIncrementRequest) (FundIncrementResponse, error) {
	log.Println("Fineractor increment fund called for NewMockClient")
	return FundIncrementResponse{}, nil
}

func (client MockClient) FundDecrement(request FundDecrementRequest) (FundDecrementResponse, error) {
	return FundDecrementResponse{}, nil
}

func (client MockClient) GetFundValue(request FundValueRequest) (FundValueResponse, error) {
	return FundValueResponse{}, nil
}

func (client MockClient) GetFundAvailablity(request FundAvailablityRequest) (FundAvailablityResponse, error) {
	return FundAvailablityResponse{}, nil
}

func (client MockClient) GetFunds(request FundsRequest) (FundsResponse, error) {
	return FundsResponse{}, nil
}
