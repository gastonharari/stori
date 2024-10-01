package processtransactions

import (
	"context"
	dmnsummary "stori/internal/summary/domain"
	dmntransactions "stori/internal/transactions/domain"
)

//go:generate mockery --name=Service --output=./mocks --structname=Service --filename=service.go
//go:generate mockery --name=CreateSummaryUC --output=./mocks --structname=CreateSummaryUC --filename=createsummaryuc.go
//go:generate mockery --name=SendEmailUC --output=./mocks --structname=SendEmailUC --filename=sendemailuc.go

type Service interface {
	ReadFile(ctx context.Context, path string) ([]dmntransactions.Transaction, error)
}

type CreateSummaryUC interface {
	Exec(ctx context.Context, transactions []dmntransactions.Transaction) dmnsummary.Summary
}

type SendEmailUC interface {
	Exec(ctx context.Context, summary dmnsummary.Summary) error
}
