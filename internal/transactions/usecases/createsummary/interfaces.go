package createsummary

import (
	"context"
	"stori/internal/transactions/domain"
)

type Service interface {
	ReadFile(ctx context.Context, path string) ([]domain.Transaction, error)
}
