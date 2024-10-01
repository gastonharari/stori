package service

import (
	"context"
	"stori/internal/summary/domain"
)

type Repository interface {
	SendEmail(ctx context.Context, emailData domain.EmailData) error
}
