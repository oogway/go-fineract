package fineract

import (
	"testing"
	"time"

	"log"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSuiteMockAddress(t *testing.T) {
	if !testing.Short() {
		t.Skip("Skipped mock tests in long mode")
	}
	client, err := makeClient(true)
	if err != nil {
		t.Fatal(err)
	}
	clientId := "144"
	Suite(t, client, clientId)
}

func TestSuiteAddress(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipped integrated tests in short mode")
	}
	client, err := makeClient(false)
	if err != nil {
		t.Fatal(err)
	}
	clientReq := &ClientInfo{
		FirstName:      "first name",
		LastName:       "last name",
		Active:         true,
		Locale:         "en",
		CountryCode:    "62",
		PhoneNumber:    toString(random(81100200000, 81100249999)),
		SubmitDate:     time.Now(),
		ActivationDate: time.Now(),
	}

	merchantName := "toko"
	merchantClientId := toString(random(11111111, 88888888))
	response, err := client.CreateClient(clientReq, merchantClientId, merchantName)
	log.Println(err)
	require.Nil(t, err)
	require.NotNil(t, response)

	SuiteAddress(t, client, toString(response.ID))
}

func SuiteAddress(t *testing.T, client *Client, clientId string) {
	t.Run("TestCreateAddress", func(t *testing.T) {
		add := Address{
			AddressLine1: "Jl. Medan Merdeka Utara",
			AddressLine2: "No. 1, Gambir",
			AddressLine3: "014/002",
			City:         "Central Jakarta City",
			Country:      "27",
			PostalCode:   "10110",
		}
		req := CreateAddressRequest{
			AddressTypeCode: "25",
			ClientId:        clientId,
			Address:         add,
		}

		_, err := client.CreateAddress(&req)
		assert.NoErrorf(t, err, "create address failed with %s", err)
	})
}
