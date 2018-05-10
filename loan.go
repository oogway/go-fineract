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

type GetLoanProductRequest struct{}

type InterestType struct {
	Id float64 `json:"Id"`
}

type GetLoanProductResponse struct {
	Term               float64
	RepaymentFrequency float64      `json:"repaymentEvery"`
	Id                 float64      `json:"Id"`
	InstallmentsCount  float64      `json:"numberOfRepayments"`
	Type               InterestType `json:"interestType"`
	InterestRate       float64      `json:"interestRatePerPeriod"`
}

type GetLoanRequest struct{}

type GetLoanResponse struct {
	Id                            float64 `json:"id"`
	ExternalId                    string  `json:"externalId"`
	ClientId                      float64 `json:"clientId"`
	EMIId                         float64 `json:"loanProductId"`
	Principal                     float64 `json:"principal"`
	Term                          float64 `json:"termFrequency"`
	InstallmentsCount             float64 `json:"numberOfRepayments"`
	RepaymentFrequency            float64 `json:"repaymentEvery"`
	InterestRatePerPeriod         float64 `json:"interestRatePerPeriod"`
	TransactionProcessingStrategy float64 `json:"transactionProcessingStrategyId"`
}

type LoanConfirmRequest struct {
	Locale           string `json:"locale"`
	DateFormat       string `json:"dateFormat"`
	ConfirmationDate string `json:"approvedOnDate"`
}

type LoanConfirmResponse struct{}

type LoanDisburseRequest struct {
	Locale           string `json:"locale"`
	DateFormat       string `json:"dateFormat"`
	DisbursementDate string `json:"actualDisbursementDate"`
}

type LoanDisburseResponse struct{}

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

func (client *Client) GetLoanProduct(loanProductId string, request *GetLoanProductRequest) (*GetLoanProductResponse, error) {
	tempPath, _ := url.Parse(path.Join("fineract-provider/api/v1/loanproducts", loanProductId))
	path := client.HostName.ResolveReference(tempPath).String()
	var response *GetLoanProductResponse
	if err := client.MakeRequest("GET", path, nil, &response); err != nil {
		log.Println("Error in geting the loan product: ", err)
		return nil, err
	}

	return response, nil
}

func (client *Client) GetLoan(loanId string, request *GetLoanRequest) (*GetLoanResponse, error) {
	tempPath, _ := url.Parse(path.Join("fineract-provider/api/v1/loans", loanId))
	path := client.HostName.ResolveReference(tempPath).String()
	var response *GetLoanResponse
	if err := client.MakeRequest("GET", path, nil, &response); err != nil {
		log.Println("Error in geting the loan: ", err)
		return nil, err
	}

	return response, nil
}

func (client *Client) LoanConfirm(loanId string, request *LoanConfirmRequest) (*LoanConfirmResponse, error) {
	tempPath, _ := url.Parse(path.Join("fineract-provider/api/v1/loans", loanId+"/?command=approve"))
	path := client.HostName.ResolveReference(tempPath).String()
	var response *LoanConfirmResponse
	if err := client.MakeRequest("POST", path, request, &response); err != nil {
		log.Println("Error in confirming loan: ", err)
		return nil, err
	}

	return response, nil
}

func (client *Client) LoanDisburse(loanId string, request *LoanDisburseRequest) (*LoanDisburseResponse, error) {
	tempPath, _ := url.Parse(path.Join("fineract-provider/api/v1/loans", loanId+"/?command=disburse"))
	path := client.HostName.ResolveReference(tempPath).String()
	var response *LoanDisburseResponse
	if err := client.MakeRequest("POST", path, request, &response); err != nil {
		log.Println("Error in disbursal of loan: ", err)
		return nil, err
	}

	return response, nil
}
