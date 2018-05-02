package fineractor

type Fineractor interface {
	FundIncrement(request FundIncrementRequest) (FundIncrementResponse, error)
	FundDecrement(request FundDecrementRequest) (FundDecrementResponse, error)
	GetFundValue(request FundValueRequest) (FundValueResponse, error)
	GetFundAvailablity(request FundAvailablityRequest) (FundAvailablityResponse, error)
	GetFunds(request FundsRequest) (FundsResponse, error)
}
