package service

import (
	"context"
	"stori/internal/transactions/domain"
)

type Repository interface {
	Create(ctx context.Context, transactions []domain.Transaction) error
}
