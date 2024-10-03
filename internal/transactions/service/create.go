package service

import (
	"context"
	"stori/internal/transactions/domain"
)

func (s Service) Create(ctx context.Context, transactions []domain.Transaction) error {
	err := s.repository.Create(ctx, transactions)
	if err != nil {
		return err
	}
	//TODO: send sns event created transaction
	return nil
}
