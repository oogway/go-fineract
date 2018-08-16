package fineract

import (
	"testing"
	"time"

	"log"

	"github.com/bmizerany/assert"
	"github.com/stretchr/testify/require"
)

func TestSuiteMockClient(t *testing.T) {
	if !testing.Short() {
		t.Skip("Skipped mock tests in long mode")
	}
	client, err := makeClient(true)
	if err != nil {
		t.Fatal(err)
	}
	ClientSuite(t, client)
}

func TestSuiteClient(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipped integrated tests in short mode")
	}
	client, err := makeClient(false)
	if err != nil {
		t.Fatal(err)
	}
	ClientSuite(t, client)
	ClientISuite(t, client)
}

func ClientISuite(t *testing.T, client *Client) {
	t.Run("TestCreateClient: 2 requests with same custId returns same fi-clientId", func(t *testing.T) {
		custId := toString(random(1, 999))

		clientReq := &ClientInfo{
			FirstName:      "first name",
			LastName:       "last name",
			Active:         true,
			Locale:         "en",
			CountryCode:    "62",
			PhoneNumber:    toString(random(81100200000, 81100249999)),
			SubmitDate:     time.Now(),
			ActivationDate: time.Now(),
			DeclaredIncome: 10000,
			Occupation:     "student",
			Email:          "abc@gmail.com",
			ExternalId:     toString(random(1, 999)),
		}

		fResponse, err := client.CreateClient(clientReq, custId, "toko")
		require.Nil(t, err)

		sResponse, err := client.CreateClient(clientReq, custId, "toko")
		require.Nil(t, err)

		assert.Equal(t, fResponse.ID, sResponse.ID, "new client got created despite supplying same merchant customerID")
	})

	t.Run("TestCreateClient: 2 requests with same externalId returns error", func(t *testing.T) {
		externalId := toString(random(1, 999))

		clientReq := &ClientInfo{
			FirstName:      "first name",
			LastName:       "last name",
			Active:         true,
			Locale:         "en",
			CountryCode:    "62",
			PhoneNumber:    toString(random(81100200000, 81100249999)),
			SubmitDate:     time.Now(),
			ActivationDate: time.Now(),
			DeclaredIncome: 10000,
			Occupation:     "student",
			Email:          "abc@gmail.com",
			ExternalId:     externalId,
		}

		_, err := client.CreateClient(clientReq, toString(random(1, 999)), "toko")
		require.Nil(t, err)

		_, err = client.CreateClient(clientReq, toString(random(1, 999)), "toko")
		assert.NotEqual(t, nil, err, "duplicate customer creation with same externalId should have failed")
	})
}

func ClientSuite(t *testing.T, client *Client) {
	t.Run("TestCreateClient", func(t *testing.T) {
		clientReq := &ClientInfo{
			FirstName:      "first name",
			LastName:       "last name",
			Active:         true,
			Locale:         "en",
			CountryCode:    "62",
			PhoneNumber:    toString(random(81100200000, 81100249999)),
			SubmitDate:     time.Now(),
			ActivationDate: time.Now(),
			DeclaredIncome: 10000,
			Occupation:     "student",
			Email:          "abc@gmail.com",
			ExternalId:     toString(random(1111, 9999)),
		}

		response, err := client.CreateClient(clientReq, toString(random(0, 99999)), "toko")
		if err != nil {
			t.Fatalf(err.Error())
		}
		require.NotNil(t, response)
		require.NotNil(t, response.ID)
		log.Println(response.ID)
	})
}
