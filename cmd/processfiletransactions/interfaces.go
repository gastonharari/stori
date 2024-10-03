package processfiletransactions

import (
	"context"
	"stori/internal/transactions/domain"
)

type ProcesstransactionsUC interface {
	Exec(ctx context.Context, userEmail string, transactions []domain.Transaction) error
}
