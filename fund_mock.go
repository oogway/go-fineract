package fineractor

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
)

func (mockClient *MockClient) GetFieldsMap() map[string]interface{} {
	fields := make(map[string]interface{})
	return fields
}

func (mockClient *MockClient) MakeRequest(directoryPath string, fileName string, request interface{}, response interface{}) error {
	jsonResp, err := ioutil.ReadFile(path.Join(directoryPath, fileName))
	if err != nil {
		log.Println(err.Error())
		return err
	}

	if err = json.Unmarshal(jsonResp, &response); err != nil {
		rawMessage := json.RawMessage([]byte(err.Error()))
		return &FineractError{ErrCodeSerialization, &rawMessage}
	}

	return nil
}

func (mockClient *MockClient) FundIncrement(fundId string, request FundIncrementRequest) (*FundIncrementResponse, error) {
	var response *FundIncrementResponse
	if err := mockClient.MakeRequest(mockClient.DirectoryPath, "fund_increment.json", nil, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (mockClient *MockClient) FundDecrement(fundId string, request FundDecrementRequest) (*FundDecrementResponse, error) {
	var response *FundDecrementResponse
	if err := mockClient.MakeRequest(mockClient.DirectoryPath, "fund_decrement.json", nil, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (mockClient *MockClient) GetFundValue(fundId string, request FundValueRequest) (*FundValueResponse, error) {
	var response *FundValueResponse
	if err := mockClient.MakeRequest(mockClient.DirectoryPath, "fund_value.json", nil, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (mockClient *MockClient) GetFunds(request FundsRequest) (*FundsResponse, error) {
	return nil, nil
}
