package fineractor

import (
	"log"
)

func (client NewClient) IncrementFund(request []byte) error {
	log.Println("Fineractor increment fund called for NewClient")
	return nil
}
