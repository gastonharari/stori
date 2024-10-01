package sendemail

import (
	"context"
	"stori/internal/summary/domain"
)

//go:generate mockery --name=Service --output=./mocks --structname=Service --filename=service.go

type Service interface {
	SendEmail(ctx context.Context, summary domain.Summary, email string) error
}
