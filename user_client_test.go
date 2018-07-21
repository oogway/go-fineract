// build test
package fineract

import (
	"log"
	"testing"
)

func TestUserInfo(t *testing.T) {
	client, err := makeClient(true)
	if err != nil {
		t.Fatalf("Cannot create new client: %v", err)
	}
	if err != nil {
		log.Println(err)
	}

	clientId := "337"
	resp, err := client.GetClientInfo(clientId)
	if err != nil {
		log.Println(err)
	}

	log.Println(resp)
}
