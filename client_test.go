package fineract

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateClient(t *testing.T) {
	client, err := makeClient(true)
	require.Nil(t, err)

	clientReq := &ClientInfo{
		FirstName:      "first name",
		LastName:       "last name",
		Active:         true,
		Locale:         "en",
		CountryCode:    "62",
		PhoneNumber:    "8123123",
		SubmitDate:     time.Now(),
		ActivationDate: time.Now(),
	}

	response, err := client.CreateClient(clientReq, "merchant_user_id", "merchant_name")
	require.Nil(t, err)
	require.NotNil(t, response)
	require.Equal(t, int64(1001), response.ID)
}
