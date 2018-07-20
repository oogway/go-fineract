package fineract

import (
	"testing"
	"github.com/stretchr/testify/require"
	"time"
)

func TestCreateClient(t *testing.T) {
	client, err := makeClient(true)
	require.Nil(t, err)
	clientReq := &ClientInfo{
		FirstName:      "first name",
		LastName:       "last name",
		Active:         true,
		Locale:         "en",
		MobileNo:       "628123123",
		SubmitDate:     time.Now(),
		ActivationDate: time.Now(),
	}
	response, err := client.CreateClient(clientReq, "merchant_user_id", "merchant_name")
	require.Nil(t, err)
	require.NotNil(t, response)
	require.Equal(t, int64(1001), response.ID)
}