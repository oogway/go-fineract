package fineract

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestGetAllCharges(t *testing.T) {
	client, err := makeClient(true)
	if err != nil {
		t.Fatal(err)
	}

	charges, err := client.GetAllCharges()
	require.Nil(t, err)
	require.Len(t, charges, 3)
	for _, charge := range charges {
		require.NotEmpty(t, charge.Id)
		require.NotEmpty(t, charge.Name)
		require.NotEmpty(t, charge.Amount)
		require.NotEmpty(t, charge.ChargeTime)
		require.NotEmpty(t, charge.ChargeAppliesTo)
		require.NotNil(t, charge.Currency)
	}
}
