package fineract

import (
	"net/http"
	"net/url"
	"path"
	"strconv"
	"fmt"
)

type LoanContract struct {
	LoanContract string `json:"loan_contract"`
	LoanLocale   string `json:"loan_locale"`
	SignedAt     string `json:"signed_at"`
	Locale       string `json:"locale"`
	SignStatus   bool   `json:"sign_status"`
	DateFormat   string `json:"dateFormat"`
}

type CreateLoanContractRequest struct {
	LoanId       int64
	LoanContract LoanContract
}

const (
	LoanContractTableName = "datatables/m_loan_contract"
)

type CreateLoanContractResponse struct {
	OfficeId   int64 `json:"officeId"`
	ClientId   int64 `json:"clientId"`
	LoanId     int64 `json:"loanId"`
	ResourceID int64 `json:"resourceId"`
}

func (client *Client) CreateLoanContract(r *CreateLoanContractRequest) (*CreateLoanContractResponse, error) {
	if r.LoanId <= 0 {
		return nil, fmt.Errorf("Invalid loan id")
	}
	tempPath, err := url.Parse(contractURL(r.LoanId))
	if err != nil {
		return nil, err
	}
	response := CreateLoanContractResponse{}

	err = client.MakeRequest(http.MethodPost, client.HostName.ResolveReference(tempPath).String(), r.LoanContract, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func contractURL(loanId int64) string {
	return path.Join(loanContractPath(), strconv.FormatInt(loanId, 10))
}

func loanContractPath() string {
	return path.Join(baseURL, LoanContractTableName)
}
