package fineract

import (
	"log"
	"net/url"
	"path"
)

type LoanCreateRequest struct {
	Locale                        string `json:"locale"`
	DateFormat                    string `json:"dateFormat"`
	ExternalId                    string `json:"externalId"`
	ClientId                      string `json:"clientId"`
	EMIId                         string `json:"productId"`
	LoanAmount                    string `json:"principal"`
	Term                          string `json:"loanTermFrequency"`
	TermFrequency                 string `json:"loanTermFrequencyType"`
	Type                          string `json:"loanType"`
	InstallmentsCount             string `json:"numberOfRepayments"`
	RepaymentFrequency            string `json:"repaymentEvery"`
	RepaymentFrequencyType        string `json:"repaymentFrequencyType"`
	InterestRatePerPeriod         string `json:"interestRatePerPeriod"`
	AmortizationType              string `json:"amortizationType"`
	InterestType                  string `json:"interestType"`
	InterestCalculationPeriodType string `json:"interestCalculationPeriodType"`
	TransactionProcessingStrategy string `json:"transactionProcessingStrategyId"`
	ExpectedDisbursementDate      string `json:"expectedDisbursementDate"`
	SubmittedOnDate               string `json:"submittedOnDate"`
	LinkAccountId                 string `json:"linkAccountId"`
}

type LoanCreateResponse struct {
	Id         float64 `json:"loanId"`
	OfficeId   float64 `json:"officeId"`
	CustomerId float64 `json:"clientId"`
	ResourceId float64 `json:"resourceId"`
}

type LoanUpdateRequest struct {
	Locale                string `json:"locale"`
	DateFormat            string `json:"dateFormat"`
	EMIId                 string `json:"productId"`
	Term                  string `json:"loanTermFrequency"`
	InstallmentsCount     string `json:"numberOfRepayments"`
	RepaymentFrequency    string `json:"repaymentEvery"`
	InterestRatePerPeriod string `json:"interestRatePerPeriod"`
}

type LoanUpdateResponse struct {
	Id         float64 `json:"loanId"`
	OfficeId   float64 `json:"officeId"`
	CustomerId float64 `json:"clientId"`
	ResourceId float64 `json:"resourceId"`
}

func (client *Client) LoanCreate(request *LoanCreateRequest) (*LoanCreateResponse, error) {
	tempPath, _ := url.Parse("fineract-provider/api/v1/loans")
	path := client.HostName.ResolveReference(tempPath).String()
	var response *LoanCreateResponse
	if err := client.MakeRequest("POST", path, request, &response); err != nil {
		log.Println("Error in create loan: ", err)
		return nil, err
	}

	return response, nil
}

func (client *Client) LoanUpdate(loanId string, request *LoanUpdateRequest) (*LoanUpdateResponse, error) {
	tempPath, _ := url.Parse(path.Join("fineract-provider/api/v1/loans", loanId))
	path := client.HostName.ResolveReference(tempPath).String()
	var response *LoanUpdateResponse
	if err := client.MakeRequest("PUT", path, request, &response); err != nil {
		log.Println("Error in update loan: ", err)
		return nil, err
	}

	return response, nil
}
