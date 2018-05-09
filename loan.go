package fineract

import (
	"log"
	"net/url"
)

type CreateLoanRequest struct {
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

type CreateLoanResponse struct {
	Id         float64 `json:"loanId"`
	OfficeId   float64 `json:"officeId"`
	CustomerId float64 `json:"clientId"`
	ResourceId float64 `json:"resourceId"`
}

type UpdateLoanRequest struct {
	EMIId                 string `json:"productId"`
	Term                  string `json:"loanTermFrequency"`
	InstallmentsCount     string `json:"numberOfRepayments"`
	RepaymentFrequency    string `json:"repaymentEvery"`
	InterestRatePerPeriod string `json:"interestRatePerPeriod"`
}

type UpdateLoanResponse struct {
	Id         float64 `json:"loanId"`
	OfficeId   float64 `json:"officeId"`
	CustomerId float64 `json:"clientId"`
	ResourceId float64 `json:"resourceId"`
}

func (client *Client) CreateLoan(request *CreateLoanRequest) (*CreateLoanResponse, error) {
	tempPath, _ := url.Parse("fineract-provider/api/v1/loans")
	path := client.HostName.ResolveReference(tempPath).String()
	var response *CreateLoanResponse
	if err := client.MakeRequest("POST", path, request, &response); err != nil {
		log.Println("Error in create loan: ", err)
		return nil, err
	}

	return response, nil
}
