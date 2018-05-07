package fineractor

import (
	"fmt"
	"testing"
)

func TestNewClient(t *testing.T) {
	hostName := "https://demo.openmf.org/fineract-provider/api/v1/savingsaccounts/"
	userName := "mifos"
	password := "password"

	t.Run("Should decrement fund", func(t *testing.T) {
		client, err := NewClient(hostName, userName, password, FineractOption{})
		// s := NewServer(client)
		if err != nil {
			t.Fatalf("Cannot create new client: %v", err)
		}

		fields := client.GetFieldsMap()
		if (fields["HostName"] != hostName) || (fields["UserName"] != userName) || (fields["Password"] != password) {
			t.Fatalf("Values not set properly for client: %v", client)
		}

		clientNext, err := NewClient(hostName, userName, password, FineractOption{})

		if fmt.Sprintf("%p", client) != fmt.Sprintf("%p", clientNext) {
			t.Fatal("Client should be initialised only once and same client should be returned on next call")
		}
	})
}
