package fineract

import (
	"fmt"
	"testing"
)

func TestNewClient(t *testing.T) {
	t.Run("TestNewClient", func(t *testing.T) {
		client, err := NewClient("https://"+fineractHost, fineractUser, fineractPassword, FineractOption{SkipVerify: true})
		if err != nil {
			t.Fatalf("Cannot create new client: %v", err)
		}

		clientNext, err := NewClient("https://"+fineractHost, fineractUser, fineractPassword, FineractOption{SkipVerify: true})

		if fmt.Sprintf("%p", client) != fmt.Sprintf("%p", clientNext) {
			t.Fatal("Client should be initialised only once and same client should be returned on next call")
		}
	})
}
