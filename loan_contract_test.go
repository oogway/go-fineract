package fineract

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestSuiteMockLoanContract(t *testing.T) {
	if !testing.Short() {
		t.Skip("Skipped mock tests in long mode")
	}
	client, err := makeClient(true)
	if err != nil {
		t.Fatal(err)
	}
	SuiteLoanContract(t, client)
}

func SuiteLoanContract(t *testing.T, client *Client) {
	t.Run("TestCreateLoanContract", func(t *testing.T) {
		contract := LoanContract{
			LoanContract: "contract content",
			LoanLocale:   "id",
			SignedAt:     "01 October 2018 11:59",
			Locale:       "en",
			DateFormat:   "dd MMMM yyyy HH:mm",
			SignStatus:   true,
		}
		req := &CreateLoanContractRequest{
			LoanId:       1,
			LoanContract: contract,
		}

		res, err := client.CreateLoanContract(req)
		assert.NoErrorf(t, err, "create loan contract failed with %s", err)
		fmt.Printf("%+v", err)
		fmt.Printf("%+v", res)
		assert.Equal(t, int64(1), res.LoanId)
		assert.Equal(t, int64(1), res.OfficeId)
		assert.Equal(t, int64(1), res.ClientId)
	})
}
