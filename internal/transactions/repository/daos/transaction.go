package daos

import (
	"stori/internal/transactions/domain"
	"time"
)

const DateFormat = "1/2"

type TransactionDAO struct {
	ID     string    `json:"id"`
	Date   time.Time `json:"date"`
	Amount float64   `json:"transaction"`
}

func (d TransactionDAO) ToDomain() domain.Transaction {
	return domain.Transaction{
		ID:     d.ID,
		Date:   d.Date,
		Amount: d.Amount,
	}
}
