package fineractor

import (
	"log"
)

func (mockClient NewMockClient) IncrementFund(request []byte) error {
	log.Println("Fineractor increment fund called for NewMockClient")
	return nil
}
