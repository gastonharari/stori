package service

import (
	"context"
	"stori/internal/transactions/domain"
)

type Repository interface {
	ReadFile(ctx context.Context, path string) ([]domain.Transaction, error)
}
