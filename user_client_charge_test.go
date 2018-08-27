package fineract

import (
	"testing"
	"time"
	"github.com/stretchr/testify/require"
)

func TestClientChargeMock(t *testing.T) {
	if !testing.Short() {
		t.Skip("Skipped mock tests in long mode")
	}
	client, err := makeClient(true)
	if err != nil {
		t.Fatal(err)
	}
	AddChargeToClientSuite(t, client)
	GetClientChargesSuite(t, client)
}

func AddChargeToClientSuite(t *testing.T, client *Client) {
	t.Run("Add charge to client should success", func(t *testing.T) {
		const (
			clientID = "124"
			chargeID = "2"
		)
		dueDate := time.Now().AddDate(0, 6, 0)

		rowID, err := client.AddChargeToClient(clientID, chargeID, dueDate, "10000.25")
		require.Nil(t, err)
		require.Equal(t, int64(164), rowID)
	})
}

func GetClientChargesSuite(t *testing.T, client *Client) {
	t.Run("Client has multiple charges", func(t *testing.T) {
		charges, err := client.GetClientCharges("100", 0, -1)
		require.Nil(t, err)
		require.Len(t, charges, 2)
		for _, charge := range charges {
			require.NotEmpty(t, charge.ChargeId)
			require.NotEmpty(t, charge.Id)
			require.NotEmpty(t, charge.Name)
			require.NotEmpty(t, charge.DueDate)
			require.NotEmpty(t, charge.Amount)
			require.NotEmpty(t, charge.ChargeTime)
			require.NotNil(t, charge.Currency)
		}
	})

	t.Run("Client has one charges", func(t *testing.T) {
		charges, err := client.GetClientCharges("101", 0, -1)
		require.Nil(t, err)
		require.Len(t, charges, 1)
		for _, charge := range charges {
			require.NotEmpty(t, charge.ChargeId)
			require.NotEmpty(t, charge.Id)
			require.NotEmpty(t, charge.Name)
			require.NotEmpty(t, charge.DueDate)
			require.NotEmpty(t, charge.Amount)
			require.NotEmpty(t, charge.ChargeTime)
			require.NotNil(t, charge.Currency)
		}
	})

	t.Run("Client has no charge", func(t *testing.T) {
		charges, err := client.GetClientCharges("102", 0, -1)
		require.Nil(t, err)
		require.Len(t, charges, 0)
	})
}
