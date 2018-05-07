package fineractor

type Fineractor interface {
	FundIncrement(request FundIncrementRequest) (*FundIncrementResponse, error)
	FundDecrement(request FundDecrementRequest) (*FundDecrementResponse, error)
	GetFundValue(request FundValueRequest) (*FundValueResponse, error)
	GetFunds(request FundsRequest) (*FundsResponse, error)
}
