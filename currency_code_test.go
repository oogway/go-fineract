package fineract

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestSuiteCurrencyCode(t *testing.T) {
	client, err := MakeClient(true)
	if err != nil {
		t.Fatal(err)
	}
	currencyCode, err := client.GetCurrencyCode()
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(currencyCode.SelectedCurrencyOptions))
	assert.Equal(t, "IDR", currencyCode.SelectedCurrencyOptions[0].Code)
}
