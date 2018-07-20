package fineract

import (
	"fmt"
	"log"
	"testing"

	"github.com/bmizerany/assert"
)

func TestSuiteLoanMock(t *testing.T) {
	if !testing.Short() {
		t.Skip("Skipped mock tests in long mode")
	}
	client, err := MakeClient(true)
	if err != nil {
		t.Fatal(err)
	}
	LoanSuite(t, client, "12")
}

func TestSuiteLoan(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipped integrated tests in short mode")
	}
	client, err := MakeClient(false)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.GetLoanProducts(&GetLoanProductsRequest{})
	if err != nil {
		t.Fatalf("Cannot get loan products: %v", err)
	}

	LoanSuite(t, client, fmt.Sprintf("%v", resp.LoanProducts[0].Id))
}

func LoanSuite(t *testing.T, client *Client, loanProductId string) {

	t.Run("TestGetLoanProduct with service & penalty charges", func(t *testing.T) {
		resp, err := client.GetLoanProduct(loanProductId, &GetLoanProductRequest{})
		if err != nil {
			t.Fatalf("Cannot get the loan product: %v", err)
		}
		log.Println(resp)
		if len(resp.Charges) == 0 {
			t.Fatalf("No charge found")
		}
		var penaltyCharge, serviceCharge = false, false
		for _, charge := range resp.Charges {
			if charge.ChargeTime.Code == "chargeTimeType.overdueInstallment" {
				penaltyCharge = true
			} else if charge.ChargeTime.Code == "chargeTimeType.instalmentFee" {
				serviceCharge = true
			}
		}
		assert.Equal(t, penaltyCharge, true)
		assert.Equal(t, serviceCharge, true)
	})
}
