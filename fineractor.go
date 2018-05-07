package fineractor

type Fineractor interface {
	FundIncrement(fundId string, request FundIncrementRequest) (*FundIncrementResponse, error)
	FundDecrement(fundId string, request FundDecrementRequest) (*FundDecrementResponse, error)
	GetFundValue(fundId string, request FundValueRequest) (*FundValueResponse, error)
	GetFunds(request FundsRequest) (*FundsResponse, error)
	GetFieldsMap() map[string]interface{}
	MakeRequest(string, string, interface{}, interface{}) error
}
