package dtos

import (
	"stori/internal/transactions/domain"
	"time"
)

const (
	DateFormat        = "1/2"
	HeaderID          = "Id"
	HeaderDate        = "Date"
	HeaderTransaction = "Transaction"
)

type TransactionDTO struct {
	ID     string
	Date   time.Time
	Amount float64
}

func (d TransactionDTO) ToDomain() domain.Transaction {
	return domain.Transaction{
		ID:     d.ID,
		Date:   d.Date,
		Amount: d.Amount,
	}
}
