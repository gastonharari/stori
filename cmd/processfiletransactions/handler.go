package processfiletransactions

import (
	"context"
	"flag"
	"log"
)

func (h handler) Handle() error {
	ctx := context.Background()

	path := flag.String("file", "", "Path to the CSV file containing transactions")
	flag.Parse()

	if *path == "" {
		log.Fatal("Please provide a valid path to the CSV file using the --file flag.")
	}

	err := h.UC.Exec(ctx, *path)
	if err != nil {
		log.Fatalf("Error processing transactions: %v", err)
	}

	log.Println("Transaction processing completed successfully.")
	return nil
}
