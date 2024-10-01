package daos

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

type TransactionDAO struct {
	ID     string
	Date   time.Time
	Amount float64
}

func (d TransactionDAO) ToDomain() domain.Transaction {
	return domain.Transaction{
		ID:     d.ID,
		Date:   d.Date,
		Amount: d.Amount,
	}
}
