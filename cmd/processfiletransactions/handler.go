package processfiletransactions

import (
	"context"
	"flag"
	"log"
	"stori/cmd/processfiletransactions/dtos"
)

func (h handler) Handle() error {
	ctx := context.Background()

	path := flag.String("file", "", "Path to the CSV file containing transactions")
	email := flag.String("email", "", "Email address to send the summary to")
	flag.Parse()

	if *path == "" {
		return dtos.ErrMissingFile
	}

	if *email == "" {
		return dtos.ErrMissingEmail
	}
	log.Println("Starting transaction processing...")

	transactions, err := ReadFile(ctx, *path)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		return err
	}

	err = h.UC.Exec(ctx, *email, transactions)
	if err != nil {
		log.Printf("Error processing transactions: %v", err)
		return err
	}

	log.Println("Transaction processing completed successfully.")
	return nil
}
