package fineract

import (
	"fmt"
	"testing"
	"time"
)

func makeClient(mock bool) (*Client, error) {
	if mock {
		return NewClient("", "", "", FineractOption{
			Transport: &MockTransport{DirectoryPath: "testdata"},
		})
	}
	return NewClient("https://"+fineractHost, fineractUser, fineractPassword, FineractOption{SkipVerify: true})
}

func TestSuiteMock(t *testing.T) {
	if !testing.Short() {
		t.Skip("Skipped mock tests in long mode")
	}
	client, err := makeClient(true)
	if err != nil {
		t.Fatal(err)
	}

	fundId := "1884"
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

	//retrieve fundID from GetFunds
	req := FundsRequest{}
	resp, err := client.GetFunds(&req)
	if err != nil {
		t.Fatalf("retrieve list of fund(s): %v", err)
	}
	if len(resp.FundDetail) == 0 {
		t.Fatalf("No fund retrieved, atleast one fund should exists")
	}
	fundId := fmt.Sprintf("%v", resp.FundDetail[0].Id)

	Suite(t, client, fundId)
	ISuite(t, client, fundId)
}

func Suite(t *testing.T, client *Client, fundId string) {
	var fundInitialBalance float64

	t.Run("TestGetFundValue", func(t *testing.T) {
		resp, err := client.GetFundValue(fundId, nil)
		if err != nil {
			t.Fatalf("Cannot get the fund value: %v", err)
		}

		fundInitialBalance = resp.Statement.Amount
	})

	t.Run("TestGetFunds", func(t *testing.T) {
		req := FundsRequest{}
		if _, err := client.GetFunds(&req); err != nil {
			t.Fatalf("retrieve list of fund(s): %v", err)
		}
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
}

func ISuite(t *testing.T, client *Client, fundId string) {
	var txAmount float64 = 500

	resp, err := client.GetPaymentType(&GetPaymentTypeRequest{})
	if err != nil {
		t.Fatalf("failed to get payment types: %v", err)
	}
	if len(resp.PaymentMethod) == 0 {
		t.Fatalf("no payment type found, atleast one is required")
	}
	paymentId := fmt.Sprintf("%v", resp.PaymentMethod[0].Id)

	t.Run("TestFundIncrement", func(t *testing.T) {
		before, err := client.GetFundValue(fundId, nil)
		if err != nil {
			t.Fatalf("Cannot get the fund value: %v", err)
		}

		//increment
		req := &FundIncrementRequest{
			Locale:            "en",
			DateFormat:        "dd MMMM yyyy",
			TransactionDate:   time.Now().Format("02 January 2006"),
			TransactionAmount: fmt.Sprintf("%v", txAmount),
			PaymentTypeId:     paymentId,
		}
		_, err = client.FundIncrement(fundId, req)
		if err != nil {
			t.Fatalf("Could not increment the fund value: %v", err)
		}

		after, err := client.GetFundValue(fundId, nil)
		if err != nil {
			t.Fatalf("Cannot get the fund value: %v", err)
		}

		assertEqual(t, before.Statement.Amount+txAmount, after.Statement.Amount, "Fund balance was not incremented")
	})

	t.Run("TestFundDecrement", func(t *testing.T) {
		before, err := client.GetFundValue(fundId, nil)
		if err != nil {
			t.Fatalf("Cannot get the fund value: %v", err)
		}

		//increment
		req := &FundDecrementRequest{
			Locale:            "en",
			DateFormat:        "dd MMMM yyyy",
			TransactionDate:   time.Now().Format("02 January 2006"),
			TransactionAmount: fmt.Sprintf("%v", txAmount),
			PaymentTypeId:     paymentId,
		}
		_, err = client.FundDecrement(fundId, req)
		if err != nil {
			t.Fatalf("Could not decrement the fund value: %v", err)
		}

		after, err := client.GetFundValue(fundId, nil)
		if err != nil {
			t.Fatalf("Cannot get the fund value: %v", err)
		}

		assertEqual(t, before.Statement.Amount-txAmount, after.Statement.Amount, "Fund balance was not decremented")
	})
}

func assertEqual(t *testing.T, a interface{}, b interface{}, msg string) {
	if a != b {
		t.Fatalf("%s != %s : %s", a, b, msg)
	}
}
