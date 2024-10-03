package processfiletransactions

import (
	"context"
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"time"

	"stori/cmd/processfiletransactions/dtos"
	"stori/internal/transactions/domain"
	"stori/pkg/structs"
)

func ReadFile(ctx context.Context, filename string) ([]domain.Transaction, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.Join(domain.ErrorFileNotFound, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	header, err := reader.Read()
	if err != nil {
		return nil, err
	}
	if header[0] != dtos.HeaderID || header[1] != dtos.HeaderDate || header[2] != dtos.HeaderTransaction {
		return nil, domain.ErrorInvalidHeader
	}
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var transactions []dtos.TransactionDTO
	for _, record := range records {
		id := record[0]
		if id == "" {
			return nil, domain.ErrorInvalidID
		}

		date, err := time.Parse(dtos.DateFormat, record[1])
		if err != nil {
			return nil, errors.Join(domain.ErrorInvalidDate, err)
		}

		amount, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, errors.Join(domain.ErrorInvalidAmount, err)
		}

		transactions = append(transactions, dtos.TransactionDTO{
			ID:     id,
			Date:   date,
			Amount: amount,
		})
	}

	return structs.Map(transactions, dtos.TransactionDTO.ToDomain), nil
}
