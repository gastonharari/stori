package stori

import (
	"log"
	"stori/cmd/processfiletransactions"
)

func main() {
	handlerProcessTransactions, err := processfiletransactions.NewHandler()
	if err != nil {
		log.Fatalf("Error processing transactions: %v", err)
	}
	err = handlerProcessTransactions.Handle()
	if err != nil {
		log.Fatalf("Error processing transactions: %v", err)
	}
}
