package fineractor

import (
	"fmt"
	"testing"
)

func TestNewClient(t *testing.T) {
	hostName := "https://demo.openmf.org/fineract-provider/api/v1/savingsaccounts/"
	userName := "mifos"
	password := "password"

	t.Run("Should decrement fund", func(t *testing.T) {
		client, err := NewClient(hostName, userName, password, FineractOption{
			Transport: &MockTransport{DirectoryPath: "testdata"},
		})
		if err != nil {
			t.Fatalf("Cannot create new client: %v", err)
		}

		clientNext, err := NewClient(hostName, userName, password, FineractOption{})

		if fmt.Sprintf("%p", client) != fmt.Sprintf("%p", clientNext) {
			t.Fatal("Client should be initialised only once and same client should be returned on next call")
		}
	})
}

func TestGetFundValue(t *testing.T) {
	t.Run("Should decrement fund", func(t *testing.T) {
		client, err := NewClient("https://demo.openmf.org/fineract-provider/api/v1/savingsaccounts/", "mifos", "password", FineractOption{
			Transport: &MockTransport{DirectoryPath: "testdata"},
		})
		if err != nil {
			t.Fatalf("Cannot create new client: %v", err)
		}

		resp, err := client.GetFundValue("1884", nil)
		if err != nil {
			t.Fatalf("Cannot get the fund value: %v", err)
		}

		if resp.Statement.Amount != 4410 {
			t.Fatalf("Cannot get fund value: %v", err)
		}
	})
}

func TestFundIncrement(t *testing.T) {
	t.Run("Should decrement fund", func(t *testing.T) {
		client, err := NewClient("https://demo.openmf.org/fineract-provider/api/v1/savingsaccounts/", "mifos", "password", FineractOption{
			Transport: &MockTransport{DirectoryPath: "testdata"},
		})
		if err != nil {
			t.Fatalf("Cannot create new client: %v", err)
		}

		req := FundIncrementRequest{}
		if _, err = client.FundIncrement("1884", &req); err != nil {
			t.Fatalf("Cannot get the fund value: %v", err)
		}
	})
}

func TestDecrementIncrement(t *testing.T) {
	t.Run("Should decrement fund", func(t *testing.T) {
		client, err := NewClient("https://demo.openmf.org/fineract-provider/api/v1/savingsaccounts/", "mifos", "password", FineractOption{
			Transport: &MockTransport{DirectoryPath: "testdata"},
		})
		if err != nil {
			t.Fatalf("Cannot create new client: %v", err)
		}

		req := FundDecrementRequest{}
		if _, err = client.FundDecrement("1884", &req); err != nil {
			t.Fatalf("Cannot get the fund value: %v", err)
		}
	})
}
