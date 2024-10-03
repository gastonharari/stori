package service

import (
	"context"
	"stori/internal/summary/domain"
)

//go:generate mockery --name=Repository --output=./mocks --structname=Repository --filename=repository.go
type Repository interface {
	SendEmail(ctx context.Context, emailData domain.EmailData) error
}
