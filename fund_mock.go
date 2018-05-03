package fineractor

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
)

func (mockClient *MockClient) FundIncrement(fv FundIncrementRequest) (FundIncrementResponse, error) {
	jsonResp, err := ioutil.ReadFile(path.Join(mockClient.DirectoryPath, "fund_increment_success.json"))
	if err != nil {
		log.Println(err.Error())
		return FundIncrementResponse{}, err
	}

	var response FundIncrementResponse
	json.Unmarshal(jsonResp, &response)
	return response, nil
}

func (client *MockClient) FundDecrement(request FundDecrementRequest) (FundDecrementResponse, error) {
	jsonResp, err := ioutil.ReadFile(path.Join(mockClient.DirectoryPath, "fund_decrement_success.json"))
	if err != nil {
		log.Println(err.Error())
		return FundDecrementResponse{}, err
	}

	var response FundDecrementResponse
	json.Unmarshal(jsonResp, &response)
	return response, nil
}

func (client *MockClient) GetFundValue(request FundValueRequest) (FundValueResponse, error) {
	return FundValueResponse{}, nil
}

func (client *MockClient) GetFundAvailablity(request FundAvailablityRequest) (FundAvailablityResponse, error) {
	return FundAvailablityResponse{}, nil
}

func (client *MockClient) GetFunds(request FundsRequest) (FundsResponse, error) {
	return FundsResponse{}, nil
}
