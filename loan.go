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
	Id                 float64      `json:"Id"`
	RepaymentFrequency float64      `json:"repaymentEvery"`
	InstallmentsCount  float64      `json:"numberOfRepayments"`
	Type               InterestType `json:"interestType"`
	InterestRate       float64      `json:"interestRatePerPeriod"`
}

type GetLoanProductsRequest struct{}

type GetLoanProductsResponse struct {
	LoanProducts []GetLoanProductResponse
}

type GetLoanRequest struct{}

type Status struct {
	Code                string `json:"code"`
	Value               string `json:"Value"`
	PendingApproval     bool   `json:"pendingApproval"`
	WaitingForDisbursal bool   `json:"waitingForDisbursal"`
}

type GetLoanResponse struct {
	Id                            int64   `json:"id"`
	ExternalId                    string  `json:"externalId"`
	ClientId                      float64 `json:"clientId"`
	EMIId                         float64 `json:"loanProductId"`
	Principal                     float64 `json:"principal"`
	Term                          float64 `json:"termFrequency"`
	InstallmentsCount             float64 `json:"numberOfRepayments"`
	RepaymentFrequency            float64 `json:"repaymentEvery"`
	InterestRatePerPeriod         float64 `json:"interestRatePerPeriod"`
	TransactionProcessingStrategy float64 `json:"transactionProcessingStrategyId"`
	Lstatus                       Status  `json:"status"`
}

type GetAllLoanRequest struct {
	ClientId int64
}

type GetAllLoanResponse struct {
	Loans []GetLoanResponse `json:"pageItems"`
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

type LoanCalculateScheduleRequest struct {
	Locale                          string  `json:"locale"`
	DateFormat                      string  `json:"dateFormat"`
	LoanAmount                      string  `json:"principal,omitempty"`
	ProductId                       string  `json:"productId,omitempty"`
	ClientId                        string  `json:"clientId,omitempty"`
	LoanTermFrequency               uint64  `json:"loanTermFrequency,omitempty"`
	LoanTermFrequencyType           float64 `json:"loanTermFrequencyType,omitempty"`
	NumberOfRepayments              uint64  `json:"numberOfRepayments,omitempty"`
	RepaymentEvery                  uint64  `json:"repaymentEvery,omitempty"`
	RepaymentFrequencyType          uint64  `json:"repaymentFrequencyType,omitempty"`
	AmortizationType                uint64  `json:"amortizationType,omitempty"`
	InterestRatePerPeriod           float64 `json:"interestRatePerPeriod"`
	InterestType                    float64 `json:"interestType"`
	InterestCalculationPeriodType   uint64  `json:"interestCalculationPeriodType,omitempty"`
	ExpectedDisbursementDate        string  `json:"expectedDisbursementDate,omitempty"`
	SubmittedOnDate                 string  `json:"submittedOnDate,omitempty"`
	TransactionProcessingStrategyId uint64  `json:"transactionProcessingStrategyId,omitempty"`
	LoanType                        string  `json:"loanType,omitempty"`
}

type LoanPeriod struct {
	Period                          uint64  `json:"period,omitempty"`
	FromDate                        []uint  `json:"fromDate,omitempty"`
	DueDate                         []uint  `json:"dueDate,omitempty"`
	PrincipalDisbursed              float64 `json:"principalDisbursed,omitempty"`
	PrincipalLoanBalanceOutstanding float64 `json:"principalLoanBalanceOutstanding,omitempty"`
	FeeChargesOutstanding           float64 `json:"feeChargesOutstanding,omitempty"`
	DaysInPeriod                    uint64  `json:"daysInPeriod,omitempty"`
	PrincipalOriginalDue            float64 `json:"principalOriginalDue,omitempty"`
	PrincipalDue                    float64 `json:"principalDue,omitempty"`
	PrincipalOutstanding            float64 `json:"principalOutstanding,omitempty"`
	InterestOriginalDue             float64 `json:"interestOriginalDue,omitempty"`
	InterestDue                     float64 `json:"interestDue,omitempty"`
	InterestOutstanding             float64 `json:"interestOutstanding,omitempty"`
	FeeChargesDue                   float64 `json:"feeChargesDue,omitempty"`
	PenaltyChargesDue               float64 `json:"penaltyChargesDue,omitempty"`
	TotalOriginalDueForPeriod       float64 `json:"totalOriginalDueForPeriod,omitempty"`
	TotalDueForPeriod               float64 `json:"totalDueForPeriod,omitempty"`
	TotalPaidForPeriod              float64 `json:"totalPaidForPeriod,omitempty"`
	TotalOutstandingForPeriod       float64 `json:"totalOutstandingForPeriod,omitempty"`
	TotalActualCostOfLoanForPeriod  float64 `json:"totalActualCostOfLoanForPeriod,omitempty"`
	TotalInstallmentAmountForPeriod float64 `json:"totalInstallmentAmountForPeriod,omitempty"`
}

type LoanCalculateScheduleResponse struct {
	TotalRepaymentExpected float64       `json:"totalRepaymentExpected,omitempty"`
	TotalInterestCharged   float64       `json:"totalInterestCharged,omitempty"`
	LoanTermInDays         uint64        `json:"loanTermInDays,omitempty"`
	Periods                []*LoanPeriod `json:"periods,omitempty"`
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

func (client *Client) GetAllLoan(request *GetAllLoanRequest) (*GetAllLoanResponse, error) {
	tempPath, _ := url.Parse("fineract-provider/api/v1/loans")
	path := client.HostName.ResolveReference(tempPath).String()
	var response *GetAllLoanResponse
	if err := client.MakeRequest("GET", path, nil, &response); err != nil {
		log.Println("Error in geting the loan: ", err)
		return nil, err
	}

	var clientLoans []GetLoanResponse
	for _, loan := range response.Loans {
		if loan.ClientId == float64(request.ClientId) {
			clientLoans = append(clientLoans, loan)
		}
	}
	return &GetAllLoanResponse{
		Loans: clientLoans,
	}, nil
}

func (client *Client) LoanConfirm(loanId string, request *LoanConfirmRequest) (*LoanConfirmResponse, error) {
	tempPath, _ := url.Parse(path.Join("fineract-provider/api/v1/loans", loanId+"?command=approve"))
	path := client.HostName.ResolveReference(tempPath).String()
	var response *LoanConfirmResponse
	if err := client.MakeRequest("POST", path, request, &response); err != nil {
		log.Println("Error in confirming loan: ", err)
		return nil, err
	}

	return response, nil
}

func (client *Client) LoanDisburse(loanId string, request *LoanDisburseRequest) (*LoanDisburseResponse, error) {
	tempPath, _ := url.Parse(path.Join("fineract-provider/api/v1/loans", loanId+"?command=disburse"))
	path := client.HostName.ResolveReference(tempPath).String()
	var response *LoanDisburseResponse
	if err := client.MakeRequest("POST", path, request, &response); err != nil {
		log.Println("Error in disbursal of loan: ", err)
		return nil, err
	}

	return response, nil
}

func (client *Client) LoanCalculateSchedule(request *LoanCalculateScheduleRequest) (*LoanCalculateScheduleResponse, error) {
	tempPath, _ := url.Parse("fineract-provider/api/v1/loans?command=calculateLoanSchedule")
	path := client.HostName.ResolveReference(tempPath).String()
	var response *LoanCalculateScheduleResponse
	if err := client.MakeRequest("POST", path, request, &response); err != nil {
		log.Println("Error in disbursal of loan: ", err)
		return nil, err
	}

	return response, nil
}

func (client *Client) GetLoanProducts(request *GetLoanProductsRequest) (*GetLoanProductsResponse, error) {
	tempPath, _ := url.Parse("fineract-provider/api/v1/loanproducts")
	path := client.HostName.ResolveReference(tempPath).String()
	var response []GetLoanProductResponse
	if err := client.MakeRequest("GET", path, nil, &response); err != nil {
		log.Println("Error in geting the loan products: ", err)
		return nil, err
	}

	return &GetLoanProductsResponse{LoanProducts: response}, nil
}
