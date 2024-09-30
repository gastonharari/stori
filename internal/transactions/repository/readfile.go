package repository

import (
	"context"
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"time"

	"stori/internal/transactions/domain"
	"stori/internal/transactions/repository/daos"
	"stori/pkg/structs"
)

func (r Repository) ReadFile(ctx context.Context, filename string) ([]domain.Transaction, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.Join(domain.ErrorFileNotFound, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var transactions []daos.TransactionDAO
	for _, record := range records {
		id := record[0]
		if id == "" {
			return nil, domain.ErrorInvalidID
		}

		date, err := time.Parse(daos.DateFormat, record[1])
		if err != nil {
			return nil, errors.Join(domain.ErrorInvalidDate, err)
		}

		amount, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, errors.Join(domain.ErrorInvalidAmount, err)
		}

		transactions = append(transactions, daos.TransactionDAO{
			ID:     id,
			Date:   date,
			Amount: amount,
		})
	}

	return structs.Map(transactions, daos.TransactionDAO.ToDomain), nil
}
