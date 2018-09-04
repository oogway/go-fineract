package fineract

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestSuiteMockAuth(t *testing.T) {
	if !testing.Short() {
		t.Skip("Skipped mock tests in long mode")
	}
	client, err := makeClient(true)
	if err != nil {
		t.Fatal(err)
	}
	AuthSuite(t, client)
}

func AuthSuite(t *testing.T, client *Client) {
	t.Run("Test authentication success", func(t *testing.T) {
		authReq := &AuthRequest{
			Username: "mifos",
			Password: "mifos",
		}

		response, err := client.Auth(authReq)
		if err != nil {
			t.Fatalf(err.Error())
		}
		require.Nil(t, err, "Error should be nil")
		require.NotNil(t, response, "response should not be nil")
		require.NotNil(t, response.Username, "username should not be nil")
		require.NotNil(t, response.OfficeId, "OfficeId should not be nil")
		require.NotNil(t, response.OfficeName, "OfficeName should not be nil")
		require.NotNil(t, response.Permissions, "Permission should not be nil")
		require.NotNil(t, response.Token, "Token should not be nil")
	})
}
