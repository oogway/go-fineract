package fineractor

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
)

func (mockClient *MockClient) FundIncrement(fv FundIncrementRequest) (*FundIncrementResponse, error) {
	jsonResp, err := ioutil.ReadFile(path.Join(mockClient.DirectoryPath, "fund_increment.json"))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var response FundIncrementResponse
	json.Unmarshal(jsonResp, &response)
	return &response, nil
}

func (mockClient *MockClient) FundDecrement(request FundDecrementRequest) (*FundDecrementResponse, error) {
	jsonResp, err := ioutil.ReadFile(path.Join(mockClient.DirectoryPath, "fund_decrement.json"))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var response FundDecrementResponse
	json.Unmarshal(jsonResp, &response)
	return &response, nil
}

func (mockClient *MockClient) GetFundValue(request FundValueRequest) (*FundValueResponse, error) {
	jsonResp, err := ioutil.ReadFile(path.Join(mockClient.DirectoryPath, "fund_value.json"))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var response FundValueResponse
	json.Unmarshal(jsonResp, &response)
	return &response, nil
}

func (mockClient *MockClient) GetFunds(request FundsRequest) (*FundsResponse, error) {
	return nil, nil
}
