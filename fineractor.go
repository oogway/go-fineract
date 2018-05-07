package fineractor

type Fineractor interface {
	FundIncrement(string, FundIncrementRequest) (*FundIncrementResponse, error)
	FundDecrement(string, FundDecrementRequest) (*FundDecrementResponse, error)
	GetFundValue(string, FundValueRequest) (*FundValueResponse, error)
	GetFunds(FundsRequest) (*FundsResponse, error)
	GetFieldsMap() map[string]interface{}
	MakeRequest(string, string, interface{}, interface{}) error
}
