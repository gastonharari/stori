package service

import (
	"context"
	"stori/internal/transactions/domain"
)

func (s Service) ReadFile(ctx context.Context, path string) ([]domain.Transaction, error) {
	transactions, err := s.repository.ReadFile(ctx, path)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
