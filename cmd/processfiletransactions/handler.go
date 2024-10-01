package processfiletransactions

import (
	"context"
	"flag"
	"log"
)

func (h handler) Handle() error {
	ctx := context.Background()

	path := flag.String("file", "", "Path to the CSV file containing transactions")
	email := flag.String("email", "", "Email address to send the summary to")
	flag.Parse()

	if *path == "" {
		log.Fatal("Please provide a valid path to the CSV file using the --file flag.")
	}

	if *email == "" {
		log.Fatal("Please provide a valid email address using the --email flag.")
	}
	log.Println("Starting transaction processing...")

	err := h.UC.Exec(ctx, *path, *email)
	if err != nil {
		log.Fatalf("Error processing transactions: %v", err)
	}

	log.Println("Transaction processing completed successfully.")
	return nil
}
