package fineract

import (
	"fmt"
	"testing"
	"time"

	"github.com/bmizerany/assert"
)

const (
	LendingName = "TSFund"
)

var (
	Lending string
)

func TestSuiteMock(t *testing.T) {
	if !testing.Short() {
		t.Skip("Skipped mock tests in long mode")
	}
	client, err := makeClient(true)
	if err != nil {
		t.Fatal(err)
	}

	fundId := "144"
	Suite(t, client, fundId)
}

func TestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipped integrated tests in short mode")
	}
	client, err := makeClient(false)
	if err != nil {
		t.Fatal(err)
	}

	//retrieve fundType
	id, err := client.GetFundType(&GetFundTypeRequest{Name: LendingName})
	if err != nil {
		t.Fatalf("retrieve fund type: %v", err)
	}
	Lending = toString(id)

	//retrieve fundID from GetFunds
	req := FundsRequest{Type: Lending}
	resp, err := client.GetFunds(&req)
	if err != nil {
		t.Fatalf("retrieve list of fund(s): %v", err)
	}
	if len(resp.Fund) == 0 {
		t.Fatalf("no fund retrieved, atleast one fund should exists")
	}
	fundId := fmt.Sprintf("%v", resp.Fund[0].Id)

	Suite(t, client, fundId)
	ISuite(t, client, fundId)
}

func Suite(t *testing.T, client *Client, fundId string) {
	var fundInitialBalance float64
	var accntId string

	t.Run("TestGetFundValue", func(t *testing.T) {
		resp, currency, err := client.GetFundValue(fundId)
		if err != nil {
			t.Fatalf("Cannot get the fund value: %v", err)
		}
		fundInitialBalance = resp
		assert.NotEqual(t, currency, "")
	})

	t.Run("TestGetFundAccountId", func(t *testing.T) {
		resp, err := client.GetFundAccountId(fundId)
		if err != nil || resp.PrincipalAccountId == "" {
			t.Fatalf("retrieve accountId(s) for fund: %v", err)
		}
		accntId = resp.PrincipalAccountId
	})

	t.Run("TestGetFunds", func(t *testing.T) {
		resp, err := client.GetFunds(&FundsRequest{Type: Lending})
		if err != nil || resp.TotalFilteredRecords == 0 {
			t.Fatalf("retrieve list of fund(s): %v", err)
		}
		assert.NotEqual(t, resp.Fund[0].Currency.Code, "")
	})

	t.Run("TestGetFund", func(t *testing.T) {
		resp, err := client.GetFund(fundId)
		if err != nil || resp.Id == 0 {
			t.Fatalf("retrieve core details of fund: %v", err)
		}
	})

	t.Run("TestGetAccount", func(t *testing.T) {
		resp, err := client.GetAccount(accntId)
		if err != nil || resp.ProductId == 0 {
			t.Fatalf("retrieve core details of account: %v", err)
		}
		assert.NotEqual(t, resp.Currency.Code, "")
	})

	t.Run("TestGetPaymentType", func(t *testing.T) {
		resp, err := client.GetPaymentType(&GetPaymentTypeRequest{})
		if err != nil {
			t.Fatalf("Cannot retrieve payment types: %v", err)
		}
		if len(resp.PaymentMethod) == 0 {
			t.Fatalf("No payment type found, atleast one is required")
		}
	})

	t.Run("TestGetFundAccounts", func(t *testing.T) {
		if resp, err := client.GetFundAccounts(fundId); err != nil || len(resp.FundAccount) == 0 {
			t.Fatalf("retrieve fund account: %v", err)
		}
	})
}

func ISuite(t *testing.T, client *Client, fundId string) {
	var txAmount float64 = 500
	var accountId string

	resp, err := client.GetPaymentType(&GetPaymentTypeRequest{})
	if err != nil {
		t.Fatalf("failed to get payment types: %v", err)
	}
	if len(resp.PaymentMethod) == 0 {
		t.Fatalf("no payment type found, atleast one is required")
	}
	paymentId := fmt.Sprintf("%v", resp.PaymentMethod[0].Id)

	//getAccountId
	response, err := client.GetFundAccounts(fundId)
	if err != nil {
		t.Fatalf("failed to get accountId: %v", err)
	}

	for _, cursor := range response.FundAccount {
		if cursor.ProductName == toString(Principal) && cursor.Status.Value == active {
			accountId = cursor.AccountNo
			break
		}
	}
	if accountId == "" {
		t.Fatalf("failed to get accountId: %v", err)
	}

	t.Run("TestFundIncrement", func(t *testing.T) {
		before, currency, err := client.GetFundValue(fundId)
		if err != nil {
			t.Fatalf("Cannot get the fund value: %v", err)
		}

		//increment
		req := &TxRequest{
			Locale:            "en",
			DateFormat:        "dd MMMM yyyy",
			TransactionDate:   time.Now().Format("02 January 2006"),
			TransactionAmount: fmt.Sprintf("%v", txAmount),
			PaymentTypeId:     paymentId,
		}
		_, err = client.AccountDeposit(accountId, req)
		if err != nil {
			t.Fatalf("Could not increment the fund value: %v", err)
		}

		after, currency, err := client.GetFundValue(fundId)
		if err != nil {
			t.Fatalf("Cannot get the fund value: %v", err)
		}

		assertEqual(t, before+txAmount, after, "Fund balance was not incremented")
		assert.NotEqual(t, currency, "")
	})

	t.Run("TestFundDecrement", func(t *testing.T) {
		before, currency, err := client.GetFundValue(fundId)
		if err != nil {
			t.Fatalf("Cannot get the fund value: %v", err)
		}

		//decrement
		req := &TxRequest{
			Locale:            "en",
			DateFormat:        "dd MMMM yyyy",
			TransactionDate:   time.Now().Format("02 January 2006"),
			TransactionAmount: fmt.Sprintf("%v", txAmount),
			PaymentTypeId:     paymentId,
		}
		_, err = client.AccountWithdraw(accountId, req)
		if err != nil {
			t.Fatalf("Could not decrement the fund value: %v", err)
		}

		after, currency, err := client.GetFundValue(fundId)
		if err != nil {
			t.Fatalf("Cannot get the fund value: %v", err)
		}

		assertEqual(t, before-txAmount, after, "Fund balance was not decremented")
		assert.NotEqual(t, currency, "")
	})
}

func assertEqual(t *testing.T, a interface{}, b interface{}, msg string) {
	if a != b {
		t.Fatalf("%s != %s : %s", a, b, msg)
	}
}
