package fineract

import (
	"testing"
	"time"

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
		}

		response, err := client.CreateClient(clientReq, toString(random(0, 99999)), "toko")
		if err != nil {
			t.Fatalf(err.Error())
		}
		require.NotNil(t, response)
		require.NotNil(t, response.ID)
	})
}
